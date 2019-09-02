package repositories

import (
	"fmt"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/datasource"
	"skill_blog/com.github.bobacsmall/pkg/setting"
	"testing"
)

func TestTagRepository_AddTag(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()
	tag := datamodels.Tag{
		Name:  "Java",
		State: 1,
	}
	err := tagRepository.AddTag(tag)
	fmt.Println(err)
}

func TestTagRepository_ExistTagById(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()
	b, e := tagRepository.ExistTagById(1)
	fmt.Println("====", b, e)
}

func TestTagRepository_ExistTagByName(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()
	b, e := tagRepository.ExistTagByName("java1")
	fmt.Println("====", b, e)
}

func TestTagRepository_GetTags(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()
	tags, e := tagRepository.GetTags(0, 0, datamodels.Tag{
	})
	fmt.Println("====", tags, e)
}

func TestTagRepository_EditTag(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()
	tag := tagRepository.EditTag(2, datamodels.Tag{
		Name: "C#",
	})
	fmt.Println(tag)
}

func TestTagRepository_GetTagTotal(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()

	i, e := tagRepository.GetTagTotal(datamodels.Tag{})
	fmt.Println("------",i,e)
}

func TestTagRepository_DeleteTagById(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()

	id := tagRepository.DeleteTagById(5)
	fmt.Println("---->",id)
}
func TestTagRepository_CleanAllTag(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()
	tag := tagRepository.CleanAllTag()

	fmt.Println(tag)
}

func TestTagRepository_GetTag(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	tagRepository := NewTagRepository()
	tag, e := tagRepository.GetTag(5)
	fmt.Println(tag,e)
}