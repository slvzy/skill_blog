package viewmodels

import "skill_blog/com.github.bobacsmall/datamodels"

type Tag struct {
	datamodels.Tag
}

// deal validate
func (t Tag) IsValid() bool {
	return t.ID > 0
}
