package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"gopkg.in/go-playground/validator.v9"
	"skill_blog/com.github.bobacsmall/pkg/response"
	"skill_blog/com.github.bobacsmall/services"
	"skill_blog/com.github.bobacsmall/web/viewmodels"
)

type TagController struct {
	// context is auto-binded by Iris on each request,
	Ctx iris.Context

	// Our ArticleService, it's an interface which
	// is binded from the main application.
	Service services.TagService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session

	response.ApiResponse
}

func NewTagController() *TagController {
	return &TagController{Service: services.NewTagService()}
}

func (c *TagController) PostAdd() {
	var validate = validator.New()
	var tag viewmodels.Tag
	if err := c.Ctx.ReadForm(&tag); err != nil {
		// Handle error.
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString("request params error")
		return

	}
	// Returns InvalidValidationError for bad validation input,
	// nil or ValidationErrors ( []FieldError )
	err := validate.Struct(tag)
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
	err = c.Service.AddTag(tag.Tag)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("标签添加成功", nil))
	}
}

func (c *TagController) PostEdit() {
	var validate = validator.New()
	var tag viewmodels.Tag
	if err := c.Ctx.ReadForm(&tag); err != nil {
		// Handle error.
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString("request params error")
		return

	}
	if !tag.IsValid() {
		c.Ctx.JSON(response.ApiResultFail("ID not null", nil))
		return
	}
	// Returns InvalidValidationError for bad validation input,
	// nil or ValidationErrors ( []FieldError )
	err := validate.Struct(tag)
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
	err = c.Service.EditTag(tag.ID, tag.Tag)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("标签修改成功", nil))
	}
}

func (c *TagController) GetDeleteBy(id uint) {
	if id <= 0 {
		c.Ctx.JSON(response.ApiResultFail("ID not null", nil))
		return
	}
	err := c.Service.DeleteTagById(id)
	if err != nil {
		c.Ctx.JSON(response.ApiResultFail(err.Error(), nil))
	} else {
		c.Ctx.JSON(response.ApiResultSuccess("标签删除成功", nil))
	}
}
