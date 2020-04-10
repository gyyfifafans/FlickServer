package model

import "FlickServer/common"

type MusicData struct {
	Id           int64  `json:"id"`
	MovieUrl     string `json:"movieURL" orm:"size(200)"`
	ThumbnailUrl string `json:"thumbnailURL" orm:"size(200)"`
	Title        string `json:"title" orm:"size(64)"`
	Artist       string `json:"artist" orm:"size(64)"`
	MovieLength  string `json:"movieLength" orm:"size(64)"`
	Tags         string `json:"tags" orm:"size(64)"`
	UpdateTime   int64  `json:"updateTime"`
	CreateTime   int64  `json:"createTime"`
}

func (self *MusicData) QueryWithId(id int64) (*MusicData, error) {
	db := common.NewOrm()
	r := &MusicData{}
	if err := db.QueryTable("MusicData").Filter("id", id).One(r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}
