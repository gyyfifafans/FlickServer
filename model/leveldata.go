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
	//MusicData   *MusicData `json:"music_data" orm:"rel(fk)"`
}

//func (self *LevelData) QueryWithMovieUrl(url string)([]*LevelData, error) {
//
//}

func (self *LevelData) QueryAllLevels() ([]*LevelData, error) {
	db := common.NewOrm()
	r := make([]*LevelData, 0, 60)
	m := make([]*MusicData, 0, 14)
	//初期先只提供这么多
	if _, err := db.QueryTable("MusicData").Limit(14).All(&m); err != nil {
		return nil, err
	}

	for v := range m {
		//var tr =  []*LevelData{}
		var t = make([]*LevelData, 0, 10)
		if _, err := db.QueryTable("LevelData").Filter("movie_url", m[v].MovieUrl).Filter("creator", "MIZUSHIKI").Filter("creator", "mizushiki").All(&t); err != nil {
			return nil, err
		} else {
			println(len(t))
			//println(tr[0].MovieUrl)
			r = append(t)

		}
	}
	return r, nil

}
