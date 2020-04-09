package model

import "FlickServer/common"

var Config *common.Config

func init() {
	Config = common.NewConfig()
}
