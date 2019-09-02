package repositories

import (
	"github.com/jinzhu/gorm"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/datasource"
)

type OauthTokenRepository interface {
	OauthTokenCreate(oauthToke datamodels.OauthToken) (datamodels.Token, error)
	GetOauthTokenByToken(token string) (*datamodels.OauthToken, error)
	UpdateOauthTokenByUserId(userId uint) (*datamodels.OauthToken, error)
}

func NewOauthTokenRepository() OauthTokenRepository {
	return &oauthTokenRepository{db: datasource.DB}
}

type oauthTokenRepository struct {
	db *gorm.DB
}

func (at oauthTokenRepository) OauthTokenCreate(oauthToke datamodels.OauthToken) (datamodels.Token, error) {
	err := at.db.Create(&oauthToke).Error
	return datamodels.Token{Token: oauthToke.Token}, err
}

func (at oauthTokenRepository) GetOauthTokenByToken(token string) (*datamodels.OauthToken, error) {
	var ot datamodels.OauthToken
	// note  deal auth get token db is nil
	if at.db == nil {
		at.db = datasource.DB
	}
	err := at.db.Where("token=?", token).First(&ot).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &ot, nil
}

func (at oauthTokenRepository) UpdateOauthTokenByUserId(userId uint) (*datamodels.OauthToken, error) {
	var ot datamodels.OauthToken
	err := at.db.Model(ot).Where("revoked = ?", false).Where("user_id = ?", userId).Updates(map[string]interface{}{"revoked": true}).Error
	return &ot, err
}
