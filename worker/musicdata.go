package worker

import (
	"FlickServer/model"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

//这个是具体做什么用的，需要和表的结构一致么？好像无法去到传进来的参数
type MusicDataParam struct {
	Id         int64  `json:"id"`
	Title      string `json:"title" orm:"size(64)"`
	Artist     string `json:"artist" orm:"size(64)"`
	UpdateTime int64  `json:"updateTime"`
}

func MusicDataGetInitMusic(c *gin.Context) {
	var param = MusicDataParam{}
	if err := c.Bind(&param); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	}
	musicData := &model.MusicData{}
	if r, err := musicData.QueryAllMusics(); err != nil {
		if err.Error() == orm.ErrNoRows.Error() {
			respJSON(c, Result{
				Status: 500,
				Msg:    "查不到音乐",
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

func MusicDataGetLastUpdateMusic(c *gin.Context, t int64) {
	var param = MusicDataParam{}
	if err := c.Bind(&param); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	}
	musicData := &model.MusicData{}
	if r, err := musicData.QueryWithUpdateTime(t); err != nil {
		if err.Error() == orm.ErrNoRows.Error() {
			respJSON(c, Result{
				Status: 500,
				Msg:    "查不到音乐",
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
