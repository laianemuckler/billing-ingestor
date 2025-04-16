package domain

import "time"

type Pulse struct {
	ID         string
	TenantId   string
	ProductSKU string
	UsedAmount int
	UseUnit    string
	CreatedAt  time.Time
}

type PulseInput struct {
	TenantId   string `json:"tenant_id"`
	ProductSKU string `json:"product_sku"`
	UsedAmount int    `json:"used_amount"`
	UseUnit    string `json:"use_unit"`
}

type PulseAggregate struct {
	TenantId        string
	ProductSKU      string
	UseUnit         string
	TotalUsedAmount int
	AggregationDate string
	PulseKey			string
}
