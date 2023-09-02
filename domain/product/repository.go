package product

type Repository interface {
	Create(prod *Product) error
	Get(productCode string) (*Product, error)
}
