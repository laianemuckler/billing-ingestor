package service

import (
	"billing-ingestor/internal/pulses/domain"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type PulseService struct {
	repository domain.PulseRepository
}

func NewPulseService(repository domain.PulseRepository) *PulseService {
	return &PulseService{
		repository: repository,
	}
}

func (s *PulseService) ProcessPulse(ctx context.Context, input domain.PulseInput) error {
	log.Printf("Recebido novo Pulse: Tenant=%s | SKU=%s | Unit=%s | Amount=%d",
		input.TenantId, input.ProductSKU, input.UseUnit, input.UsedAmount)

	if input.TenantId == "" || input.ProductSKU == "" || input.UseUnit == "" || input.UsedAmount == 0 {
		return fmt.Errorf("missing required fields")
	}

	pulse := domain.Pulse{
		ID:         generateID(),
		TenantId:   input.TenantId,
		ProductSKU: input.ProductSKU,
		UsedAmount: input.UsedAmount,
		UseUnit:    input.UseUnit,
		CreatedAt:  time.Now(),
	}

	key := fmt.Sprintf("%s:%s:%s:%s", pulse.TenantId, pulse.ProductSKU, pulse.UseUnit, pulse.CreatedAt.Format("2006-01-02"))

	err := s.repository.StorePulse(ctx, key, pulse)
	if err != nil {
		return err
	}

	return nil
}

func (s *PulseService) GetAggregatedData(ctx context.Context) ([]domain.PulseAggregate, error) {
	return s.repository.GetAggregatedData(ctx)
}

func (s *PulseService) GetAggregatedDataByKey(ctx context.Context, key string) (domain.PulseAggregate, error) {
	return s.repository.GetAggregatedDataByKey(ctx, key)
}

func (s *PulseService) SendAggregatedData(ctx context.Context) error {
	aggregates, err := s.repository.GetAggregatedData(ctx)
	if err != nil {
		return err
	}

	dataByDay := make(map[string][]domain.PulseAggregate)

	for _, aggregate := range aggregates {
		dataByDay[aggregate.AggregationDate] = append(dataByDay[aggregate.AggregationDate], aggregate)
	}

	for date, aggregatesForDay := range dataByDay {
		err := sendToProcessor(aggregatesForDay, date)
		if err != nil {
			return fmt.Errorf("failed to send data for date %s: %v", date, err)
		}
	}

	return s.repository.ClearAggregatedData(ctx)
}

func (s *PulseService) ClearAggregatedData(ctx context.Context) error {
	return s.repository.ClearAggregatedData(ctx)
}

func generateID() string {
	return uuid.New().String()
}

func sendToProcessor(aggregates []domain.PulseAggregate, date string) error {
	log.Printf("Sending %d records for date %s to processor...", len(aggregates), date)
	for _, a := range aggregates {
		log.Printf("Tenant: %s | SKU: %s | Unit: %s | Total: %d", a.TenantId, a.ProductSKU, a.UseUnit, a.TotalUsedAmount)
	}
	return nil
}
