package service

import (
	"ex-di-interfaces/model"
)

type ProductRepository interface {
	FindBy(id int) *model.Product
	FindAll() []model.Product
	Save(newProduct model.Product) *model.Product
}

type DefaultProductService struct {
	Repo ProductRepository
}

func (ps DefaultProductService) GetAllProducts() []model.Product {
	return ps.Repo.FindAll()
}

func (ps DefaultProductService) GetProductById(id int) *model.Product {
	return ps.Repo.FindBy(id)
}

func (ps DefaultProductService) AddProduct(name string, category model.Category, price float32) *model.Product {
	newProduct := model.Product{Name: name, Category: category, Price: price}

	return ps.Repo.Save(newProduct)
}

func NewProductService(repo ProductRepository) DefaultProductService {
	return DefaultProductService{
		Repo: repo,
	}
}
