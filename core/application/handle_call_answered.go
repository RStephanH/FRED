package application

import "fred-core/domain/call"

type CallAnsweredHandler struct {
	store *CallStore
}

func NewCallAnsweredHandler(store *CallStore) *CallAnsweredHandler {
	return &CallAnsweredHandler{store: store}
}

func (h *CallAnsweredHandler) Handle(event call.CallAnswered) {
	c := h.store.Get(event.CallID)
	if c == nil {
		return
	}
	c.On(event)
}
