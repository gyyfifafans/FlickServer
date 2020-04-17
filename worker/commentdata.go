package worker

type CommentDataParam struct {
	Id         int64  `json:"id"`
	LevelId    string `json:"levelID" orm:"size(80)"`
	Comment    string `json:"comment" orm:"size(1024)"`
	UserId     string `json:"userID" orm:size(200)"`
	Updatetime int64  `json:"updateTime"`
}
