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
		c.State = CallStateNew
	case CallAnswered:
		c.State = CallStateActive
	case CallEnded:
		c.State = CallStateEnded
	default:
		log.Fatalf("Unknown event type: %T, event: %v", event, event)
	}
}
