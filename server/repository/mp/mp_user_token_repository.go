package mp

import (
	"gorm.io/gorm"
	"server/models/mp"
	"server/repository/base"
)

type MpUserTokenRepository struct {
	*base.BaseRepository
}

func NewMpUserTokenRepository(DB *gorm.DB) *MpUserTokenRepository {
	return &MpUserTokenRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建用户令牌
func (r *MpUserTokenRepository) Create(token *mp.MpUserToken) error {
	return r.DB.Create(token).Error
}

// 根据Token获取令牌
func (r *MpUserTokenRepository) FindByToken(token string) (*mp.MpUserToken, error) {
	var userToken mp.MpUserToken
	err := r.DB.Where("token = ?", token).First(&userToken).Error
	return &userToken, err
}

// 根据用户ID获取令牌列表
func (r *MpUserTokenRepository) FindByUserID(userID uint) ([]mp.MpUserToken, error) {
	var tokens []mp.MpUserToken
	err := r.DB.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&tokens).Error
	return tokens, err
}

// 删除令牌
func (r *MpUserTokenRepository) Delete(id uint) error {
	return r.DB.Delete(&mp.MpUserToken{}, id).Error
}

// 删除用户的所有令牌
func (r *MpUserTokenRepository) DeleteByUserID(userID uint) error {
	return r.DB.Where("user_id = ?", userID).
		Delete(&mp.MpUserToken{}).Error
}

// 删除过期令牌
func (r *MpUserTokenRepository) DeleteExpired() error {
	return r.DB.Where("expire_time < NOW()").
		Delete(&mp.MpUserToken{}).Error
}