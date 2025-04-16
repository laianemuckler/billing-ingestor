package domain

import "context"

type PulseRepository interface {
	StorePulse(ctx context.Context, key string, pulse Pulse) error
	GetAggregatedData(ctx context.Context) ([]PulseAggregate, error)
	GetAggregatedDataByKey(ctx context.Context, key string) (PulseAggregate, error)
	ClearAggregatedData(ctx context.Context) error
	SeedAggregatedData()
}
