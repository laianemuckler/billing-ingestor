package service

import (
	"billing-ingestor/internal/pulses/domain"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPulseRepository struct {
	mock.Mock
}

func (m *MockPulseRepository) StorePulse(ctx context.Context, key string, pulse domain.Pulse) error {
	args := m.Called(ctx, key, pulse)
	return args.Error(0)
}

func (m *MockPulseRepository) GetAggregatedData(ctx context.Context) ([]domain.PulseAggregate, error) {
	args := m.Called(ctx)
	if ret := args.Get(0); ret != nil {
		return ret.([]domain.PulseAggregate), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPulseRepository) GetAggregatedDataByKey(ctx context.Context, key string) (domain.PulseAggregate, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(domain.PulseAggregate), args.Error(1)
}

func (m *MockPulseRepository) ClearAggregatedData(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockPulseRepository) SeedAggregatedData() {
	m.Called()
}

func TestProcessPulse(t *testing.T) {
	mockRepo := new(MockPulseRepository)
	service := NewPulseService(mockRepo)

	mockRepo.On("StorePulse", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	pulseInput := domain.PulseInput{
		TenantId:   "tenant1",
		ProductSKU: "sku1",
		UseUnit:    "unit1",
		UsedAmount: 100,
	}

	err := service.ProcessPulse(context.Background(), pulseInput)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestProcessPulse_MissingFields(t *testing.T) {
	mockRepo := new(MockPulseRepository)
	service := NewPulseService(mockRepo)

	pulseInput := domain.PulseInput{
		TenantId:   "",
		ProductSKU: "sku1",
		UseUnit:    "unit1",
		UsedAmount: 100,
	}

	err := service.ProcessPulse(context.Background(), pulseInput)

	assert.NotNil(t, err)
	assert.Equal(t, "missing required fields", err.Error())
}

func TestSendAggregatedData(t *testing.T) {
	mockRepo := new(MockPulseRepository)
	service := NewPulseService(mockRepo)

	aggregates := []domain.PulseAggregate{
		{
			TenantId:        "tenant1",
			ProductSKU:      "sku1",
			UseUnit:         "unit1",
			TotalUsedAmount: 100,
			AggregationDate: "2025-04-14",
		},
	}

	mockRepo.On("GetAggregatedData", mock.Anything).Return(aggregates, nil)
	mockRepo.On("ClearAggregatedData", mock.Anything).Return(nil)

	err := service.SendAggregatedData(context.Background())

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestSendAggregatedData_ErrorOnGetAggregatedData(t *testing.T) {
	mockRepo := new(MockPulseRepository)
	service := NewPulseService(mockRepo)

	mockRepo.On("GetAggregatedData", mock.Anything).Return(nil, fmt.Errorf("database error"))

	err := service.SendAggregatedData(context.Background())

	assert.NotNil(t, err)
	assert.Equal(t, "database error", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestClearAggregatedData(t *testing.T) {
	mockRepo := new(MockPulseRepository)
	service := NewPulseService(mockRepo)

	mockRepo.On("ClearAggregatedData", mock.Anything).Return(nil)

	err := service.ClearAggregatedData(context.Background())

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
