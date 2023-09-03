package campaign

import (
	"github.com/ozkansen/campaign-module/pkg/time"
)

type Campaign struct {
	Name        string
	ProductCode string

	// Duration is given in hours.
	Duration int

	// PriceManipulationLimit the maximum percentage that you can increase or
	// decrease the price of product according to demand
	PriceManipulationLimit int
	TargetSalesCount       int
	CreatedAt              time.Time
}

func NewCampaign(name, productCode string, duration, priceManipulationLimit, targetSalesCount int) (*Campaign, error) {
	if name == "" ||
		productCode == "" ||
		duration <= 0 ||
		priceManipulationLimit <= 0 ||
		priceManipulationLimit >= 100 ||
		targetSalesCount <= 0 {
		return nil, ErrInvalidValue
	}
	return &Campaign{
		Name:                   name,
		ProductCode:            productCode,
		Duration:               duration,
		PriceManipulationLimit: priceManipulationLimit,
		TargetSalesCount:       targetSalesCount,
		CreatedAt:              time.Now(),
	}, nil
}
