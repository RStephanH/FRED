package call

type Command interface {
	isCommand()
}

type AnswerCall struct {
	CallID string
}

func (c AnswerCall) isCommand() {}

type Playsound struct {
	CallID string
	File   string
}

func (c Playsound) isCommand() {}

type StartRecording struct {
	CallID string
	Format string
}

func (c StartRecording) isCommand() {}

type StopPlayback struct {
}

func (c StopPlayback) isCommand() {}
