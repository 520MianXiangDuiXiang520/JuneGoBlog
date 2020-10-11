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

const (
	MaxUsernameLength = 25 // username 的最大长度
	MinUsernameLength = 4  // username 的最小长度
)

const (
	AbstractSplitStr = "<!-- more -->"
)

const (
	VisitorPermission = 1 // 游客
	AdminPermission   = 2 // 管理员
)

const CacheTagsSplitStr = "-"

// 评论类型
const (
	RootTalkType  = 1
	ChildTalkType = 2
)

const MaxArticleTitleLen = 100
