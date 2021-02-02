package health

import (
	"github.com/Komdosh/go-bookstore-utils/logger"
	"net/http"
)

const (
	health = "alive"
)

var (
	HealthController healthControllerInterface = &healthController{}
)

type healthControllerInterface interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type healthController struct{}

func (c *healthController) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(health))
	logger.Info("Health...")
}
