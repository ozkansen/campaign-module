package memory

import (
	"sync"

	"github.com/ozkansen/campaign-module/domain/product"
)

var _ product.Repository = (*MemoryProductRepository)(nil)

type MemoryProductRepository struct {
	products map[string]*product.Product
	mutex    sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[string]*product.Product),
	}
}

func (m *MemoryProductRepository) Create(prod *product.Product) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, exists := m.products[prod.ProductCode]
	if exists {
		return product.ErrProductAlreadyExists
	}
	m.products[prod.ProductCode] = prod
	return nil
}

func (m *MemoryProductRepository) Get(productCode string) (*product.Product, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	prod, exists := m.products[productCode]
	if !exists {
		return nil, product.ErrProductNotFound
	}
	return prod, nil
}
