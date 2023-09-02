package order

import (
	"time"
)

// TimeNow for Order time manipulation
var TimeNow = time.Now

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
		Price:       price,
		CreatedAt:   TimeNow(),
	}, nil
}
