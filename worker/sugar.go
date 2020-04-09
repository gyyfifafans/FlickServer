package worker

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Status int64       `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

type List struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func respJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
