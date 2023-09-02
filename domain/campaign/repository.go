package campaign

type Repository interface {
	Create(camp *Campaign) error
	Get(name string) (*Campaign, error)
	GetFromProductCode(productCode string) (*Campaign, error)
}
