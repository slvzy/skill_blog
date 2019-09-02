package controllers

import (
	"github.com/kataras/iris"
	"skill_blog/com.github.bobacsmall/pkg/response"
	"skill_blog/com.github.bobacsmall/services"
	"skill_blog/com.github.bobacsmall/web/viewmodels"

	"github.com/kataras/iris/sessions"
	"gopkg.in/go-playground/validator.v9"
)

type ArticleController struct {
	// context is auto-binded by Iris on each request,
	Ctx iris.Context

	// Our ArticleService, it's an interface which
	// is binded from the main application.
	Service services.ArticleService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session

	response.ApiResponse
}

func NewArticleController() *ArticleController {
	return &ArticleController{Service: services.NewArticleService()}
}

func (c *ArticleController) PostAdd() {
	var validate = validator.New()
	var article viewmodels.Article
	if err := c.Ctx.ReadForm(&article); err != nil {
		// Handle error.
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString("request params error")
		return

	}
	// use redis
	//_ = redis.Client.Set("admin", "12344",0).Err()

	// Returns InvalidValidationError for bad validation input,
	// nil or ValidationErrors ( []FieldError )
	err := validate.Struct(article)
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
	// save
	err = c.Service.AddArticle(article.Article)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("success", nil))
	}
}

// 修改
func (c *ArticleController) PostEdit() {
	var validate = validator.New()
	var article viewmodels.Article
	if err := c.Ctx.ReadForm(&article); err != nil {
		// Handle error.
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString("request params error")
		return

	}
	// 判断Id是否为空
	if !article.IsValid() {
		c.Ctx.JSON(response.ApiResultFail("ID not null", nil))
		return
	}
	// params validate
	err := validate.Struct(article)
	if err != nil {
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
	// update
	err = c.Service.EditArticle(article.ID, article.Article)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("success", nil))
	}
}

// 查询
func (c *ArticleController) GetBy(id uint) {
	if id == 0 {
		c.Ctx.JSON(response.ApiResultFail("Id value not null", nil))
		return
	}
	// get by id
	article, err := c.Service.GetArticle(id)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("get success", article))
	}
}

// 删除
func (c *ArticleController) GetDeleteBy(id uint) {
	if id == 0 {
		c.Ctx.JSON(response.ApiResultFail("Id value not null", nil))
		return
	}
	err := c.Service.DeleteArticleById(id)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("delete success", nil))
	}
}
