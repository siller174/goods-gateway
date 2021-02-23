package middleware

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/api/http/response"
	"github.com/siller174/goodsGateway/pkg/logger"
	"net/http"
)

func (m *MiddleWare) Handle(f func(r *http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(r.Context())
		defer cancelFunc()
		r.WithContext(ctx)

		result, err := f(r)
		if err != nil {
			m.errorHandler.Handle(ctx, w, err)
			return
		}

		m.response(ctx, w, result)
	}
}

func (m *MiddleWare) response(ctx context.Context, w http.ResponseWriter, val interface{}) {
	err := response.WriteJSON(w, http.StatusOK, val)
	if err != nil {
		logger.ErrorCtx(ctx, "Could not write middleware %v to response", val)
		m.errorHandler.Handle(ctx, w, err)
		return
	}
}
