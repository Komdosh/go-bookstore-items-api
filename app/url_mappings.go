package app

import (
	"github.com/Komdosh/go-bookstore-items-api/controllers/health"
	"github.com/Komdosh/go-bookstore-items-api/controllers/items"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", health.HealthController.Health).Methods(http.MethodGet)

	router.HandleFunc("/items", items.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", items.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", items.ItemsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", items.ItemsController.Update).Methods(http.MethodPut)
	router.HandleFunc("/items/{id}", items.ItemsController.Update).Methods(http.MethodPatch)
	router.HandleFunc("/items/{id}", items.ItemsController.Delete).Methods(http.MethodDelete)
}
