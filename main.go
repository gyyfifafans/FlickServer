package main

import (
	"FlickServer/common"
	"FlickServer/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func loadConfig() {
	model.Config.Load("config/app.ini")
}

func initDatabase() {
	common.InitOrm("default", false)
	common.ConnectMySQL(
		model.Config.Get("mysql.user"),
		model.Config.Get("mysql.pass"),
		model.Config.Get("mysql.database"),
		model.Config.Get("mysql.host")+":"+model.Config.Get("mysql.port"),
	)
	model.RegisterModels()
	parseArgs() // 处理命令行
}

func initLogger() {

}

func startServer() {
	if model.Config.Int64("server.DEBUG") == 1 {
		gin.SetMode(gin.DebugMode)	// 全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	flickServer := gin.New()
	registerApi(flickServer)

	host := model.Config.Get("server.host")
	port := model.Config.Get("server.port")
	flickServer.Run(fmt.Sprintf("%s:%s", host, port))
}

func main() {
	// 初始化
	loadConfig()
	initDatabase()
	initLogger()
	// 启动服务器
	startServer()
}

func parseArgs() {
	switch len(os.Args) {
	case 2:
		// 重建数据库
		if os.Args[1] == "-syncdb" {
			common.Syncdb()
			os.Exit(0)
		}
	}
}