package worker

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Status int64       `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg,omitempty"`
}

type List struct {
	Total int64       `json:"total,omitempty"`
	List  interface{} `json:"list,omitempty"`
}

func respJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
