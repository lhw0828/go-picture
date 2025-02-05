package types

type PageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type DeleteRequest struct {
	Id int64 `path:"id"`
}

type GetRequest struct {
	Id int64 `path:"id"`
}

type PageResponse struct {
	Total     int64       `json:"total"`
	List      interface{} `json:"list"`
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
	TotalPage int         `json:"totalPage"`
}
