package product

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ozkansen/campaign-module/domain/product"
)

type productRepositoryStubs struct {
	createRetErr error
	getRetVal    *product.Product
	getRetErr    error
}

func (pr *productRepositoryStubs) Create(prod *product.Product) error {
	return pr.createRetErr
}

func (pr *productRepositoryStubs) Get(productCode string) (*product.Product, error) {
	return pr.getRetVal, pr.getRetErr
}

func newProductFuncStub(p *product.Product, err error) newProduct {
	return func(productCode string, price int64, stock int) (*product.Product, error) {
		return p, err
	}
}

func TestProductService_Create(t *testing.T) {
	type fields struct {
		products   product.Repository
		newProduct newProduct
	}
	type args struct {
		productCode string
		price       int64
		stock       int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create a valid product",
			fields: fields{
				products:   &productRepositoryStubs{createRetErr: nil},
				newProduct: newProductFuncStub(&product.Product{ProductCode: "P1"}, nil),
			},
			args:    args{productCode: "P1"},
			wantErr: false,
		},
		{
			name: "invalid product create",
			fields: fields{
				products:   &productRepositoryStubs{},
				newProduct: newProductFuncStub(&product.Product{ProductCode: "P1"}, errors.New("error")),
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "repository error",
			fields: fields{
				products:   &productRepositoryStubs{createRetErr: errors.New("error")},
				newProduct: newProductFuncStub(nil, nil),
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{
				products:   tt.fields.products,
				newProduct: tt.fields.newProduct,
			}
			if err := ps.Create(tt.args.productCode, tt.args.price, tt.args.stock); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductService_Get(t *testing.T) {
	type fields struct {
		products product.Repository
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
			name: "get a valid product",
			fields: fields{
				products: &productRepositoryStubs{getRetVal: &product.Product{ProductCode: "P1"}},
			},
			args:    args{productCode: "P1"},
			want:    &product.Product{ProductCode: "P1"},
			wantErr: false,
		},
		{
			name: "invalid product",
			fields: fields{
				products: &productRepositoryStubs{getRetVal: nil, getRetErr: errors.New("error")},
			},
			args:    args{productCode: "P2"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{
				products: tt.fields.products,
			}
			got, err := ps.Get(tt.args.productCode)
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
