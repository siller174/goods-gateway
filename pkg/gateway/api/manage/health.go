package manage

import (
	"github.com/siller174/goodsGateway/pkg/gateway/api/http/response"
	"net/http"
)

const HealthRoute = "/health"

type Health struct {
}

func NewHealthApi() *Health {
	return &Health{

	}
}

func (health *Health) Handle() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		_ = response.WriteJSON(w, http.StatusOK, []byte(`{"status": "UP"}`))
	}
}
