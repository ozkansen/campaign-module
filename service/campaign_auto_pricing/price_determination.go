package campaign_auto_pricing

import (
	"github.com/ozkansen/campaign-module/pkg/time"
)

type PriceDetermination func(basePrice int64, priceManipulationLimit, totalSales, targetSales int, startedDate time.Time, duration int) (int64, error)

func AlwaysMinPrice(basePrice int64, priceManipulationLimit, _, _ int, startedDate time.Time, duration int) (int64, error) {
	if time.Now().After(startedDate.Add(time.Duration(duration) * time.Hour)) {
		return basePrice, nil
	}
	return basePrice - (basePrice*int64(priceManipulationLimit))/100, nil
}

func LineerPricing(basePrice int64, priceManipulationLimit, totalSales, targetSales int, startedDate time.Time, duration int) (int64, error) {
	if time.Now().After(startedDate.Add(time.Duration(duration) * time.Hour)) {
		return basePrice, nil
	}

	minPrice := basePrice - (basePrice*int64(priceManipulationLimit))/100
	maxPrice := basePrice + (basePrice*int64(priceManipulationLimit))/100
	priceRange := maxPrice - minPrice

	salesRate := float64(totalSales) / float64(targetSales)
	if salesRate == 0 {
		return minPrice, nil
	}
	return minPrice + int64(float64(priceRange)*salesRate), nil
}
