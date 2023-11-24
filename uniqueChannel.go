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

func (uc *UniqueChannel) Send(value *MessageData) {
	if value == nil {
		return
	}

	uc.mu.Lock()
	defer uc.mu.Unlock()

	if id := string(value.Id); !uc.seenVals[id] {
		uc.Ch <- value
		uc.seenVals[id] = true
	}
}

func (uc *UniqueChannel) Remove(value *MessageData) {
	if value == nil {
		return
	}

	uc.mu.Lock()
	defer uc.mu.Unlock()

	delete(uc.seenVals, string(value.Id))
}
