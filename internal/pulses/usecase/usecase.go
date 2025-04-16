package usecase

import (
	"billing-ingestor/internal/pulses/domain"
	"context"
)

type PulseUsecase interface {
	ProcessPulse(ctx context.Context, input domain.PulseInput) error
	GetAggregatedData(ctx context.Context) ([] domain.PulseAggregate, error)
	GetAggregatedDataByKey(ctx context.Context, key string) (domain.PulseAggregate, error)
	SendAggregatedData(ctx context.Context) error
	ClearAggregatedData(ctx context.Context) error
}
