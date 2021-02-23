package middleware

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/logger"
	"github.com/siller174/goodsGateway/pkg/utils/rquid"
	"net/http"
)

// ContextRequest parse x-request-id from Header and set in context
func (m *MiddleWare) ContextRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := setUid(r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func setUid(r *http.Request) context.Context {
	rqId := r.Header.Get("x-request-id")

	if rqId == "" {
		rqId = rquid.CreateReqUid()
	}

	return context.WithValue(r.Context(), logger.RequestLogField, rqId)
}
