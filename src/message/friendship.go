package message

import "JuneGoBlog/src/dao"

type FriendShipListResp struct {
	Header BaseRespHeader                `json:"header"`         // 响应头
	Total int                            `json:"total"`          // 友链总数
	FriendShipList []dao.FriendShipLink  `json:"friendShipList"` // 友链列表
}
