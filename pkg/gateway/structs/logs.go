package structs

type SaveLogsReq struct {
	Text string `json:"text"`
}

type SaveLogsRsp struct {
	Result bool `json:"result"`
}
