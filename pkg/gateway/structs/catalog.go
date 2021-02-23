package structs

type GetCatalogsReq struct {
	Shop *Shop `json:"shop"`
}

type GetCatalogsRsp struct {
	Shop     *Shop      `json:"shop"`
	Catalogs []*Catalog `json:"catalogs"`
}

type Catalog struct {
	CategoryID int    `json:"categoryID"`
	URL        string `json:"url"`
	Count      int    `json:"count"`
}
