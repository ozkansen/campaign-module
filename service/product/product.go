package product

import (
	"github.com/ozkansen/campaign-module/domain/product"
	"github.com/ozkansen/campaign-module/domain/product/memory"
)

type (
	Configuration func(ps *ProductService) error
	newProduct    func(productCode string, price int64, stock int) (*product.Product, error)
)

type ProductService struct {
	products   product.Repository
	newProduct newProduct
}

func New(cfgs ...Configuration) (*ProductService, error) {
	ps := &ProductService{
		newProduct: product.NewProduct,
	}
	for _, cfg := range cfgs {
		err := cfg(ps)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

func WithProductRepository(pr product.Repository) Configuration {
	return func(ps *ProductService) error {
		ps.products = pr
		return nil
	}
}

func WithMemoryProductRepository() Configuration {
	pr := memory.New()
	return WithProductRepository(pr)
}

func (ps *ProductService) Create(productCode string, price int64, stock int) error {
	newProd, err := ps.newProduct(productCode, price, stock)
	if err != nil {
		return err
	}
	err = ps.products.Create(newProd)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) Get(productCode string) (*product.Product, error) {
	return ps.products.Get(productCode)
}
