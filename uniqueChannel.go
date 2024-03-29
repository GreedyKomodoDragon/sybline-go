package handler

import "sync"

type UniqueChannel struct {
	Ch       chan *MessageData
	seenVals map[string]bool
	mu       *sync.Mutex
}

func NewUniqueChannel(cap int) *UniqueChannel {
	return &UniqueChannel{
		Ch:       make(chan *MessageData, cap),
		seenVals: make(map[string]bool),
		mu:       &sync.Mutex{},
	}
}

func (uc *UniqueChannel) Send(value *MessageData) bool {
	if value == nil {
		return false
	}

	uc.mu.Lock()
	defer uc.mu.Unlock()

	if id := string(value.Id); !uc.seenVals[id] {
		uc.seenVals[id] = true
		uc.Ch <- value

		return true
	}

	return false
}

func (uc *UniqueChannel) Remove(value *MessageData) {
	if value == nil {
		return
	}

	uc.mu.Lock()
	defer uc.mu.Unlock()

	delete(uc.seenVals, string(value.Id))
}
