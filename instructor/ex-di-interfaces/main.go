package main

import (
	"encoding/json"
	"ex-di-interfaces/repository"
	"ex-di-interfaces/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ProductRouter struct {
	service service.DefaultProductService
}

func NewProductRouter(ps service.DefaultProductService) ProductRouter {
	return ProductRouter{ps}
}

func main() {
	// creating instances of dependencies
	dbRepo := repository.NewDbProductRepo()
	// wiring your application (configuration code)
	ps := service.NewProductService(dbRepo)
	router := NewProductRouter(ps)

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(profilingMiddleware)

	r.HandleFunc("/products", router.productsHandler).Methods(http.MethodGet)
	r.HandleFunc("/products", router.newProductsHandler).Methods(http.MethodPost)
	r.HandleFunc("/products/{id:[0-9]+}", router.singleProductHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func profilingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request: ", r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Outgoing response: ")
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("Endpoint: %s | Elasped: %v \n", r.URL.Path, time.Since(start))
	})
}

func (pr ProductRouter) newProductsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: In the lab complete this functionality
}

func (pr ProductRouter) productsHandler(w http.ResponseWriter, r *http.Request) {
	products := pr.service.GetAllProducts()
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Some error occoured"))
	}
}

func (pr ProductRouter) singleProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productId, _ := strconv.Atoi(id)

	product := pr.service.GetProductById(productId)
	w.Header().Add("Content-Type", "application/json")

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		errRes := ErrorResponse{"product does not exist"}
		json.NewEncoder(w).Encode(errRes)
	} else {
		if err := json.NewEncoder(w).Encode(product); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Some error occoured"))
		}
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
