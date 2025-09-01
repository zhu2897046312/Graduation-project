package mp

import (
	"gorm.io/gorm"
	"server/models/mp"
	"server/repository/base"
	"time"
)

type MpResetPwdTokensRepository struct {
	*base.BaseRepository
}

func NewMpResetPwdTokensRepository(DB *gorm.DB) *MpResetPwdTokensRepository {
	return &MpResetPwdTokensRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建重置密码令牌
func (r *MpResetPwdTokensRepository) Create(token *mp.MpResetPwdTokens) error {
	return r.DB.Create(token).Error
}

// 根据Token获取重置密码记录
func (r *MpResetPwdTokensRepository) FindByToken(token string) (*mp.MpResetPwdTokens, error) {
	var record mp.MpResetPwdTokens
	err := r.DB.Where("token = ?", token).First(&record).Error
	return &record, err
}

// 根据Email获取重置密码记录
func (r *MpResetPwdTokensRepository) FindByEmail(email string) (*mp.MpResetPwdTokens, error) {
	var record mp.MpResetPwdTokens
	err := r.DB.Where("email = ?", email).First(&record).Error
	return &record, err
}

// 更新令牌使用计数
func (r *MpResetPwdTokensRepository) IncrementCount(email string) error {
	return r.DB.Model(&mp.MpResetPwdTokens{}).
		Where("email = ?", email).
		Update("count", gorm.Expr("count + 1")).Error
}

// 删除过期令牌
func (r *MpResetPwdTokensRepository) DeleteExpired() error {
	return r.DB.Where("expiration_time < ?", time.Now()).
		Delete(&mp.MpResetPwdTokens{}).Error
}

// 删除特定令牌
func (r *MpResetPwdTokensRepository) Delete(email string) error {
	return r.DB.Where("email = ?", email).
		Delete(&mp.MpResetPwdTokens{}).Error
}