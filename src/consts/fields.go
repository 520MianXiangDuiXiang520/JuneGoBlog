package consts

// redis Key
const (
	ArticleIDListCache   = "JuneGo:ArticleIDList" // 文章ID列表
	ArticleInfoHashCache = "JuneGo:ArticleInfo:"
	TagsInfoHashCache    = "JuneGo:TagInfo:"
)

// 友链状态
const (
	FriendShipApproving    = 1 // 审批中
	FriendShipApprovalPass = 2 // 审批通过
	FriendShipApprovalFail = 3 // 审批失败
)
