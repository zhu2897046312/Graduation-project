package repository

import (
	"gorm.io/gorm"
	"server/models/mp"
)

type MpUserRepository struct {
	*BaseRepository
}

func NewMpUserRepository(db *gorm.DB) *MpUserRepository {
	return &MpUserRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建用户
func (r *MpUserRepository) Create(user *mp.MpUser) error {
	return r.db.Create(user).Error
}

// 更新用户
func (r *MpUserRepository) Update(user *mp.MpUser) error {
	return r.db.Save(user).Error
}

// 根据ID获取用户
func (r *MpUserRepository) FindByID(id int) (*mp.MpUser, error) {
	var user mp.MpUser
	err := r.db.First(&user, id).Error
	return &user, err
}

// 根据Email获取用户
func (r *MpUserRepository) FindByEmail(email string) (*mp.MpUser, error) {
	var user mp.MpUser
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// 更新用户状态
func (r *MpUserRepository) UpdateStatus(id int, status int8) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("user_status", status).Error
}

// 更新用户密码
func (r *MpUserRepository) UpdatePassword(id int, newPassword string) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("password", newPassword).Error
}

// 更新邮箱验证状态
func (r *MpUserRepository) VerifyEmail(id int) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("email_verify", 1).Error
}

// 更新用户token
func (r *MpUserRepository) UpdateToken(id int, token string) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("token", token).Error
}