package memory

import (
	"sync"

	"github.com/ozkansen/campaign-module/domain/order"
)

var _ order.Repository = (*MemoryOrderRepository)(nil)

type MemoryOrderRepository struct {
	orders map[string][]*order.Order
	mutex  sync.Mutex
}

func New() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make(map[string][]*order.Order),
	}
}

func (m *MemoryOrderRepository) Create(ord *order.Order) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	productOrders, exists := m.orders[ord.ProductCode]
	if !exists || productOrders == nil {
		productOrders = make([]*order.Order, 0)
		m.orders[ord.ProductCode] = productOrders
	}
	m.orders[ord.ProductCode] = append(productOrders, ord)
	return nil
}

func (m *MemoryOrderRepository) Get(productCode string) ([]*order.Order, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	productOrders, exists := m.orders[productCode]
	if !exists || productOrders == nil {
		return nil, order.ErrOrdersNotAvailable
	}
	return productOrders, nil
}
