package services

import (
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/repositories"
)

type OauthTokenService interface {
	OauthTokenCreate(oauthToke datamodels.OauthToken) (datamodels.Token, error)
	GetOauthTokenByToken(token string) (*datamodels.OauthToken, error)
	UpdateOauthTokenByUserId(userId uint) (*datamodels.OauthToken, error)
}

func NewOauthTokenService() OauthTokenService {
	return &oauthTokenService{rep: repositories.NewOauthTokenRepository()}
}

type oauthTokenService struct {
	rep repositories.OauthTokenRepository
}

func (a oauthTokenService) OauthTokenCreate(oauthToke datamodels.OauthToken) (datamodels.Token, error) {
	return a.rep.OauthTokenCreate(oauthToke)
}

func (a oauthTokenService) GetOauthTokenByToken(token string) (*datamodels.OauthToken, error) {
	return a.rep.GetOauthTokenByToken(token)
}

func (a oauthTokenService) UpdateOauthTokenByUserId(userId uint) (*datamodels.OauthToken, error) {
	return a.rep.UpdateOauthTokenByUserId(userId)
}
