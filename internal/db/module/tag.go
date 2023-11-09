package module

type Tag struct {
	Id       int64  `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name,omitempty"`
	CreateTs int64  `json:"create_ts" bson:"create_ts,omitempty"`
}
