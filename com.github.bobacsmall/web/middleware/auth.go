package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"net/http"
	"skill_blog/com.github.bobacsmall/services"
	"time"
)

/**
 * 判断 token 是否有效
 * 如果有效 就获取信息并且保存到请求里面
 * @method AuthToken
 * @param  {[type]}  ctx       iris.Context    [description]
 */

// get data
// note 根据用户的token获取用户信息并设置context里面
var oauthToken = services.NewOauthTokenService()

func AuthToken(ctx iris.Context) {
	u := ctx.Values().Get("jwt").(*jwt.Token)          //获取 token 信息
	token, _ := oauthToken.GetOauthTokenByToken(u.Raw) //获取 access_token 信息
	if token.Revoked || token.ExpressIn < time.Now().Unix() {
		ctx.StatusCode(http.StatusUnauthorized)
		//ctx.Next()
		return
	} else {
		ctx.Values().Set("auth_user_id", token.UserId)
	}
	ctx.Next() // execute the "after" handler registered via `DoneGlobal`.
}
