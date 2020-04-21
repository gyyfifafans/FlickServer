package worker

import (
	"FlickServer/model"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

type LevelDataParam struct {
	MovieUrl string `json:"movieURL" orm:"size(200)"`
	Level    string `json:"level" orm:"size(64)"`
	Notes    string `json:"notes" orm:"type(text)"`
}

func LevelDataGetInitLevel(c *gin.Context) {
	var param = LevelDataParam{}
	if err := c.Bind(&param); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	}
	levelData := &model.LevelData{}
	if r, err := levelData.QueryAllLevels(); err != nil {
		if err.Error() == orm.ErrNoRows.Error() {
			respJSON(c, Result{
				Status: 500,
				Msg:    "查不到等级",
			})
			return
		}
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	} else {
		respJSON(c, Result{
			Status: 200,
			Data:   r,
		})
	}
	print("go there!!!")
}
