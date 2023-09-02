package campaign

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ozkansen/campaign-module/domain/campaign"
)

type campaignRepositoryStubs struct {
	createRetErr error
	getRetVal    *campaign.Campaign
	getRetErr    error
}

func (cr *campaignRepositoryStubs) Create(camp *campaign.Campaign) error {
	return cr.createRetErr
}
func (cr *campaignRepositoryStubs) Get(name string) (*campaign.Campaign, error) {
	return cr.getRetVal, cr.getRetErr
}

func newCampaignFuncStub(c *campaign.Campaign, err error) newCampaign {
	return func(name, productCode string, duration, priceManipulationLimit, targetSalesCount int) (*campaign.Campaign, error) {
		return c, err
	}
}

func TestCampaignService_Create(t *testing.T) {
	type fields struct {
		campaigns   campaign.Repository
		newCampaign newCampaign
	}
	type args struct {
		name                   string
		productCode            string
		duration               int
		priceManipulationLimit int
		targetSalesCount       int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create a valid campaign",
			fields: fields{
				campaigns:   &campaignRepositoryStubs{createRetErr: nil},
				newCampaign: newCampaignFuncStub(&campaign.Campaign{}, nil),
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "invalid campaign create",
			fields: fields{
				campaigns:   &campaignRepositoryStubs{createRetErr: nil},
				newCampaign: newCampaignFuncStub(&campaign.Campaign{}, errors.New("invalid")),
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "repository error",
			fields: fields{
				campaigns:   &campaignRepositoryStubs{createRetErr: errors.New("error")},
				newCampaign: newCampaignFuncStub(&campaign.Campaign{}, nil),
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CampaignService{
				campaigns:   tt.fields.campaigns,
				newCampaign: tt.fields.newCampaign,
			}
			if err := c.Create(tt.args.name, tt.args.productCode, tt.args.duration, tt.args.priceManipulationLimit, tt.args.targetSalesCount); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCampaignService_Get(t *testing.T) {
	type fields struct {
		campaigns   campaign.Repository
		newCampaign newCampaign
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
			name: "get valid campaign",
			fields: fields{
				campaigns: &campaignRepositoryStubs{getRetVal: &campaign.Campaign{}, getRetErr: nil},
			},
			args:    args{},
			want:    &campaign.Campaign{},
			wantErr: false,
		},
		{
			name: "repository error",
			fields: fields{
				campaigns: &campaignRepositoryStubs{getRetVal: nil, getRetErr: errors.New("error")},
			},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CampaignService{
				campaigns:   tt.fields.campaigns,
				newCampaign: tt.fields.newCampaign,
			}
			got, err := cs.Get(tt.args.name)
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
