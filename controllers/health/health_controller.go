package health

import (
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
}
