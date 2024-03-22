package service

import (
	mocks "ex-di-interfaces/mocks/service"
	"ex-di-interfaces/model"
	"testing"
)

// type MockProductRepository struct {
// }

// func (m MockProductRepository) FindBy(id int) *model.Product {
// 	return nil
// }
// func (m MockProductRepository) FindAll() []model.Product {
// 	return []model.Product{{1, "name", model.BEVERAGE, 100.02}}
// }
// func (m MockProductRepository) Save(newProduct model.Product) *model.Product {
// 	return nil
// }

// STATE based tests
// Interaction based test

func Test_Get_all_products(t *testing.T) {
	// What do we want to test? I want to test the interaction
	// when calling service.GetAllProducts() a call to repo.FindAll() is made
	// think in terms of responsibility

	// ARRANGE
	mockRepo := mocks.NewProductRepository(t)
	ps := NewProductService(mockRepo)

	// ACT & ASSERT I am also setting it as an expected that a call to FindAll function will be made
	mockRepo.On("FindAll").Return([]model.Product{{1, "name", model.BEVERAGE, 100.02}})

	//ACT
	ps.GetAllProducts()
}

func Test_add_products(t *testing.T) {
	// What do we want to test? I want to test the interaction
	// when calling service.AddProduct() a call to repo.Save(Product) is made

	// ARRANGE
	mockRepo := mocks.NewProductRepository(t)
	ps := NewProductService(mockRepo)

	// ACT & ASSERT I am also setting it as an expected that a call to FindAll function will be made
	newProduct := model.NewProduct(0, "test product", model.BOOKS, 123.45)

	mockRepo.On("Save", newProduct).Return(&newProduct)

	//ACT
	ps.AddProduct("test product", model.BOOKS, 123.45)
}

func Test_addition_functionality(t *testing.T) {
	// what do I want to test?
	// AAA -> Arrange, Act, Assert

	// ARRANGE
	expected := 300

	// ACT
	actual := add(100, 200)
	// ASSERT
	if actual != expected {
		t.Fail()
	}

}

// Function under test
func add(x, y int) int {
	return x + y
}

/*
	func TestAddProduct(t *testing.T) {
		newProduct := AddProduct("new test product", model.TOYS, 100.23)

		if newProduct.Name != "new test product" {
			t.Errorf("Not the same product ")
		}

		if len(GetAllProducts()) != 6 {
			t.Errorf("Expected product list should have 6 products")
		}
	}

func TestGetAllProducts(t *testing.T) {

		if got := GetAllProducts(); got != nil {
			if len(got) != 5 {
				t.Errorf("Expected products list should have 5 products")
			}
		}
	}

func TestGetProductById(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want *model.Product
	}{
		{
			"With product Id 1",
			1,
			&model.Product{Id: 1, Name: "InMem - Baked goods", Category: model.FOOD, Price: 100.0},
		},
		{
			"With product Id 2",
			2,
			&model.Product{Id: 2, Name: "InMem - Go programming", Category: model.BOOKS, Price: 55.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProductById(tt.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductById() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
