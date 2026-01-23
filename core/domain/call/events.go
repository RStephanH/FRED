package call

// This file contain the event that will be handled , it's kind of abstraction or translation of the ARI event
// Events describe something that already happened.
// They do not do anything

type CallStarted struct {
	CallID string
}

func NewCall(id string) *Call {
	return &Call{
		ID:    id,
		State: CallStateNew,
	}
}
