package repositories

import (
	"fmt"
	"skill_blog/com.github.bobacsmall/datasource"
	"skill_blog/com.github.bobacsmall/pkg/setting"
	"testing"
)

func TestOauthTokenRepository_GetOauthTokenByToken(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewOauthTokenRepository()
	oauthToken, e := repository.GetOauthTokenByToken("111111")
	fmt.Println(oauthToken,e)
}
