package dto

type Product struct {
	ID string `json:id`
	Name string `json:name`
	Price float64 `json:price`
	Status string `json:status`
}

func newProduct() * Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product, error) {
	if(p.ID != "") {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = P.Status
	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}