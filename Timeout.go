package handler

import "time"

type Timeout interface {
	Increment()
	Sleep()
	HasThreadReached() bool
}

type LinearTimeout struct {
	increment int
	threshold int
	seconds   int
}

func NewTimeout(threshold, seconds int) Timeout {
	return &LinearTimeout{
		threshold: threshold,
		increment: 0,
		seconds:   seconds,
	}
}

func (t *LinearTimeout) Increment() {
	t.increment++
}

func (t *LinearTimeout) Sleep() {
	for i := 0; i < t.increment; i++ {
		time.Sleep(time.Second * time.Duration(t.seconds))
	}
}

func (t *LinearTimeout) HasThreadReached() bool {
	return t.increment > t.threshold
}
