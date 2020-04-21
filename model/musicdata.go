package model

import "FlickServer/common"

type MusicData struct {
	Id int64 `json:"id"`
	//因为movie_url和leveldata中的不一致，本来是作为过滤没编辑完level的不展示音乐
	MovieUrl     string `json:"movieURL" orm:"size(200)"`
	ThumbnailUrl string `json:"thumbnailURL" orm:"size(200)"`
	Title        string `json:"title" orm:"size(64)"`
	Artist       string `json:"artist" orm:"size(64)"`
	MovieLength  string `json:"movieLength" orm:"size(64)"`
	Tags         string `json:"tags" orm:"type(text)"`
	UpdateTime   int64  `json:"updateTime"`
	CreateTime   int64  `json:"createTime"`
	//Levels        []*LevelData `orm:"reverse(many)"`
}

func (self *MusicData) QueryAllMusics() ([]*MusicData, error) {
	db := common.NewOrm()
	r := make([]*MusicData, 0, 15)
	//初期先只提供这么多
	if _, err := db.QueryTable("MusicData").Limit(14).All(&r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (self *MusicData) QueryWithUpdateTime(time int64) ([]*MusicData, error) {
	db := common.NewOrm()
	r := make([]*MusicData, 0, 14)
	if _, err := db.QueryTable("MusicData").Filter("update_time__gt", time).Limit(14).All(&r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}
