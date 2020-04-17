package model

import "FlickServer/common"

type LevelData struct {
	Id          int64  `json:"id"`
	MovieUrl    string `json:"movieURL" orm:"size(200)"`
	Level       string `json:"level" orm:"size(64)"`
	Creator     string `json:"creator" orm:"size(64)"`
	Description string `json:"description" orm:"size(1024)"`
	Speed       string `json:"speed" orm:"size(64)"`
	Notes       string `json:"notes" orm:"type(text)"`
	UpdateTime  int64  `json:"updateTime"`
	CreateTime  int64  `json:"createTime"`
	PlayCount   int64  `json:"playCount"`
}

func (self *LevelData) QueryWithId(id int64) (*LevelData, error) {
	db := common.NewOrm()
	r := &LevelData{}
	if err := db.QueryTable("LevelData").Filter("id", id).One(r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (self *LevelData) QueryWithMovieUrl() {

}
