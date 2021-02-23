package middleware

import "github.com/siller174/goodsGateway/pkg/gateway/api/http/errors/handler"

type MiddleWare struct {
	errorHandler *handler.ErrorHandle
}

func NewMiddleWare(errorHandler *handler.ErrorHandle) *MiddleWare {
	return &MiddleWare{
		errorHandler: errorHandler,
	}
}

