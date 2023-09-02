package campaign

import (
	"github.com/ozkansen/campaign-module/domain/campaign"
	"github.com/ozkansen/campaign-module/domain/campaign/memory"
)

type (
	Configuration func(cs *CampaignService) error
	newCampaign   func(name, productCode string, duration, priceManipulationLimit, targetSalesCount int) (*campaign.Campaign, error)
)

type CampaignService struct {
	campaigns   campaign.Repository
	newCampaign newCampaign
}

func New(cfgs ...Configuration) (*CampaignService, error) {
	cs := &CampaignService{
		newCampaign: campaign.NewCampaign,
	}
	for _, cfg := range cfgs {
		err := cfg(cs)
		if err != nil {
			return nil, err
		}
	}
	return cs, nil
}

func WithCampaignRepository(cr campaign.Repository) Configuration {
	return func(cs *CampaignService) error {
		cs.campaigns = cr
		return nil
	}
}

func WithMemoryCampaignRepository() Configuration {
	cr := memory.New()
	return WithCampaignRepository(cr)
}

func (cs *CampaignService) Create(name, productCode string, duration, priceManipulationLimit int, targetSalesCount int) error {
	camp, err := cs.newCampaign(name, productCode, duration, priceManipulationLimit, targetSalesCount)
	if err != nil {
		return err
	}
	err = cs.campaigns.Create(camp)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CampaignService) Get(name string) (*campaign.Campaign, error) {
	return cs.campaigns.Get(name)
}

func (cs *CampaignService) GetFromProductCode(productCode string) (*campaign.Campaign, error) {
	return cs.campaigns.GetFromProductCode(productCode)
}
