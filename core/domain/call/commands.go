package call

type AnswerCall struct {
	CallID string
}

type Playsound struct {
	CallID string
	File   string
}

type StartRecording struct {
	CallID string
	Format string
}

type StopPlayback struct {
}
