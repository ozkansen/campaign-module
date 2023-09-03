package campaign_auto_pricing

import (
	"testing"

	"github.com/ozkansen/campaign-module/pkg/time"
)

func TestAlwaysMinPrice(t *testing.T) {
	type args struct {
		basePrice              int64
		priceManipulationLimit int
		totalSales             int
		targetSales            int
		startedDate            time.Time
		duration               int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				basePrice:              100,
				priceManipulationLimit: 20,
				totalSales:             0,
				targetSales:            0,
				startedDate:            time.Now(),
				duration:               1,
			},
			want:    80,
			wantErr: false,
		},
		{
			name: "expired test",
			args: args{
				basePrice:              100,
				priceManipulationLimit: 20,
				totalSales:             0,
				targetSales:            0,
				startedDate:            time.Now().Add(-2 * time.Hour),
				duration:               1,
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AlwaysMinPrice(tt.args.basePrice, tt.args.priceManipulationLimit, tt.args.totalSales, tt.args.targetSales, tt.args.startedDate, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlwaysMinPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AlwaysMinPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineerPricing(t *testing.T) {
	type args struct {
		basePrice              int64
		priceManipulationLimit int
		totalSales             int
		targetSales            int
		startedDate            time.Time
		duration               int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				basePrice:              100,
				priceManipulationLimit: 20,
				totalSales:             0,
				targetSales:            100,
				startedDate:            time.Now(),
				duration:               10,
			},
			want:    80,
			wantErr: false,
		},
		{
			name: "valid test 2",
			args: args{
				basePrice:              100,
				priceManipulationLimit: 20,
				totalSales:             50,
				targetSales:            100,
				startedDate:            time.Now(),
				duration:               10,
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "valid test 3",
			args: args{
				basePrice:              100,
				priceManipulationLimit: 20,
				totalSales:             100,
				targetSales:            100,
				startedDate:            time.Now(),
				duration:               10,
			},
			want:    120,
			wantErr: false,
		},
		{
			name: "expired campaign",
			args: args{
				basePrice:              100,
				priceManipulationLimit: 20,
				totalSales:             20,
				targetSales:            100,
				startedDate:            time.Now(),
				duration:               -10,
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LineerPricing(tt.args.basePrice, tt.args.priceManipulationLimit, tt.args.totalSales, tt.args.targetSales, tt.args.startedDate, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineerPricing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LineerPricing() got = %v, want %v", got, tt.want)
			}
		})
	}
}
