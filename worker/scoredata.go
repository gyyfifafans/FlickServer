package worker

type ScoreDataParam struct {
	Id         int64  `json:"id"`
	LevelId    string `json:"levelID" orm:"size(64)"`
	Score      string `json:"score" orm:"size(64)"`
	UserId     string `json:"userID" orm:size(200)"`
	UpdateTime int64  `json:"updateTime"`
}
