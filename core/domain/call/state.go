package call

type CallState string

const (
	CallStateNew    CallState = "NEW"
	CallStateActive CallState = "ACTIVE"
	CallStateEnded  CallState = "ENDED"
)
