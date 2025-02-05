package types

type PageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type PageResponse struct {
	Total     int64       `json:"total"`
	List      interface{} `json:"list"`
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
	TotalPage int         `json:"totalPage"`
}