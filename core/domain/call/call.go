package call

import "log"

type Call struct {
	ID    string
	State CallState
}

func NewCall(id string) *Call {
	return &Call{
		ID:    id,
		State: CallStateNew,
	}
}

func (c *Call) GetID() string {
	return c.ID
}

func (c *Call) GetState() CallState {
	return c.State
}

// TODO: the event type in a variable and also indicates the timestamp of the event
func (c *Call) On(event interface{}) {
	switch event.(type) {
	case CallStarted:
		// Only valid if call is brand new (creation case)
		if c.State == CallStateNew {
			// already NEW, nothing to change
		}
	case CallAnswered:
		// Only NEW → ACTIVE is allowed
		if c.State == CallStateNew {
			c.State = CallStateActive
		}
	case CallEnded:
		// NEW → ENDED or ACTIVE → ENDED
		if c.State == CallStateNew || c.State == CallStateActive {
			c.State = CallStateEnded
		}
	default:
		log.Fatalf("Unknown event type: %T, event: %v", event, event)
	}
}
