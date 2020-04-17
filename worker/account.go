package worker

import (
	"FlickServer/model"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

type AccountParam struct {
	Username string `json:"username" orm:"size(80)"`
	Password string `json:"password" orm:"size(64)"`
}

func AccountRegister(c *gin.Context) {
	var param AccountParam
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
	if has, err := account.HasUsername(param.Username); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	} else if has {
		respJSON(c, Result{
			Status: 500,
			Msg:    "账号已经注册过",
		})
		return
	}
	if r, err := account.Add(param.Username, param.Password); err != nil {
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
	var param AccountParam
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
		if err.Error() == orm.ErrNoRows.Error() {
			respJSON(c, Result{
				Status: 500,
				Msg:    "查不到账号",
			})
			return
		}
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

func AccountGetUserName(c *gin.Context) {
	var param = AccountParam{}
	if err := c.Bind(&param); err != nil {
		respJSON(c, Result{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	}
	print("go there!!!")
}
