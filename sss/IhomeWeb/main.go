package main

import (
        "github.com/micro/go-log"
        "github.com/micro/go-web"
        "github.com/julienschmidt/httprouter"
        "sss/IhomeWeb/handler"
        _ "sss/IhomeWeb/models"
    "net/http"
)

func main() {
	// create new web service
	//创建web服务
    service := web.NewService(
            web.Name("go.micro.web.IhomeWeb"),
            web.Version("latest"),
            web.Address(":22333"),
    )

	// initialise service
	//服务初始化
    if err := service.Init(); err != nil {
            log.Fatal(err)
    }

	// register html handler
	//注册 静态文件
	//service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	//注册路由解析
	//service.HandleFunc("/example/call", handler.ExampleCall)

	rou :=httprouter.New()
	//文件服务器 映射  静态页面

	rou.NotFound =  http.FileServer(http.Dir("html"))


	//模板
	//rou.GET("/example/call", handler.ExampleCall)
	//获取地域信息服务
    rou.GET("/api/v1.0/areas", handler.GetArea)


	//获取图片验证码服务/api/v1.0/imagecode/61bb9476-2a28-4180-9ce6-56a99e16003a
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)
	// 短信验证码服务
	rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmscd)
	// 注册
	rou.POST("/api/v1.0/users", handler.PostRet)


	//欺骗浏览器  session index
	rou.GET("/api/v1.0/session", handler.GetSession)
	//登陆服务
	rou.POST("/api/v1.0/sessions", handler.PostLogin)
	//退出登陆
	rou.DELETE("/api/v1.0/session",handler.DeleteSession)



	//获取用户信息
	rou.GET("/api/v1.0/user", handler.GetUserInfo)







	//index
	rou.GET("/api/v1.0/house/index", handler.GetIndex)



	//将router 注册到 服务
	service.Handle("/",rou)

	// run service
	//服务运行
    if err := service.Run(); err != nil {
            log.Fatal(err)
    }
}
