package order

import (
	"time"
)

// TimeNow for Order time manipulation
var TimeNow = time.Now

type Order struct {
	ProductCode string
	Quantity    int
	CreatedAt   time.Time
}

func NewOrder(productCode string, quantity int) (*Order, error) {
	if productCode == "" || quantity <= 0 {
		return nil, ErrInvalidValue
	}
	return &Order{
		ProductCode: productCode,
		Quantity:    quantity,
		CreatedAt:   TimeNow(),
	}, nil
}
