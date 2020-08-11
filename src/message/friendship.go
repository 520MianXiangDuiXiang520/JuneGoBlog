package message

import "JuneGoBlog/src/dao"

type FriendShipListResp struct {
	Header BaseRespHeader                `json:"header"`         // 响应头
	Total int                            `json:"total"`          // 友链总数
	FriendShipList []dao.FriendShipLink  `json:"friendShipList"` // 友链列表
}

// 添加友链的请求头格式
type FriendAddReq struct {
	SiteName string                      `form:"siteName"`       // 网站名称（必填）
	SiteLink string                      `form:"siteLink"`       // 网站链接（必填）
	ImgLink string                       `form:"imgLink"`        // 网站图标链接
	Intro string                         `form:"intro"`          // 网站简介
}

// 添加友链响应格式
type FriendAddResp struct {
	Header BaseRespHeader                `json:"header"`
}