package handler

import (
	"context"

	"time"
)

type syblineConsumer struct {
	ticker       *time.Ticker // periodic ticker
	Messages     *UniqueChannel
	Errs         chan error
	client       SyblineClient
	capacity     int
	holding      *safeInt
	queue        string
	tickerPeriod time.Duration
}

func newSyblineConsumer(capacity int, queue string, tickerPeriod time.Duration, client SyblineClient) *syblineConsumer {
	rv := &syblineConsumer{
		tickerPeriod: tickerPeriod,
		ticker:       time.NewTicker(tickerPeriod),
		Messages:     NewUniqueChannel(capacity),
		Errs:         make(chan error, 1),
		client:       client,
		capacity:     capacity,
		holding:      NewSafeInt(0),
		queue:        queue,
	}

	go rv.run()
	return rv
}

func (s *syblineConsumer) run() {
	for {
		<-s.ticker.C
		s.ticker.Reset(s.tickerPeriod)

		if s.holding.Value() > 0 {
			continue
		}

		messages, err := s.client.GetMessages(context.Background(), s.queue, uint32(s.capacity))
		if err != nil {
			continue
		}

		for _, m := range messages {
			if ok := s.Messages.Send(m); ok {
				s.holding.AddDelta(1)
			}
		}
	}
}

func (s *syblineConsumer) Ack(ctx context.Context, id []byte) error {
	if err := s.client.Ack(ctx, s.queue, id); err != nil {
		return err
	}

	s.holding.AddDelta(-1)
	return nil
}

func (s *syblineConsumer) BatchAck(ctx context.Context, ids [][]byte) error {
	if err := s.client.BatchAck(ctx, s.queue, ids); err != nil {
		return err
	}

	s.holding.AddDelta(-len(ids))
	return nil
}

func (s *syblineConsumer) BatchNack(ctx context.Context, ids [][]byte) error {
	if err := s.client.BatchNack(ctx, s.queue, ids); err != nil {
		return err
	}

	s.holding.AddDelta(-len(ids))
	return nil
}

func (s *syblineConsumer) Nack(ctx context.Context, id []byte) error {
	if err := s.client.Nack(ctx, s.queue, id); err != nil {
		return err
	}

	s.holding.AddDelta(-1)
	return nil
}
