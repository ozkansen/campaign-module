package memory

import (
	"reflect"
	"testing"
	"time"

	"github.com/ozkansen/campaign-module/domain/campaign"
)

func TestMemoryCampaignRepository_Create(t *testing.T) {
	TimeNow := func() time.Time {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	type fields struct {
		campaigns map[string]*campaign.Campaign
	}
	type args struct {
		camp *campaign.Campaign
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "create a valid campaign",
			fields:  fields{make(map[string]*campaign.Campaign)},
			args:    args{camp: &campaign.Campaign{Name: "C1", ProductCode: "P1", Duration: 1, PriceManipulationLimit: 10, TargetSalesCount: 100, CreatedAt: TimeNow()}},
			wantErr: false,
		},
		{
			name:    "invalid campaign already exists, same product code",
			fields:  fields{map[string]*campaign.Campaign{"C1": {Name: "C1", ProductCode: "P1", Duration: 1, PriceManipulationLimit: 10, TargetSalesCount: 100, CreatedAt: time.Now()}}},
			wantErr: true,
			args:    args{camp: &campaign.Campaign{Name: "C2", ProductCode: "P1", Duration: 1, PriceManipulationLimit: 10, TargetSalesCount: 100, CreatedAt: TimeNow()}},
		},
		{
			name:    "invalid campaign already exists, same campaign name",
			fields:  fields{map[string]*campaign.Campaign{"C1": {Name: "C1", ProductCode: "P1", Duration: 1, PriceManipulationLimit: 10, TargetSalesCount: 100, CreatedAt: time.Now()}}},
			wantErr: true,
			args:    args{camp: &campaign.Campaign{Name: "C1", ProductCode: "P2", Duration: 1, PriceManipulationLimit: 10, TargetSalesCount: 100, CreatedAt: TimeNow()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCampaignRepository{
				campaigns: tt.fields.campaigns,
			}
			if err := m.Create(tt.args.camp); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryCampaignRepository_Get(t *testing.T) {
	type fields struct {
		campaigns map[string]*campaign.Campaign
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *campaign.Campaign
		wantErr bool
	}{
		{
			name:    "get campaign",
			fields:  fields{map[string]*campaign.Campaign{"C1": {Name: "C1"}}},
			args:    args{name: "C1"},
			want:    &campaign.Campaign{Name: "C1"},
			wantErr: false,
		},
		{
			name:    "campaign not found",
			fields:  fields{make(map[string]*campaign.Campaign)},
			args:    args{name: "C1"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCampaignRepository{
				campaigns: tt.fields.campaigns,
			}
			got, err := m.Get(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
