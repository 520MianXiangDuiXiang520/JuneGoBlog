package article

// 请求文章列表格式
type ListReq struct {
	Page int       `form:"page"`           // 页数
	PageSize int   `form:"pageSize"`       // 每页请求的文章数量
	Tag string     `form:"tag"`            // 标签
}
