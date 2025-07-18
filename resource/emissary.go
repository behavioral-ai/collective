package resource

import (
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(a *agentT) {
	var paused = false
	if paused {
	}
	for {
		select {
		case <-a.ticker.T.C:
		default:
		}
		select {
		case msg := <-a.emissary.C:
			switch msg.Name {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				a.emissaryFinalize()
				return
			default:
			}
		default:
		}
	}
}
