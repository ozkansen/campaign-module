package order

import (
	"github.com/ozkansen/campaign-module/pkg/time"
)

type Order struct {
	ProductCode string
	Quantity    int
	Price       int64
	CreatedAt   time.Time
}

func NewOrder(productCode string, quantity int, price int64) (*Order, error) {
	if productCode == "" || quantity <= 0 {
		return nil, ErrInvalidValue
	}
	return &Order{
		ProductCode: productCode,
		Quantity:    quantity,
		Price:       price * int64(quantity),
		CreatedAt:   time.Now(),
	}, nil
}
