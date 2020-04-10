package main

import (
	"FlickServer/worker"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	GET = "GET"
	POST = "POST"
)

type GinHandler struct {
	Path    string
	Handler gin.HandlerFunc
	Method  string
}

func options(c *gin.Context) {
	c.String(202, "fuck you")
}

func testParam(c *gin.Context) {
	// GET /path/:id
	// c.Param("id") == ":id"
	fmt.Printf("路由参数：%+v\n", c.Param("参数名"))
	// GET /path?id=1234&name=Manu&value=
	// c.Query("id") == "1234"
	// c.Query("name") == "Manu"
	// c.Query("value") == ""
	// c.Query("wtf") == ""
	fmt.Printf("url参数：%+v\n", c.Query("参数名"))
	c.String(200, "it works!")
}

func registerApi(rounter *gin.Engine) {

	// 注册路由
	apis := []GinHandler{
		{"/account/register", worker.AccountRegister, POST},
		{"/account/login", worker.AccountLogin, POST},
		{"/get_test/:id", testParam, GET}, // 获取路由参数测试
	}

	// 路由全局设置
	store := cookie.NewStore([]byte("secret-flickServer-&^#$%&*"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true, // 对js隐藏session
	})
	rounter.Use(sessions.Sessions("flick-session", store))
	rounter.Use(gin.Logger())
	rounter.Use(gin.Recovery())
	rounter.Use(makeAccessJsMiddleware()) // 跨域处理放前面
	// 接口不分组
	{/*
		for _, v := range apis {
			rounter.OPTIONS(v.Path, options) // 浏览器http请求会先触发一次options请求
			switch v.Method {
			case POST:
				rounter.POST(v.Path, v.Handler)
			case GET:
				rounter.GET(v.Path, v.Handler)
			}
		}*/
	}
	// 接口分组
	api := rounter.Group("/v1")
	{
		for _, v := range apis {
			api.OPTIONS(v.Path, options) // 浏览器http请求会先触发一次options请求
			switch v.Method {
			case POST:
				api.POST(v.Path, v.Handler)
			case GET:
				api.GET(v.Path, v.Handler)
			}
		}
	}
	// 开放静态资源目录
	rounter.Static("/static", "./static")
}

func makeAccessJsMiddleware() gin.HandlerFunc {
	// 处理js-ajax跨域问题
	return func(c *gin.Context) {
		w := c.Writer
		//w.Header().Set("Access-Control-Allow-Origin", "*") // 允许访问所有域
		w.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin")) // 允许cookie则需要origin对应
		w.Header().Set("Access-Control-Allow-Credentials", "true")                    // 允许cookie
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Origin")
	}
}
