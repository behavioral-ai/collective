package content

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/behavioral-ai/core/io"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"net/url"
)

type text struct {
	Value string
}

/*
func toAgent(resolver Resolution) messaging.Agent {
	if resolver == nil {
		return nil
	}
	if r, ok := any(resolver).(resolution); ok {
		return r.agent
	}
	return nil
}

*/

// resolutionFunc - data store function
type resolutionFunc func(method, name, author string, body []byte, version int) ([]byte, *messaging.Status)

// addActivityFunc -
type addActivityFunc func(hostName string, agent messaging.Agent, event, source string, content any)

type resolution struct {
	//do       http2.Exchange
	notifier messaging.NotifyFunc
	activity addActivityFunc
	hosts    []string
	agent    *agentT
}

func newHttpResolver() Resolution {
	r := new(resolution)
	r.agent = newContentAgent(false, nil)
	return r
}

// GetValue - resolution get
func (r *resolution) GetValue(nsName string, version int) ([]byte, *messaging.Status) {
	return r.agent.getValue(nsName, version)
}

// PutValue - resolution put
func (r *resolution) PutValue(nsName, author string, content any, version int) *messaging.Status {
	var buf []byte
	var err error

	if nsName == "" {
		err = errors.New(fmt.Sprintf("nsName is empty on call to PutValue()"))
		return messaging.NewStatusError(http.StatusBadRequest, err, r.agent.Uri())
	}
	if content == nil {
		err = errors.New(fmt.Sprintf("content is nil on call to PutValue() for nsName : %v", nsName))
		return messaging.NewStatusError(http.StatusNoContent, err, r.agent.Uri())
	}
	switch ptr := content.(type) {
	case string:
		v := text{ptr}
		buf, err = json.Marshal(v)
		if err != nil {
			return messaging.NewStatusError(messaging.StatusJsonEncodeError, err, r.agent.Uri())
		}
	case []byte:
		buf = ptr
	case *url.URL:
		buf, err = io.ReadFile(ptr)
		if err != nil {
			return messaging.NewStatusError(messaging.StatusIOError, err, r.agent.Uri())
		}
	default:
		buf, err = json.Marshal(ptr)
		if err != nil {
			return messaging.NewStatusError(messaging.StatusJsonEncodeError, err, r.agent.Uri())
		}
	}
	if len(buf) == 0 {
		err = errors.New(fmt.Sprintf("content is empty on call to PutValue() for nsName : %v", nsName))
		return messaging.NewStatusError(http.StatusNoContent, err, r.agent.Uri())
	}
	return r.agent.putValue(nsName, author, buf, version)
}

// GetAttributes - resolution get
func (r *resolution) GetAttributes(nsName string) (map[string]string, *messaging.Status) {
	return r.agent.getAttributes(nsName)
}

// PutAttributes - resolution put
func (r *resolution) PutAttributes(nsName, author string, m map[string]string) *messaging.Status {
	return r.agent.putAttributes(nsName, author, m)
}

// AddActivity - resolution activity
func (r *resolution) AddActivity(agent messaging.Agent, event, source string, content any) {
	if r.activity != nil {
		r.activity(r.agent.hostName, agent, event, source, content)
	} else {
		// TODO: add call to append activity, include appHostName
	}
}

// Notify - resolution notify
func (r *resolution) Notify(e messaging.Event) {
	if r.notifier != nil {
		r.notifier(e)
	} else {
		// TODO: add call to notify, include appHostName
	}
}
