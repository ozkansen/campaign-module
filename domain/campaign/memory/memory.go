package memory

import (
	"sync"

	"github.com/ozkansen/campaign-module/domain/campaign"
)

var _ campaign.Repository = (*MemoryCampaignRepository)(nil)

type MemoryCampaignRepository struct {
	campaigns map[string]*campaign.Campaign
	mutex     sync.Mutex
}

func New() *MemoryCampaignRepository {
	return &MemoryCampaignRepository{
		campaigns: make(map[string]*campaign.Campaign),
	}
}

func (m *MemoryCampaignRepository) Create(camp *campaign.Campaign) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, exists := m.campaigns[camp.Name]
	if exists {
		return campaign.ErrCampaignAlreadyExist
	}
	// multiple campaigns on the same product
	for n := range m.campaigns {
		c, _ := m.campaigns[n]
		if c.ProductCode == camp.ProductCode {
			return campaign.ErrCampaignAlreadyExist
		}
	}
	m.campaigns[camp.Name] = camp
	return nil
}

func (m *MemoryCampaignRepository) Get(name string) (*campaign.Campaign, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	camp, exists := m.campaigns[name]
	if !exists {
		return nil, campaign.ErrCampaignNotFound
	}
	return camp, nil
}

func (m *MemoryCampaignRepository) GetFromProductCode(productCode string) (*campaign.Campaign, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for name := range m.campaigns {
		camp := m.campaigns[name]
		if camp.ProductCode == productCode {
			return camp, nil
		}
	}
	return nil, campaign.ErrCampaignNotFound
}
