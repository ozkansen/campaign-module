package campaign_auto_pricing

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/ozkansen/campaign-module/domain/campaign"
	"github.com/ozkansen/campaign-module/domain/product"
	"github.com/ozkansen/campaign-module/pkg/time"
)

type productServiceStubs struct {
	createRetErr error
	getRetVal    *product.Product
	getRetErr    error
}

func (p *productServiceStubs) Create(productCode string, price int64, stock int) error {
	return p.createRetErr
}

func (p *productServiceStubs) Get(productCode string) (*product.Product, error) {
	return p.getRetVal, p.getRetErr
}

type orderServiceStubs struct {
	createRetErr                 error
	getProductTotalSalesRetVal   int
	getProductTotalSalesRetErr   error
	getProductTurnOverRetVal     int64
	getProductTurnOverRetErr     error
	getProductAveragePriceRetVal int64
	getProductAveragePriceRetErr error
}

func (o *orderServiceStubs) Create(productCode string, quantity int, price int64) error {
	return o.createRetErr
}

func (o *orderServiceStubs) GetProductTotalSales(productCode string) (int, error) {
	return o.getProductTotalSalesRetVal, o.getProductTotalSalesRetErr
}

func (o *orderServiceStubs) GetProductTurnOver(productCode string) (int64, error) {
	return o.getProductTurnOverRetVal, o.getProductTurnOverRetErr
}

func (o *orderServiceStubs) GetProductAveragePrice(productCode string) (int64, error) {
	return o.getProductAveragePriceRetVal, o.getProductAveragePriceRetErr
}

type campaignServiceStubs struct {
	createRetErr             error
	getRetVal                *campaign.Campaign
	getRetErr                error
	getFromProductCodeRetVal *campaign.Campaign
	getFromProductCodeRetErr error
}

func (c *campaignServiceStubs) Create(name string, productCode string, duration int, priceManipulationLimit int, targetSalesCount int) error {
	return c.createRetErr
}

func (c *campaignServiceStubs) Get(name string) (*campaign.Campaign, error) {
	return c.getRetVal, c.getRetErr
}

func (c *campaignServiceStubs) GetFromProductCode(productCode string) (*campaign.Campaign, error) {
	return c.getFromProductCodeRetVal, c.getFromProductCodeRetErr
}

func TestCampaignAutoPricing_CreateCampaign(t *testing.T) {
	productService := &productServiceStubs{}
	orderService := &orderServiceStubs{}
	campaignService := &campaignServiceStubs{}
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	cps := &CampaignAutoPricing{
		productService:  productService,
		orderService:    orderService,
		campaignService: campaignService,
		stdOutput:       stdOut,
		errOutput:       stdErr,
	}

	t.Run("create campaign valid test", func(t *testing.T) {
		cps.CreateCampaign("C1", "P1", 1, 10, 100)
		output, err := io.ReadAll(stdOut)
		if err != nil {
			t.Error("stdout reading error")
		}
		if !bytes.Contains(output, []byte("Campaign created; name C1, product P1, duration 1, limit 10, target sales count 100")) {
			t.Error("output not as expected", string(output))
		}
		stdOut.Reset()
	})

	t.Run("product service error", func(t *testing.T) {
		productService.getRetErr = errors.New("error")
		cps.CreateCampaign("C1", "P1", 1, 10, 100)
		output, err := io.ReadAll(stdErr)
		if err != nil {
			t.Error("stderr reading error")
		}
		if !bytes.Contains(output, []byte("product service get: error")) {
			t.Error("output not as expected", string(output))
		}
		productService.getRetErr = nil
		stdErr.Reset()
	})

	t.Run("campaign service error", func(t *testing.T) {
		campaignService.createRetErr = errors.New("error")
		cps.CreateCampaign("C1", "P1", 1, 10, 100)
		output, err := io.ReadAll(stdErr)
		if err != nil {
			t.Error("stderr reading error")
		}
		if !bytes.Contains(output, []byte("campaign service create: error")) {
			t.Error("output not as expected", string(output))
		}
		campaignService.createRetErr = nil
		stdErr.Reset()
	})
}

func TestCampaignAutoPricing_CreateProduct(t *testing.T) {
	productService := &productServiceStubs{}
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cps := &CampaignAutoPricing{
		productService: productService,
		stdOutput:      stdOut,
		errOutput:      stdErr,
	}

	t.Run("create product valid test", func(t *testing.T) {
		cps.CreateProduct("P1", 100, 250)
		output, err := io.ReadAll(stdOut)
		if err != nil {
			t.Error("stdout reading error")
		}
		if !bytes.Contains(output, []byte("Product created; code P1, price 100, stock 250")) {
			t.Error("output not as expected", string(output))
		}
		stdOut.Reset()
	})

	t.Run("product service error", func(t *testing.T) {
		productService.createRetErr = errors.New("error")
		cps.CreateProduct("P1", 100, 250)
		output, err := io.ReadAll(stdErr)
		if err != nil {
			t.Error("stdErr reading error")
		}
		if !bytes.Contains(output, []byte("product creating error: error")) {
			t.Error("output not as expected", string(output))
		}
		stdErr.Reset()
	})
}

func TestCampaignAutoPricing_IncreaseTime(t *testing.T) {
	type args struct {
		hour int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base datetime change valid test",
			args: args{1},
		},
		{
			name: "base datetime change valid test 2",
			args: args{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &CampaignAutoPricing{}
			old := time.Now().Hour()
			ap.IncreaseTime(tt.args.hour)
			renew := time.Now().Hour()
			if old == renew {
				t.Error("wrong time error")
			}
			ap.IncreaseTime(-1)
		})
	}
}
