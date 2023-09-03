package order

import (
	"errors"

	"github.com/ozkansen/campaign-module/domain/order"
	"github.com/ozkansen/campaign-module/domain/order/memory"
)

type (
	Configuration func(os *OrderService) error
	newOrder      func(productCode string, quantity int, price int64) (*order.Order, error)
)

type OrderService struct {
	orders   order.Repository
	newOrder newOrder
}

func New(cfgs ...Configuration) (*OrderService, error) {
	os := &OrderService{
		newOrder: order.NewOrder,
	}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithOrderRepository(or order.Repository) Configuration {
	return func(os *OrderService) error {
		os.orders = or
		return nil
	}
}

func WithMemoryOrderRepository() Configuration {
	or := memory.New()
	return WithOrderRepository(or)
}

func (os *OrderService) Create(productCode string, quantity int, price int64) error {
	ord, err := os.newOrder(productCode, quantity, price)
	if err != nil {
		return err
	}
	return os.orders.Create(ord)
}

func (os *OrderService) GetProductTotalSales(productCode string) (int, error) {
	orders, err := os.orders.Get(productCode)
	if err != nil {
		if errors.Is(err, order.ErrOrdersNotAvailable) {
			return 0, nil
		}
		return 0, err
	}

	var salesTotalQuantity int
	for _, ord := range orders {
		salesTotalQuantity += ord.Quantity
	}
	return salesTotalQuantity, nil
}
