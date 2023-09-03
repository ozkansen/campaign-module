package main

import (
	"fmt"
	"os"

	"github.com/ozkansen/campaign-module/pkg/time"
	"github.com/ozkansen/campaign-module/service/campaign"
	"github.com/ozkansen/campaign-module/service/campaign_auto_pricing"
	"github.com/ozkansen/campaign-module/service/order"
	"github.com/ozkansen/campaign-module/service/product"
)

func init() {
	// Changes system runtime time and date
	date := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	time.Set(date)
	fmt.Println("date: ", time.Now())
}

func getCampaignAutoPricingService() (*campaign_auto_pricing.CampaignAutoPricing, error) {
	productService, err := product.New(product.WithMemoryProductRepository())
	if err != nil {
		return nil, err
	}
	orderService, err := order.New(order.WithMemoryOrderRepository())
	if err != nil {
		return nil, err
	}
	campaignService, err := campaign.New(campaign.WithMemoryCampaignRepository())
	if err != nil {
		return nil, err
	}
	cp, err := campaign_auto_pricing.New(
		campaign_auto_pricing.WithProductService(productService),
		campaign_auto_pricing.WithOrderService(orderService),
		campaign_auto_pricing.WithCampaignService(campaignService),
		campaign_auto_pricing.WithPriceAlgorithm(campaign_auto_pricing.LineerPricing),
		campaign_auto_pricing.WithOutput(os.Stdout, os.Stderr),
	)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

func main() {
	ap, err := getCampaignAutoPricingService()
	if err != nil {
		panic(err)
	}
	ap.CreateProduct("P1", 100, 100)
	ap.CreateCampaign("C1", "P1", 1, 20, 100)
	ap.CreateCampaign("C1", "P1", 1, 20, 100)
	ap.GetProductInfo("P1")
	ap.CreateOrder("P1", 10)
	ap.GetProductInfo("P1")
	ap.GetCampaignInfo("C1")
	ap.IncreaseTime(1)
	ap.CreateOrder("P1", 10)
	ap.GetCampaignInfo("C1")
	ap.IncreaseTime(1)
	ap.GetCampaignInfo("C1")
	ap.IncreaseTime(1)
	ap.GetCampaignInfo("C1")
	ap.GetProductInfo("P1")
	ap.CreateOrder("P1", 10)
	ap.CreateOrder("P1", 10)
	ap.GetProductInfo("P1")
	ap.GetCampaignInfo("C1")
}
