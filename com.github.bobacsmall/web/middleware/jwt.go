package middleware

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"skill_blog/com.github.bobacsmall/pkg/response"
	"skill_blog/com.github.bobacsmall/pkg/setting"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func JwtHandler() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(setting.AppSetting.JwtSecret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, err error) { // token验证失败获取失效
			ctx.JSON(response.ApiResultFail(err.Error(),nil))
		},
	})

}
