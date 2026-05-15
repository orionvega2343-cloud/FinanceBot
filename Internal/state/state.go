package state

import "sync"

type State struct {
	Step       int
	Type       string
	CategoryID int
	Sum        float64
	Comment    string
}

type StateManager struct {
	sm map[int64]*State
	mu sync.RWMutex
}

func NewStateManager() *StateManager {
	return &StateManager{sm: make(map[int64]*State)}
}

func (sm *StateManager) Get(UserId int64) (*State, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	state, ok := sm.sm[UserId]
	if !ok {
		return nil, false
	}
	return state, true
}

func (sm *StateManager) Set(UserId int64, state *State) {
	sm.mu.Lock()
	sm.sm[UserId] = state
	sm.mu.Unlock()
}

func (sm *StateManager) Del(UserId int64) {
	sm.mu.Lock()
	delete(sm.sm, UserId)
	sm.mu.Unlock()
}
