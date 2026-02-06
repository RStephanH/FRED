package call

// This file contain the event that will be handled , it's kind of abstraction or translation of the ARI event
// Events describe something that already happened.
// They do not do anything

type Event interface {
	GetCallID() string
}

type CallStarted struct {
	CallID string
}

func (cs CallStarted) GetCallID() string {
	return cs.CallID
}

type CallAnswered struct {
	CallID string
}

func (ca CallAnswered) GetCallID() string {
	return ca.CallID
}

type CallEnded struct {
	CallID string
}

func (ce CallEnded) GetCallID() string {
	return ce.CallID
}
