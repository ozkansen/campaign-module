package order

import (
	"errors"
	"testing"

	"github.com/ozkansen/campaign-module/domain/order"
)

type orderRepositoryStubs struct {
	createRetErr error
	getRetVal    []*order.Order
	getRetErr    error
}

func (or *orderRepositoryStubs) Create(ord *order.Order) error {
	return or.createRetErr
}

func (or *orderRepositoryStubs) Get(productCode string) ([]*order.Order, error) {
	return or.getRetVal, or.getRetErr
}

func newOrderFuncStub(o *order.Order, err error) newOrder {
	return func(productCode string, quantity int) (*order.Order, error) {
		return o, err
	}
}

func TestOrderService_Create(t *testing.T) {
	type fields struct {
		orders   order.Repository
		newOrder newOrder
	}
	type args struct {
		productCode string
		quantity    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create a valid order",
			fields: fields{
				orders:   &orderRepositoryStubs{},
				newOrder: newOrderFuncStub(&order.Order{ProductCode: "P1", Quantity: 1}, nil),
			},
			args:    args{productCode: "P1", quantity: 1},
			wantErr: false,
		},
		{
			name: "invalid order create",
			fields: fields{
				orders:   &orderRepositoryStubs{},
				newOrder: newOrderFuncStub(nil, errors.New("error")),
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "repository error",
			fields: fields{
				orders:   &orderRepositoryStubs{createRetErr: errors.New("error")},
				newOrder: newOrderFuncStub(nil, nil),
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os := &OrderService{
				orders:   tt.fields.orders,
				newOrder: tt.fields.newOrder,
			}
			if err := os.Create(tt.args.productCode, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderService_GetProductTotalSales(t *testing.T) {
	type fields struct {
		orders order.Repository
	}
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "calc valid total sales product",
			fields: fields{
				orders: &orderRepositoryStubs{getRetVal: []*order.Order{{
					ProductCode: "P1",
					Quantity:    10,
				}, {
					ProductCode: "P1",
					Quantity:    20,
				}, {
					ProductCode: "P1",
					Quantity:    40,
				}}},
			},
			args: args{
				productCode: "P1",
			},
			want:    70,
			wantErr: false,
		},
		{
			name:    "repository error",
			fields:  fields{orders: &orderRepositoryStubs{getRetErr: errors.New("error")}},
			args:    args{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os := &OrderService{
				orders: tt.fields.orders,
			}
			got, err := os.GetProductTotalSales(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductTotalSales() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetProductTotalSales() got = %v, want %v", got, tt.want)
			}
		})
	}
}
