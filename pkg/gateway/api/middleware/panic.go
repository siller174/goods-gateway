package middleware

import (
	"fmt"
	"net/http"
)

func (m *MiddleWare) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				m.errorHandler.Handle(r.Context(), w, fmt.Errorf("recover error. Error %v", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
