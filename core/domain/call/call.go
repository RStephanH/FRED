package call

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
func (c *Call) On(event interface{}) {
	switch e := event.(type) {
	case CallStarted:
		c.State = CallStateNew
	case CallAnswered:
		c.State = CallStateActive
	case CallEnded:
		c.State = CallStateEnded
	}
}
