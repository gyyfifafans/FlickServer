package worker

type LevelDataParam struct {
	Id          int64  `json:"id"`
	MovieUrl    string `json:"movieURL" orm:"size(200)"`
	Level       string `json:"level" orm:"size(64)"`
	Creator     string `json:"creator" orm:"size(64)"`
	Description string `json:"description" orm:"size(1024)"`
	Speed       string `json:"speed" orm:"size(64)"`
	Notes       string `json:"notes" orm:"type(text)"`
	UpdateTime  int64  `json:"updateTime"`
	CreateTime  int64  `json:"createTime"`
	PlayCount   int32  `json:"playCount"`
}
