package timeseries

import (
	"fmt"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/collective/operations"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	NamespaceName   = "resiliency:agent/behavioral-ai/collective/timeseries"
	defaultDuration = time.Second * 10
)

var (
	agent *agentT
)

type agentT struct {
	running  bool
	duration time.Duration

	handler  eventing.Agent
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

func init() {
	agent = newAgent(eventing.Handler)
	operations.Register(agent)
}

func newAgent(handler eventing.Agent) *agentT {
	a := new(agentT)
	a.duration = defaultDuration
	a.handler = handler

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if !a.running {
		if m.Event() == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Event() == messaging.StartupEvent {
			a.run()
			a.running = true
			return
		}
		return
	}
	if m.Event() == messaging.ShutdownEvent {
		a.running = false
	}
	switch m.Channel() {
	case messaging.ChannelEmissary:
		a.emissary.Send(m)
	case messaging.ChannelMaster:
		a.master.Send(m)
	case messaging.ChannelControl:
		a.emissary.Send(m)
		a.master.Send(m)
	default:
		a.emissary.Send(m)
	}
}

func (a *agentT) configure(m *messaging.Message) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	}
	// configure
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}

// Run - run the agent
func (a *agentT) run() {
	go masterAttend(a)
	go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}

func (a *agentT) rollup(origin Origin) *messaging.Status {
	_, status := httpRollup(origin)
	if !status.OK() {
		status.WithAgent(a.Uri())
		status.WithMessage(fmt.Sprintf("name %v", origin))
		return status
	}
	return status
}

/*
func (a *agentT) addEvents(events []Event) *messaging.Status {
	if len(events) == 0 {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument events are empty")), a.Uri())
	}
	_, status := httpPutEvents(events)
	if !status.OK() {
		status.WithAgent(a.Uri())
		status.WithMessage(fmt.Sprintf("events"))
		return status
	}
	return status
}


*/
