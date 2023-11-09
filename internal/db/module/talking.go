package module

type Talker struct {
	Email string `json:"email" bson:"email,omitempty"`
	Link  string `json:"link" bson:"link,omitempty"`
}

type Talk struct {
	Id       int64 `json:"id" bson:"id"`
	CreateTs int64 `json:"create_ts" bson:"create_ts,omitempty"`
	ParentId int64 `json:"parent_id" bson:"parent_id,omitempty"`
	Talker
	Text string `json:"text" bson:"text,omitempty"`
}
