package mp

import (
	"gorm.io/gorm"
	"server/models/mp"
	"server/repository/base"
)

type MpUserRepository struct {
	*base.BaseRepository
}

func NewMpUserRepository(DB *gorm.DB) *MpUserRepository {
	return &MpUserRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建用户
func (r *MpUserRepository) Create(user *mp.MpUser) error {
	return r.DB.Create(user).Error
}

// 更新用户
func (r *MpUserRepository) Update(user *mp.MpUser) error {
	return r.DB.Save(user).Error
}

// 根据ID获取用户
func (r *MpUserRepository) FindByID(id int) (*mp.MpUser, error) {
	var user mp.MpUser
	err := r.DB.First(&user, id).Error
	return &user, err
}

// 根据Email获取用户
func (r *MpUserRepository) FindByEmail(email string) (*mp.MpUser, error) {
	var user mp.MpUser
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// 更新用户状态
func (r *MpUserRepository) UpdateStatus(id int, status int8) error {
	return r.DB.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("user_status", status).Error
}

// 更新用户密码
func (r *MpUserRepository) UpdatePassword(id int, newPassword string) error {
	return r.DB.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("password", newPassword).Error
}

// 更新邮箱验证状态
func (r *MpUserRepository) VerifyEmail(id int) error {
	return r.DB.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("email_verify", 1).Error
}

// 更新用户token
func (r *MpUserRepository) UpdateToken(id int, token string) error {
	return r.DB.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("token", token).Error
}