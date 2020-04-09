package worker

import (
	"FlickServer/model"
	"github.com/gin-gonic/gin"
)

func AccountRegister(c *gin.Context) {
	var param model.Account
	if err := c.Bind(&param); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	}
	if len(param.Username) == 0 || len(param.Password) == 0 || len(param.Nickname) == 0 {
		respJSON(c, Result{
			Status: 500,
			Msg:    "参数非法",
		})
		return
	}
	account := &model.Account{}
	if r, err := account.Add(param.Username, param.Password, param.Nickname); err != nil {
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
}

func AccountLogin(c *gin.Context) {
	var param model.Account
	if err := c.Bind(&param); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	}
	if len(param.Username) == 0 || len(param.Password) == 0 {
		respJSON(c, Result{
			Status: 500,
			Msg:    "参数非法",
		})
		return
	}
	account := &model.Account{}
	if r, err := account.QueryWithUsername(param.Username); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	} else {
		if r.Password != param.Password {
			respJSON(c, Result{
				Status: 500,
				Msg:    "密码错误",
			})
			return
		}
		respJSON(c, Result{
			Status: 200,
			Data:   r,
		})
	}
}
