package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}