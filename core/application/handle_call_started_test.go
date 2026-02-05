package application

import (
	"fred-core/domain/call"
	"testing"
)

func TestNewCallStartedHandler(t *testing.T) {
	store := NewCallStore()
	h := NewCallStartedHandler(store)
	if h != nil {
		c := call.NewCall("123")
		h.Handle(call.CallStarted{
			CallID: c.GetID(),
		})
		call := store.Get("123")
		if call == nil {
			t.Errorf("Inexpected behavior call hasn't be created by the Handle method")
		}
	} else {
		t.Errorf("Inexpected behavior, value of handler=%v", h)
	}
}
