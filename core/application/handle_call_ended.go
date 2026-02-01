package application

import "fred-core/domain/call"

type CallEndedHandler struct {
	store *CallStore
}

func NewCallEndedHandler(store *CallStore) *CallEndedHandler {
	return &CallEndedHandler{store: store}
}

func (h *CallEndedHandler) Handle(event call.CallEnded) {
	c := h.store.Get(event.CallID)
	if c == nil {
		return
	}
	c.On(event)
}
