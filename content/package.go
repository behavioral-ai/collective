package content

import (
	"encoding/json"
	"errors"
	"fmt"
	http2 "github.com/behavioral-ai/core/http"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

type ActivityFunc func(hostName, agentUri, event, source string, content any)

// Resolution - in the real world
type Resolution interface {
	GetValue(nsName string, version int) ([]byte, *messaging.Status)
	AddValue(nsName, author string, content any, version int) *messaging.Status
	GetAttributes(nsName string) (map[string]string, *messaging.Status)
	AddAttributes(nsName, author string, m map[string]string) *messaging.Status
	AddActivity(agent messaging.Agent, event, source string, content any)
	//Notify(e messaging.Event)
}

// Resolver - content resolution in the real world
var (
	Resolver Resolution
	Agent    messaging.Agent
	Exchange http2.Exchange
)

func init() {
	Exchange = http2.Do
	r := newHttpResolver()
	Resolver = r
	Agent = r.agent
	r.agent.Run()
}

func Override(activity ActivityFunc, notifier messaging.NotifyFunc, dispatcher messaging.Dispatcher) {
	if agent, ok := any(Resolver).(*agentT); ok {
		agent.activity = activity
		agent.notifier = notifier
		agent.dispatcher = dispatcher
	}
}

// Resolve - generic typed resolution
func Resolve[T any](nsName string, version int, resolver Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", nsName)), AgentNamespaceName)
	}
	body, status := resolver.GetValue(nsName, version)
	if !status.OK() {
		return t, status
	}
	if len(body) == 0 {
		return t, messaging.NewStatusMessage(http.StatusNoContent, fmt.Sprintf("content not found for name: %v", nsName), AgentNamespaceName)
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](nsName, version, resolver)
		if !status1.OK() {
			return t, status1
		}
		*ptr = t1.Value
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, nsName)), AgentNamespaceName)
		}
	}
	return t, messaging.StatusOK()
}
