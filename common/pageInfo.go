package common

type PageInfo struct {
	PageNumber int `form:"pageNumber"`
	PageSize   int `form:"pageSize"`
	Limit      int
	Offset     int
}
