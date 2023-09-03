package campaign

import (
	"reflect"
	"testing"

	"github.com/ozkansen/campaign-module/pkg/time"
)

func TestNewCampaign(t *testing.T) {
	time.Set(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))

	type args struct {
		name                   string
		productCode            string
		duration               int
		priceManipulationLimit int
		targetSalesCount       int
	}
	tests := []struct {
		name    string
		args    args
		want    *Campaign
		wantErr bool
	}{
		{
			name:    "create a valid campaign object",
			args:    args{"example", "P1", 1, 20, 10},
			want:    &Campaign{"example", "P1", 1, 20, 10, time.Now()},
			wantErr: false,
		},
		{
			name:    "invalid name check test",
			args:    args{"", "P1", 1, 20, 10},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid product code check test",
			args:    args{"example", "", 1, 20, 10},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid duration check test",
			args:    args{"example", "P1", 0, 20, 10},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid min price manipulation limit check test",
			args:    args{"example", "P1", 1, 0, 10},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid max price manipulation limit check test",
			args:    args{"example", "P1", 1, 100, 10},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid target sales count check test",
			args:    args{"example", "P1", 1, 100, 0},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCampaign(tt.args.name, tt.args.productCode, tt.args.duration, tt.args.priceManipulationLimit, tt.args.targetSalesCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCampaign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCampaign() got = %v, want %v", got, tt.want)
			}
		})
	}
}
