package model

import "FlickServer/common"

type CommentData struct {
	Id         int64  `json:"id"`
	LevelId    string `json:"levelID" orm:"size(80)"`
	Comment    string `json:"comment" orm:"size(1024)"`
	UserId     int64  `json:"userID"`
	Updatetime int64  `json:"updateTime"`
}

func (self *CommentData) QueryWithId(id int64) (*CommentData, error) {
	db := common.NewOrm()
	r := &CommentData{}
	if err := db.QueryTable("CommentData").Filter("id", id).One(r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}