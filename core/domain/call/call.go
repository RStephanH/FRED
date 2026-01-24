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
