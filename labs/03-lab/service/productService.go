package service

import (
	"03-lab/model"
	"03-lab/repository"
)
var  repo repository.InMemoryProductRepo = repository.NewInMemoryProductRepo()
func GetAllProducts() []model.Product {
	return repo.FindAll()
}

func GetProductById(id int) *model.Product {
	return repo.FindBy(id)
}

func AddProduct(name string, category model.Category, price float32) *model.Product {
	newProduct := model.Product{Name: name, Category: category, Price: price}

	return repo.Save(newProduct)
}
