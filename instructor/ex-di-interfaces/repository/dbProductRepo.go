package repository

import "ex-di-interfaces/model"

type DbProductRepo struct {
	products map[int]model.Product
}

func (r DbProductRepo) FindBy(id int) *model.Product {
	// will be replaced by actual SQL
	if p, ok := r.products[id]; ok {
		return &p
	}
	return nil
}

func (r DbProductRepo) FindAll() []model.Product {
	v := make([]model.Product, 0)
	for _, p := range r.products {
		v = append(v, p)
	}
	return v
}

func (r DbProductRepo) Save(newProduct model.Product) *model.Product {
	id := len(r.products) + 1
	newProduct.Id = id
	r.products[id] = newProduct
	return &newProduct
}

func (r DbProductRepo) UpdateProduct(productToBeUpdated model.Product) bool {
	if product := r.FindBy(productToBeUpdated.Id); product != nil {
		r.products[productToBeUpdated.Id] = productToBeUpdated
		return true
	}
	return false
}

func NewDbProductRepo() DbProductRepo {
	db := DbProductRepo{
		products: map[int]model.Product{
			1: model.NewProduct(1, "A can of green peas", model.FOOD, 100.0),
			2: model.NewProduct(2, "Learning Java Streams", model.BOOKS, 55.0),
			3: model.NewProduct(3, "Smoothies", model.BEVERAGE, 41.46),
			4: model.NewProduct(4, "The Product Book", model.BOOKS, 697.57),
			5: model.NewProduct(5, "Plain drinking water", model.BEVERAGE, 366.90),
			6: model.NewProduct(6, "Action figures", model.TOYS, 95.50),
			7: model.NewProduct(7, "A can of green olives", model.FOOD, 302.19),
			8: model.NewProduct(8, "Creative toys", model.TOYS, 295.37),
			9: model.NewProduct(9, "Frozen Food", model.FOOD, 534.64),
		},
	}
	return db
}
