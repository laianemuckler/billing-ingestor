package memory

import (
	"billing-ingestor/internal/pulses/domain"
	"context"
	"fmt"
	"sync"
)

type pulseRepository struct {
	mu     sync.RWMutex
	pulses map[string]domain.PulseAggregate
}

func NewPulseRepository() *pulseRepository {
	return &pulseRepository{
		pulses: make(map[string]domain.PulseAggregate),
	}
}

func (r *pulseRepository) StorePulse(ctx context.Context, key string, pulse domain.Pulse) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.pulses[key]
	if exists {
		existing.TotalUsedAmount += pulse.UsedAmount
		r.pulses[key] = existing
	} else {
		r.pulses[key] = domain.PulseAggregate{
			TenantId:        pulse.TenantId,
			ProductSKU:      pulse.ProductSKU,
			UseUnit:         pulse.UseUnit,
			TotalUsedAmount: pulse.UsedAmount,
			AggregationDate: pulse.CreatedAt.Format("2006-01-02"),
            PulseKey:        key,
		}
	}

	return nil
}

func (r *pulseRepository) GetAggregatedData(ctx context.Context) ([]domain.PulseAggregate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var aggregates []domain.PulseAggregate
	for _, agg := range r.pulses {
		aggregates = append(aggregates, agg)
	}
	return aggregates, nil
}

func (r *pulseRepository) GetAggregatedDataByKey(ctx context.Context, key string) (domain.PulseAggregate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	agg, exists := r.pulses[key]
	if !exists {
		return domain.PulseAggregate{}, fmt.Errorf("no data found for key: %s", key)
	}
	return agg, nil
}

func (r *pulseRepository) ClearAggregatedData(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.pulses = make(map[string]domain.PulseAggregate)
	return nil
}

func (r *pulseRepository) SeedAggregatedData() {
	r.pulses["tenant1:sku1:unit1:2025-04-14"] = domain.PulseAggregate{
		TenantId:        "tenant1",
		ProductSKU:      "sku1",
		UseUnit:         "unit1",
		TotalUsedAmount: 150,
		AggregationDate: "2025-04-14",
        PulseKey:       "tenant1:sku1:unit1:2025-04-14",
	}

	r.pulses["tenant2:sku2:unit2:2025-04-14"] = domain.PulseAggregate{
		TenantId:        "tenant2",
		ProductSKU:      "sku2",
		UseUnit:         "unit2",
		TotalUsedAmount: 90,
		AggregationDate: "2025-04-14",
        PulseKey:       "tenant2:sku2:unit2:2025-04-14",
	}
}
