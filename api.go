package main

import (
	"FlickServer/worker"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

const (
	GET  = "GET"
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

func rounterInit(rounter *gin.Engine, apis []GinHandler) {

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

	//还有问题没看懂
	type HandlerList []GinHandler
	type HandlerMap map[string]HandlerList

	_handlers := make(HandlerMap, len(apis))
	for _, v := range apis {
		if _, exist := _handlers[v.Path]; !exist {
			_handlers[v.Path] = make(HandlerList, 0, 10)
		}
		_handlers[v.Path] = append(_handlers[v.Path], v)
	}

	api := rounter.Group("/v1")
	for path, list := range _handlers {
		api.OPTIONS(path, options)
		for _, v := range list {
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
	rounter.LoadHTMLGlob("views/*")
}

func registerApi(rounter *gin.Engine) {

	// Set a lower memory limit for multipart forms
	rounter.MaxMultipartMemory = 8 << 30 //

	// 注册路由
	apis := []GinHandler{
		{"/account/register", worker.AccountRegister, POST},
		{"/account/login", worker.AccountLogin, POST},
		{"/get_test/:id", testParam, GET}, // 获取路由参数测试
		{"/get_test", testParam, GET},     //  /get_test?id=
		{"/get_music_test", worker.MusicDataGetInitMusic, GET},
		//===========================================================================
		{"/", worker.RouterGetData, GET},
		{"/", worker.RouterPostData, POST},
		{"/test_mix", func(c *gin.Context) {
			fmt.Printf("get.\n")
		}, GET}, // 获取路由参数测试
		{"/test_mix", func(c *gin.Context) {
			fmt.Printf("post.\n")
		}, POST}, // 获取路由参数测试
		//{"/fileuplaod",upload.Fileupload,POST},
		//{"/fileopt",upload.Fileopthtml,GET},

		{"/upload", func(c *gin.Context) {
			// Multipart form
			form, _ := c.MultipartForm()
			files := form.File["upload[]"]

			// variable from the form
			priority := c.PostForm("priority")

			// listing uploading file
			for _, file := range files {
				log.Println(file.Filename)
			}

			for _, file := range files {
				filename := filepath.Base(file.Filename)
				if err := c.SaveUploadedFile(file, "static/video/"+filename); err != nil {
					c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()), priority)
					return
				}
			}
			c.String(http.StatusOK, fmt.Sprintf("%d files uploaded! ", len(files)), priority)
		}, POST},
	}

	rounterInit(rounter, apis)
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
