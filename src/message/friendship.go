package message

import "JuneGoBlog/src/dao"

type FriendShipListResp struct {
	Header BaseRespHeader                `json:"header"`                        // 响应头
	Total int                            `json:"total"`                         // 友链总数
	FriendShipList []dao.FriendShipLink  `json:"friendShipList"`                // 友链列表
}

// 添加友链的请求头格式
type FriendAddReq struct {
	SiteName string                      `form:"siteName" json:"siteName"`       // 网站名称（必填）
	SiteLink string                      `form:"siteLink" json:"siteLink"`       // 网站链接（必填）
	ImgLink string                       `form:"imgLink" json:"imgLink"`         // 网站图标链接
	Intro string                         `form:"intro" json:"intro"`             // 网站简介
}

// 添加友链响应格式
type FriendAddResp struct {
	Header BaseRespHeader                `json:"header"`
}

// 删除友链请求格式
type FriendDeleteReq struct {
	ID int                               `form:"id" json:"id"`                  // 要删除的友链ID
}

// 删除友链响应格式
type FriendDeleteResp struct {
	Header BaseRespHeader                `json:"header"`
}