package structs

type SaveGoodsReq struct {
	Shop  Shop   `json:"shop"`
	Goods []Good `json:"goods"`
}

type SaveGoodRsp struct {
	Result bool `json:"result"`
}

type Good struct {
	ShopId     int    `json:"-"`
	CategoryID int    `json:"categoryID"`
	Article    string `json:"article"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	ImageURL   string `json:"imageUrl"`
	Price      int    `json:"price"`
}
