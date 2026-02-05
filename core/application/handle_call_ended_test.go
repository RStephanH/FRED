package application

import (
	"fmt"
	"fred-core/domain/call"
	"testing"
)

func TestNewCallEndedHandler(t *testing.T) {
	store := NewCallStore()
	c := call.NewCall("1234")

	// Step 1: call starts - state should be NEW
	startHandler := NewCallStartedHandler(store)
	startHandler.Handle(call.CallStarted{
		CallID: c.GetID(),
	})
	callAfterStart := store.Get("1234")
	if callAfterStart == nil {
		t.Fatalf("call should exist after start")
	}
	if callAfterStart.State != call.CallStateNew {
		t.Fatalf("call state should be NEW after start, got %v", callAfterStart.State)
	}
	fmt.Printf("Step 1 - Call state after started: %v\n", callAfterStart.State)

	// Step 2: call is answered - state should be ACTIVE
	answerHandler := NewCallAnsweredHandler(store)
	if answerHandler != nil {
		answerHandler.Handle(call.CallAnswered{
			CallID: c.GetID(),
		})
		callAfterAnswer := store.Get("1234")
		if callAfterAnswer == nil {
			t.Fatalf("call should exist after answer")
		}
		if callAfterAnswer.State != call.CallStateActive {
			t.Fatalf("call state should be ACTIVE after answer, got %v", callAfterAnswer.State)
		}
		fmt.Printf("Step 2 - Call state after answered: %v\n", callAfterAnswer.State)
	}

	// Step 3: call is ended - state should be ENDED
	endHandler := NewCallEndedHandler(store)
	if endHandler != nil {
		endHandler.Handle(call.CallEnded{
			CallID: c.GetID(),
		})
		callAfterEnd := store.Get("1234")
		if callAfterEnd == nil {
			t.Fatalf("call should exist after end")
		}
		if callAfterEnd.State != call.CallStateEnded {
			t.Fatalf("call state should be ENDED after end, got %v", callAfterEnd.State)
		}
		fmt.Printf("Step 3 - Call state after ended: %v\n", callAfterEnd.State)
	}

	// Step 4: try to change from ENDED to ACTIVE - should remain ENDED (invalid state transition)
	if answerHandler != nil {
		answerHandler.Handle(call.CallAnswered{
			CallID: c.GetID(),
		})
		callAfterInvalidTransition := store.Get("1234")
		if callAfterInvalidTransition == nil {
			t.Fatalf("call should still exist")
		}
		if callAfterInvalidTransition.State != call.CallStateEnded {
			t.Fatalf("call state should still be ENDED after invalid transition attempt, got %v", callAfterInvalidTransition.State)
		}
		fmt.Printf("Step 4 - Call state after invalid transition (ENDED->ACTIVE): %v (unchanged as expected)\n", callAfterInvalidTransition.State)
	}
}
