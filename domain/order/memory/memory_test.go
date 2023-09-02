package memory

import (
	"reflect"
	"testing"
	"time"

	"github.com/ozkansen/campaign-module/domain/order"
)

func TestMemoryOrderRepository_Create(t *testing.T) {
	type fields struct {
		orders map[string][]*order.Order
	}
	type args struct {
		ord *order.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "create a valid order",
			fields:  fields{orders: map[string][]*order.Order{}},
			args:    args{ord: &order.Order{ProductCode: "P1"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryOrderRepository{
				orders: tt.fields.orders,
			}
			if err := m.Create(tt.args.ord); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryOrderRepository_Get(t *testing.T) {
	order.TimeNow = func() time.Time {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	type fields struct {
		orders map[string][]*order.Order
	}
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*order.Order
		wantErr bool
	}{
		{
			name:    "get a valid order",
			fields:  fields{orders: map[string][]*order.Order{"P1": {{ProductCode: "P1"}}}},
			args:    args{productCode: "P1"},
			want:    []*order.Order{{ProductCode: "P1"}},
			wantErr: false,
		},
		{
			name:    "orders not available",
			fields:  fields{orders: map[string][]*order.Order{"P1": {{ProductCode: "P1"}}}},
			args:    args{productCode: "P2"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "orders not available nil check",
			fields:  fields{orders: map[string][]*order.Order{"P1": nil}},
			args:    args{productCode: "P1"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryOrderRepository{
				orders: tt.fields.orders,
			}
			got, err := m.Get(tt.args.productCode)
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
