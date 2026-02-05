package application

import "fred-core/domain/call"

type CallStore struct {
	calls map[string]*call.Call
}

func NewCallStore() *CallStore {
	return &CallStore{
		calls: make(map[string]*call.Call),
	}
}

func (s *CallStore) Add(c *call.Call) {
	s.calls[c.ID] = c
}

func (s *CallStore) Get(id string) *call.Call {
	return s.calls[id]
}

func (s *CallStore) Remove(id string) {
	delete(s.calls, id)
}
