package campaign_auto_pricing

import (
	"errors"
	"fmt"
	"io"

	campaignDomain "github.com/ozkansen/campaign-module/domain/campaign"
	productDomain "github.com/ozkansen/campaign-module/domain/product"
	"github.com/ozkansen/campaign-module/pkg/time"
)

type ProductService interface {
	Create(productCode string, price int64, stock int) error
	Get(productCode string) (*productDomain.Product, error)
}

type OrderService interface {
	Create(productCode string, quantity int, price int64) error
	GetProductTotalSales(productCode string) (int, error)
	GetProductTurnOver(productCode string) (int64, error)
	GetProductAveragePrice(productCode string) (int64, error)
}

type CampaignService interface {
	Create(name string, productCode string, duration int, priceManipulationLimit int, targetSalesCount int) error
	Get(name string) (*campaignDomain.Campaign, error)
	GetFromProductCode(productCode string) (*campaignDomain.Campaign, error)
}

type Configuration func(cp *CampaignAutoPricing) error

func WithOrderService(o OrderService) Configuration {
	return func(cp *CampaignAutoPricing) error {
		cp.orderService = o
		return nil
	}
}

func WithCampaignService(c CampaignService) Configuration {
	return func(cp *CampaignAutoPricing) error {
		cp.campaignService = c
		return nil
	}
}

func WithProductService(p ProductService) Configuration {
	return func(cp *CampaignAutoPricing) error {
		cp.productService = p
		return nil
	}
}

func WithOutput(std io.Writer, err io.Writer) Configuration {
	return func(cp *CampaignAutoPricing) error {
		cp.stdOutput = std
		cp.errOutput = err
		return nil
	}
}

func WithPriceAlgorithm(alg PriceDetermination) Configuration {
	return func(cp *CampaignAutoPricing) error {
		cp.priceDetermination = alg
		return nil
	}
}

type CampaignAutoPricing struct {
	priceDetermination PriceDetermination

	productService  ProductService
	orderService    OrderService
	campaignService CampaignService

	stdOutput io.Writer
	errOutput io.Writer
}

func New(cfgs ...Configuration) (*CampaignAutoPricing, error) {
	cp := &CampaignAutoPricing{}
	for _, cfg := range cfgs {
		err := cfg(cp)
		if err != nil {
			return nil, err
		}
	}
	if cp.productService == nil ||
		cp.campaignService == nil ||
		cp.orderService == nil ||
		cp.priceDetermination == nil ||
		cp.stdOutput == nil ||
		cp.errOutput == nil {
		return nil, errors.New("error defining dependencies")
	}
	return cp, nil
}

func (ap *CampaignAutoPricing) CreateProduct(productCode string, price int64, stock int) {
	err := ap.productService.Create(productCode, price, stock)
	if err != nil {
		ap.writeStdErr("product creating error: %v", err)
		return
	}
	ap.writeStdOut("Product created; code %s, price %d, stock %d", productCode, price, stock)
}

func (ap *CampaignAutoPricing) CreateOrder(productCode string, quantity int) {
	prod, err := ap.productService.Get(productCode)
	if err != nil {
		ap.writeStdErr("product service get: %v", err)
		return
	}
	total, err := ap.orderService.GetProductTotalSales(productCode)
	if err != nil {
		ap.writeStdErr("order service get product total sales: %v", err)
		return
	}
	if (total + quantity) > prod.Stock {
		ap.writeStdErr("not enough stocks available")
		return
	}
	camp, err := ap.campaignService.GetFromProductCode(productCode)
	if err != nil {
		ap.writeStdErr("campaign service get from product code: %v", err)
		return
	}
	currentPrice, err := ap.priceDetermination(prod.Price, camp.PriceManipulationLimit, total, camp.TargetSalesCount, camp.CreatedAt, camp.Duration)
	if err != nil {
		ap.writeStdErr("price determination error: %v", err)
		return
	}
	err = ap.orderService.Create(productCode, quantity, currentPrice)
	if err != nil {
		ap.writeStdErr("order service create: %v", err)
		return
	}
	ap.writeStdOut("Order created; product %s, quantity %d", productCode, quantity)
}

func (ap *CampaignAutoPricing) CreateCampaign(name string, productCode string, duration int, priceManipulationLimit int, targetSalesCount int) {
	_, err := ap.productService.Get(productCode)
	if err != nil {
		ap.writeStdErr("product service get: %v", err)
		return
	}
	err = ap.campaignService.Create(name, productCode, duration, priceManipulationLimit, targetSalesCount)
	if err != nil {
		ap.writeStdErr("campaign service create: %v", err)
		return
	}
	ap.writeStdOut(
		"Campaign created; name %s, product %s, duration %d, limit %d, target sales count %d",
		name, productCode, duration, priceManipulationLimit, targetSalesCount,
	)
}

func (ap *CampaignAutoPricing) GetCampaignInfo(campaignName string) {
	camp, err := ap.campaignService.Get(campaignName)
	if err != nil {
		ap.writeStdErr("campaign service get error: %v", err)
		return
	}

	totalSales, err := ap.orderService.GetProductTotalSales(camp.ProductCode)
	if err != nil {
		ap.writeStdErr("order service get product total sales error: %v", err)
		return
	}

	campFinishDate := camp.CreatedAt.Add(time.Duration(camp.Duration) * time.Hour)
	campaignStatus := "Ended"
	if time.Now().Before(campFinishDate) {
		campaignStatus = "Active"
	}

	turnOver, err := ap.orderService.GetProductTurnOver(camp.ProductCode)
	if err != nil {
		ap.writeStdErr("order service get product turnover error: %v", err)
		return
	}

	averagePrice, err := ap.orderService.GetProductAveragePrice(camp.ProductCode)
	if err != nil {
		ap.writeStdErr("order service get product average price error: %v", err)
		return
	}

	ap.writeStdOut(
		"Campaign %s info; Status %s, Target Sales %d, Total Sales %d, Turnover %d, Average Item Price %d",
		camp.Name, campaignStatus, camp.TargetSalesCount, totalSales, turnOver, averagePrice,
	)
}

func (ap *CampaignAutoPricing) GetProductInfo(productCode string) {
	camp, err := ap.campaignService.GetFromProductCode(productCode)
	if err != nil {
		ap.writeStdErr("campaign service get from product code error: %v", err)
		return
	}
	prod, err := ap.productService.Get(productCode)
	if err != nil {
		ap.writeStdErr("product service get error: %v", err)
		return
	}
	totalSales, err := ap.orderService.GetProductTotalSales(productCode)
	if err != nil {
		ap.writeStdErr("order service get product total sales error: %v", err)
		return
	}
	currentPrice, err := ap.priceDetermination(
		prod.Price,
		camp.PriceManipulationLimit,
		totalSales,
		camp.TargetSalesCount,
		camp.CreatedAt,
		camp.Duration,
	)
	if err != nil {
		ap.writeStdErr("price determination error: %v", err)
		return
	}

	ap.writeStdOut("Product %s info; price %d, stock:%d", prod.ProductCode, currentPrice, prod.Stock-totalSales)
}

func (ap *CampaignAutoPricing) IncreaseTime(hour int) {
	time.Increase(hour)
}

func (ap *CampaignAutoPricing) writeStdOut(format string, a ...any) {
	format = fmt.Sprintf("%s\n", format)
	fmt.Fprintf(ap.stdOutput, format, a...)
}

func (ap *CampaignAutoPricing) writeStdErr(format string, a ...any) {
	format = fmt.Sprintf("%s\n", format)
	fmt.Fprintf(ap.errOutput, format, a...)
}
