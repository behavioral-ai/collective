package namespace

import (
	"github.com/behavioral-ai/core/messaging"
)

// master attention
func masterAttend(a *agentT) {
	paused := false
	if paused {
	}

	for {
		select {
		case msg := <-a.master.C:
			switch msg.Name {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				a.masterFinalize()
				return
			default:
			}
		default:
		}
	}
}
