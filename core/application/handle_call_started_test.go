package application

import (
	"fred-core/domain/call"
	"testing"
)

func TestNewCallStartedHandler(t *testing.T) {
	store := NewCallStore()
	h := NewCallStartedHandler(store)
	if h != nil {
		h.Handle(
			call.CallStarted{
				CallID: "123",
			})
		call := store.Get("123")
		if call == nil {
			t.Errorf("Inexpected behavior call hasn't be created by the Handle method")
		}
	} else {
		t.Errorf("Inexpected behavior, value of handler=%v", h)
	}
}
