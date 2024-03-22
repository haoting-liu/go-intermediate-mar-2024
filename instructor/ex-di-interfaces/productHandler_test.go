package main

import (
	"ex-di-interfaces/mocks"
	"ex-di-interfaces/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_get_all_products_handler(t *testing.T) {
	// Arrange
	mockService := mocks.NewProductService(t)
	pr := NewProductRouter(mockService)
	mockService.On("GetAllProducts").Return([]model.Product{{1, "name", model.BEVERAGE, 100.02}})

	router := mux.NewRouter()
	router.HandleFunc("/products", pr.productsHandler)
	request, _ := http.NewRequest(http.MethodGet, "/products", nil)
	recorder := httptest.NewRecorder()
	// Act
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Error("Failed because of invalid content type")
	}
	if recorder.Code != http.StatusOK {
		t.Error("Failed because of invalid status code")
	}
}
