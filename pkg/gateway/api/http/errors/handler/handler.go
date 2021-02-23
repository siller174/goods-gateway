package handler

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/api/http/errors"
	"github.com/siller174/goodsGateway/pkg/gateway/api/http/response"
	"github.com/siller174/goodsGateway/pkg/logger"
	"net/http"
)

type ErrorHandle struct {
	devMode bool
}

func NewErrorHandler(devMode bool) *ErrorHandle {
	return &ErrorHandle{
		devMode: devMode,
	}
}

func (handler *ErrorHandle) Handle(ctx context.Context, w http.ResponseWriter, err error) {
	var (
		innerErr error
		resError errors.HTTPError
	)

	if err == nil {
		return
	}

	switch err.(type) {
	case errors.HTTPError:
		resError = err.(errors.HTTPError)
	default:
		resError = errors.NewInternalErr(err)
	}

	logger.ErrorCtx(ctx, resError.Error())

	innerErr = response.WriteJSON(w, resError.GetStatus(), er{resError.ToResponse(handler.devMode)})
	if innerErr != nil {
		logger.ErrorCtx(ctx, err.Error())
	}
}

type er struct {
	Error string
}
