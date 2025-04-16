package fakequeue

import (
	"billing-ingestor/internal/pulses/usecase"
	"context"
	"log"
	"time"
)

type FakeQueue struct {
	usecase usecase.PulseUsecase
}

func NewFakeQueue(usecase usecase.PulseUsecase) *FakeQueue {
	return &FakeQueue{
		usecase: usecase,
	}
}

func (q *FakeQueue) Start() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Starting the batch data sending...")

		aggregates, err := q.usecase.GetAggregatedData(context.Background())
		if err != nil {
			log.Printf("Error retrieving aggregated data before sending: %v", err)
			continue
		}

		for _, aggregate := range aggregates {
			log.Printf("Sending and cleaning data for Pulse with key: %s", aggregate.PulseKey)
		}

		if err := q.usecase.SendAggregatedData(context.Background()); err != nil {
			log.Printf("Error sending aggregated data: %v", err)
			continue
		}

		if err := q.usecase.ClearAggregatedData(context.Background()); err != nil {
			log.Printf("Error clearing aggregated data: %v", err)
		} else {
			log.Println("Aggregated data sent and cleared successfully.")
		}
	}
}
