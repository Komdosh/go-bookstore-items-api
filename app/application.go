package app

import (
	"github.com/Komdosh/go-bookstore-items-api/clients/elasticsearch"
	"github.com/Komdosh/go-bookstore-utils/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Addr: ":8081",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	logger.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
