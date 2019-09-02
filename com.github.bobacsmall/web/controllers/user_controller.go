package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"gopkg.in/go-playground/validator.v9"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/pkg/response"
	"skill_blog/com.github.bobacsmall/pkg/setting"
	"skill_blog/com.github.bobacsmall/services"
	"skill_blog/com.github.bobacsmall/web/viewmodels"
	"time"
)

type UserController struct {
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service      services.UserService
	TokenService services.OauthTokenService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

func NewUserController() *UserController {
	return &UserController{
		Service:      services.NewUserService(),
		TokenService: services.NewOauthTokenService()}
}

func (c *UserController) PostRegister() {
	var validate = validator.New()
	var user viewmodels.Users
	if err := c.Ctx.ReadForm(&user); err != nil {
		// Handle error.
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString("request params error")
		return

	}
	// Returns InvalidValidationError for bad validation input,
	// nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		// This check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.Ctx.StatusCode(iris.StatusInternalServerError)
			c.Ctx.WriteString(err.Error())
			return
		}
		c.Ctx.StatusCode(iris.StatusBadRequest)
		for _, err := range err.(validator.ValidationErrors) {
			c.Ctx.WriteString(err.Field() + " " + err.ActualTag())
			return
		}
	}
	// call userRepository
	err = c.Service.Register(user.User)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("注册成功", nil))
	}
}

func (c *UserController) PostLogin() {
	var validate = validator.New()
	var user viewmodels.Users
	if err := c.Ctx.ReadForm(&user); err != nil {
		// Handle error.
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString("request params error")
		return

	}
	// Returns InvalidValidationError for bad validation input,
	// nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		// This check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.Ctx.StatusCode(iris.StatusInternalServerError)
			c.Ctx.WriteString(err.Error())
			return
		}
		c.Ctx.StatusCode(iris.StatusBadRequest)
		for _, err := range err.(validator.ValidationErrors) {
			c.Ctx.WriteString(err.Field() + " " + err.ActualTag())
			return
		}
	}

	uu, err := c.Service.Login(user.User)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {

		// generate token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(setting.AppSetting.JwtSecret))
		if err != nil {
			c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
			return
		}

		oauth_token := datamodels.OauthToken{}
		oauth_token.Token = tokenString
		oauth_token.UserId = uu.ID
		oauth_token.Secret = setting.AppSetting.JwtSecret
		oauth_token.Revoked = false
		oauth_token.ExpressIn = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		var t = datamodels.Token{}
		t, err = c.TokenService.OauthTokenCreate(oauth_token)
		if err != nil {
			c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
			return
		}
		c.Ctx.JSON(response.ApiResultSuccess("登录成功", t))
	}
}
