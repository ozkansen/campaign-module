package memory

import (
	"reflect"
	"testing"

	"github.com/ozkansen/campaign-module/domain/product"
)

func TestMemoryProductRepository_Create(t *testing.T) {
	type fields struct {
		products map[string]*product.Product
	}
	type args struct {
		prod *product.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "create a valid product",
			fields:  fields{products: make(map[string]*product.Product)},
			args:    args{&product.Product{}},
			wantErr: false,
		},
		{
			name:    "product already exists",
			fields:  fields{products: map[string]*product.Product{"P1": {ProductCode: "P1"}}},
			args:    args{prod: &product.Product{ProductCode: "P1"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryProductRepository{
				products: tt.fields.products,
			}
			if err := m.Create(tt.args.prod); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {
	type fields struct {
		products map[string]*product.Product
	}
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *product.Product
		wantErr bool
	}{
		{
			name:    "get a valid product",
			fields:  fields{products: map[string]*product.Product{"P1": &product.Product{ProductCode: "P1"}, "P2": {ProductCode: "P2"}}},
			args:    args{productCode: "P1"},
			want:    &product.Product{ProductCode: "P1"},
			wantErr: false,
		},
		{
			name:    "product not found",
			fields:  fields{products: map[string]*product.Product{"P1": &product.Product{ProductCode: "P1"}, "P2": {ProductCode: "P2"}}},
			args:    args{productCode: "P3"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryProductRepository{
				products: tt.fields.products,
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
