package content

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

// initializedEphemeralResolver - in memory resolver, initialized with state
func initializedEphemeralResolver(activity, notify bool) Resolution {
	r := new(resolution)
	if notify {
		r.notifier = messaging.Notify
	} else {
		r.notifier = func(event messaging.Event) {}
	}
	if activity {
		r.activity = func(hostName string, agent messaging.Agent, event, source string, content any) {
			fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
		}
	}
	r.agent = newContentAgent(true, nil)
	r.agent.notifier = r.notifier
	r.agent.Run()
	return r
}
