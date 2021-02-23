package manage

import (
	"context"
	"fmt"
	"github.com/siller174/goodsGateway/pkg/gateway/repository"
	"github.com/siller174/goodsGateway/pkg/gateway/service/articleservice"
	"github.com/siller174/goodsGateway/pkg/gateway/service/catalogservice"
	"github.com/siller174/goodsGateway/pkg/gateway/service/communication"
	"github.com/siller174/goodsGateway/pkg/gateway/service/itemservice"
	"github.com/siller174/goodsGateway/pkg/gateway/service/priceservice"
	"github.com/siller174/goodsGateway/pkg/gateway/service/shopservice"
	"github.com/siller174/goodsGateway/pkg/gateway/service/telegram"
	"github.com/siller174/goodsGateway/pkg/gateway/structs"
	"github.com/siller174/goodsGateway/pkg/logger"
)

type Manager struct {
	commutator *communication.Commutator

	tgBot *telegram.Bot

	catalogService *catalogservice.Service
	itemService    *itemservice.Service
	priceService   *priceservice.PriceService
	articleService *articleservice.Service
	shopService    *shopservice.Service
}

func (m *Manager) GetCatalogs(ctx context.Context, req *structs.GetCatalogsReq) (*structs.GetCatalogsRsp, error) {

	items, err := m.catalogService.Get(ctx, req.Shop.ID)
	if err != nil {
		return nil, err
	}

	catalogs := make([]*structs.Catalog, len(items))

	for k, v := range items {
		catalogs[k] = &structs.Catalog{
			CategoryID: v.CategoryID,
			URL:        v.Url,
			Count:      v.Count,
		}
	}

	return &structs.GetCatalogsRsp{
		Shop: &structs.Shop{
			ID: req.Shop.ID,
		},
		Catalogs: catalogs,
	}, nil
}

func (m *Manager) SaveGoods(ctx context.Context, req *structs.SaveGoodsReq) (*structs.SaveGoodRsp, error) {

	for _, good := range req.Goods {
		logger.InfoCtx(ctx, "Start process good %+v in shop %v", good, req.Shop.ID)

		isExist, err := m.articleService.IsExist(ctx, good.Article, req.Shop.ID)

		if err != nil {
			logger.ErrorCtx(ctx, "Cannot check exist for good %+v from shop %v. Error %v", good, req.Shop.ID, err)
			return nil, err
		}

		if isExist {
			itemID, err := m.itemService.GetId(ctx, good.Name, good.CategoryID)
			if err != nil {
				logger.ErrorCtx(ctx, "Cannot get itemId for good %+v from shop %v. Error %v", good, req.Shop.ID, err)
				return nil, err
			}

			lastPrice, err := m.priceService.GetLastPrice(ctx, req.Shop.ID, itemID)
			if err != nil {
				logger.ErrorCtx(ctx, "Cannot get last price for good %+v from shop %v. Error %v", good, req.Shop.ID, err)
				return nil, err
			}

			if lastPrice > good.Price {
				logger.InfoCtx(ctx, "SALE!. SendGood good %+v in shop %v to commutator for check sale", good, req.Shop.ID)
				go m.sendGoodToMedia(ctx, req.Shop.ID, good, lastPrice)
			}
			if lastPrice != good.Price {
				logger.InfoCtx(ctx, "New price for good %+v in shop %v", good, req.Shop.ID)
				err = m.priceService.Create(ctx, req.Shop.ID, itemID, good.Price)
				if err != nil {
					logger.ErrorCtx(ctx, "Cannot create price for good %+v from shop %v. Error %v", good, req.Shop.ID, err)
					return nil, err
				}
			}
		} else {
			logger.DebugCtx(ctx, "Create new good %+v in shop %v", good, req.Shop.ID)
			good.ShopId = req.Shop.ID
			err := m.itemService.Create(ctx, good)
			if err != nil {
				logger.ErrorCtx(ctx, "Cannot save good %+v from shop %v. Error %v", good, req.Shop.ID, err)
				return nil, err
			}
		}
		logger.InfoCtx(ctx, "Finish process good %+v in shop %v", good, req.Shop.ID)
	}

	return &structs.SaveGoodRsp{
		Result: true,
	}, nil
}

func (m *Manager) sendGoodToMedia(ctx context.Context, shopiD int, good structs.Good, oldPrice int) {
	proc := 100 - ((good.Price * 100) / oldPrice)

	if proc < 30 {
		logger.InfoCtx(ctx, "Do not send good %+v to media. Discount %v", good, proc)
		return
	}

	shopName, err := m.shopService.GetShopByID(ctx, shopiD)

	if err != nil {
		logger.ErrorCtx(ctx, "Cannot send good %+v to media. Error %v", good, err)
		return
	}

	g := communication.Good{
		ShopName: *shopName,
		Name:     good.Name,
		Price:    good.Price,
		OldPrice: oldPrice,
		ImageUrl: good.ImageURL,
		Url:      good.URL,
		Discount: proc,
	}

	m.commutator.SendGood(ctx, g)
}

func (m *Manager) SaveLogsReq(ctx context.Context, req *structs.SaveLogsReq) (*structs.SaveLogsRsp, error) {
	logger.InfoCtx(ctx, "Handle request to save log. Text %v", req.Text)

	m.tgBot.SendMsgToAdmin(fmt.Sprintf("Message from API\n\n %v", req.Text))

	return &structs.SaveLogsRsp{Result: true}, nil
}

func NewManage(mysql *repository.Mysql, commutator *communication.Commutator, tgBot *telegram.Bot) *Manager {
	catalogService := catalogservice.NewService(catalogservice.NewDAO(mysql))
	itemService := itemservice.NewService(itemservice.NewDAO(mysql))
	priceService := priceservice.NewPriceService(priceservice.NewDAO(mysql))
	articleService := articleservice.NewService(articleservice.NewDAO(mysql))
	shopService := shopservice.NewService(shopservice.NewDAO(mysql))

	return &Manager{
		commutator:     commutator,
		tgBot:          tgBot,
		catalogService: catalogService,
		itemService:    itemService,
		priceService:   priceService,
		articleService: articleService,
		shopService:    shopService,
	}
}
