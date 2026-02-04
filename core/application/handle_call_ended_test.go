package application

import (
	"fmt"
	"fred-core/domain/call"
	"testing"
)

func TestNewCallEndedHandler(t *testing.T) {
	store := NewCallStore()
	c := call.NewCall("1234")

	//Step 1 : call starts
	startHandler := NewCallStartedHandler(store)
	startHandler.Handle(call.CallStarted{
		CallID: c.GetID(),
	})

	//Step 2 ; call is answered
	answerHandler := NewCallAnsweredHandler(store)
	if answerHandler != nil {
		answerHandler.Handle(call.CallAnswered{
			CallID: c.GetID(),
		})
		call := store.Get("1234")
		if call == nil {
			t.Fatalf("call should exist")
		} else {
			fmt.Printf("The value of the call is call = %v, his state is = %v", call.ID, call.State)
		}
	}

	//Step 3 : call is ended
	endHandler := NewCallEndedHandler(store)
	if endHandler != nil {
		endHandler.Handle(call.CallEnded{
			CallID: c.GetID(),
		})
		call := store.Get("1234")
		if call == nil {
			t.Fatalf("call should exist")
		} else {
			fmt.Printf("The value of the call is call = %v, his state is = %v", call.ID, call.State)
		}
	}

}
