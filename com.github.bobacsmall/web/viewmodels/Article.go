package viewmodels

import "skill_blog/com.github.bobacsmall/datamodels"

type Article struct {
	datamodels.Article
}

func (a Article) IsValid() bool  {
	return a.ID >0
}