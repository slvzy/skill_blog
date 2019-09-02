package main

import (
	"github.com/betacraft/yaag/yaag"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"skill_blog/com.github.bobacsmall/datasource"
	"skill_blog/com.github.bobacsmall/pkg/redis"
	resp "skill_blog/com.github.bobacsmall/pkg/response"
	"skill_blog/com.github.bobacsmall/pkg/setting"
	"skill_blog/com.github.bobacsmall/route"
	"skill_blog/com.github.bobacsmall/web/middleware"
)

func init() {
	// init config
	setting.Setup()
	// init database
	datasource.Setup()
	// init redis
	redis.Setup()
}

func main() {
	// init instance
	app := newApp()
	// Recover middleware recovers from any panics and writes a 500 if there was one.
	app.Use(recover.New())
	// log middleware
	app.Use(middleware.RequestLog())
	// add router
	route.InitRouter(app)

	// print router
	for _, v := range app.GetRoutes() {
		if v.Method == "OPTIONS" {
			continue
		}
		app.Logger().Infof("method：" + v.Method + "  path：" + v.Path)
	}
	app.Run(iris.Addr(":9090"))
}

func newApp() *iris.Application {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	app.Use(crs) //
	/*response.StaticWeb("/public", "./web/public/")
	response.RegisterView(iris.HTML("./web/views/", ".html"))*/
	app.AllowMethods(iris.MethodOptions)
	//response.Use(middleware.GetJWT().Serve) //是否启用jwt中间件
	app.Configure(iris.WithOptimizations)

	// start open generate document
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: setting.AppSetting.Name,
		DocPath:  "index.html", //设置绝对路径
		BaseUrls: map[string]string{
			"Production": setting.AppSetting.Url,
			"Staging":    "test",
		},
	})

	// shutdown application close database register interrupt
	iris.RegisterOnInterrupt(func() {
		datasource.CloseDB()
	})

	// 404 error deal
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(resp.ApiResult(0, "404 Not Found", nil))
	})

	// 500 error deal
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Oups something went wrong, try again")
	})
	return app
}
