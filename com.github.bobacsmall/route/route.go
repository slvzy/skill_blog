package route

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"skill_blog/com.github.bobacsmall/web/controllers"
	"skill_blog/com.github.bobacsmall/web/middleware"
)

func InitRouter(app *iris.Application) {

	// 配置路由
	v1 := app.Party("/article/")
	v1.Use(middleware.NewYaag())
	v1.Use(middleware.JwtHandler().Serve, middleware.AuthToken, before)
	mvc.New(v1).Handle(controllers.NewArticleController())

	// user party
	uv1 := app.Party("/user/")
	mvc.New(uv1).Handle(controllers.NewUserController())

	// tag
	tv1 := app.Party("/tag/")
	tv1.Use(middleware.NewYaag())
	tv1.Use(middleware.JwtHandler().Serve, middleware.AuthToken)
	mvc.New(tv1).Handle(controllers.NewTagController())
	// other router

}

func before(ctx iris.Context) {
	fmt.Println("======before===========")
	ctx.Next()
}

func after(ctx iris.Context) {
	fmt.Println("======after===========")
	ctx.Next()
}

func middleHander(ctx iris.Context) {
	fmt.Println("=======middleHander==========")
	ctx.Next()
}
