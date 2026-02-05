package application

import "fred-core/domain/call"

type CallStartedHandler struct {
	Store *CallStore
}

func NewCallStartedHandler(store *CallStore) *CallStartedHandler {
	return &CallStartedHandler{
		Store: store,
	}
}

func (h *CallStartedHandler) Handle(event call.CallStarted) {
	newCall := call.NewCall(event.CallID)
	h.Store.Add(newCall)
}
