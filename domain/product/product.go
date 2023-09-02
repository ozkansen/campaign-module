package product

type Product struct {
	ProductCode string
	Price       int64
	Stock       int
}

func NewProduct(productCode string, price int64, stock int) (*Product, error) {
	if productCode == "" || price <= 0 || stock <= 0 {
		return nil, ErrInvalidValue
	}
	return &Product{
		ProductCode: productCode,
		Price:       price,
		Stock:       stock,
	}, nil
}
