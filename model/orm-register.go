package model

import "FlickServer/common"

func RegisterModels() {
	// 注册数据结构到ORM
	common.RegisterModel(new(Account))
	common.RegisterModel(new(CommentData))
	common.RegisterModel(new(LevelData))
	common.RegisterModel(new(ScoreData))
	common.RegisterModel(new(MusicData))

}
