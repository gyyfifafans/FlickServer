package model

import "FlickServer/common"

type ScoreData struct {
	Id         int64  `json:"id"`
	LevelId    string `json:"levelID" orm:"size(64)"`
	Score      string `json:"score" orm:"size(64)"`
	UserId     string `json:"userID" orm:size(200)"`
	UpdateTime int64  `json:"updateTime"`
}

func (self *ScoreData) QueryWithId(id int64) (*ScoreData, error) {
	db := common.NewOrm()
	r := &ScoreData{}
	if err := db.QueryTable("ScoreData").Filter("id", id).One(r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}
