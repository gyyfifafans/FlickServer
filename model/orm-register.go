package model

import "FlickServer/common"

func RegisterModels()  {
	// 注册数据结构到ORM
	common.RegisterModel(new(Account))
}
