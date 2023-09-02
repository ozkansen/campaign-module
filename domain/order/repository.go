package order

type Repository interface {
	Create(ord *Order) error
	Get(productCode string) ([]*Order, error)
}
