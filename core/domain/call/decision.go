package call

import (
	"fmt"
)

type CallDecision struct {
}

func (d CallDecision) OnEvent(call Call, event Event) []Command {
	switch event.(type) {
	case CallStarted:
		return []Command{AnswerCall{CallID: call.ID}}
	default:
		fmt.Printf("unexpected event %T\n", event)
		return nil
	}
}
