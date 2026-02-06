package call

import "fmt"

type CallDecision struct {
}

func (d CallDecision) OnEvent(call Call, event Event) {
	switch event.(type) {
	case CallStarted:
		// return the command adapted so for now let's just print something
		// return []Command{}
		fmt.Println("Call has been started")
	}
}
