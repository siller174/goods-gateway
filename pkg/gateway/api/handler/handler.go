package handler

import (
	"encoding/json"
	cerror "errors"
	"github.com/siller174/goodsGateway/pkg/gateway/api/http/errors"
	"github.com/siller174/goodsGateway/pkg/gateway/service/manage"
	"net/http"
	"strconv"

	"github.com/siller174/goodsGateway/pkg/gateway/structs"
	"github.com/siller174/goodsGateway/pkg/logger"
)

const api = "/api/v1"
const RouteGetCatalogs = api + "/catalogs"
const RouteSaveGoods = api + "/goods"
const RouteSaveLogs = api + "/logs"

type Handler struct {
	service *manage.Manager
}

func NewHandler(service *manage.Manager) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetCatalogs(r *http.Request) (interface{}, error) {
	var (
		ctx                = r.Context()
		vars               = r.URL.Query()
		shopID, haveShopId = vars["shopId"]
	)

	logger.InfoCtx(ctx, "GetCatalogs request")

	if !haveShopId {
		return nil, errors.NewBadRequest(cerror.New("ShopId not found in query"))
	}

	id, err := strconv.ParseInt(shopID[0], 10, 64)
	if err != nil {
		return nil, errors.NewBadRequest(err)
	}

	req := &structs.GetCatalogsReq{
		Shop: &structs.Shop{
			ID: int(id),
		},
	}

	rsp, err := h.service.GetCatalogs(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (h *Handler) SaveGoods(r *http.Request) (interface{}, error) {
	var (
		ctx = r.Context()
		req = new(structs.SaveGoodsReq)
	)

	logger.InfoCtx(ctx, "SaveGoods request")

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := h.service.SaveGoods(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (h *Handler) SaveLogs(r *http.Request) (interface{}, error) {
	var (
		ctx = r.Context()
		req = new(structs.SaveLogsReq)
	)

	logger.InfoCtx(ctx, "SaveLogs request")

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := h.service.SaveLogsReq(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
