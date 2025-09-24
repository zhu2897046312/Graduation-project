package repository

import (
	"gorm.io/gorm"
	"server/models/mp"
	"server/models/common"
	"time"
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
func (r *MpUserRepository) FindByID(id common.MyID) (*mp.MpUser, error) {
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
func (r *MpUserRepository) UpdateStatus(id common.MyID, status int8) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("user_status", status).Error
}

// 更新用户密码
func (r *MpUserRepository) UpdatePassword(id common.MyID, newPassword string) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("password", newPassword).Error
}

// 更新邮箱验证状态
func (r *MpUserRepository) VerifyEmail(id common.MyID) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("email_verify", 1).Error
}

// 更新用户token
func (r *MpUserRepository) UpdateToken(id common.MyID, token string) error {
	return r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("token", token).Error
}

func (r *MpUserRepository) ListWithPagination(params mp.MpUserListParam) ([]mp.MpUser, int64, error) {
	var products []mp.MpUser
	var total int64
	// 构建查询
	query := r.db.Model(&mp.MpUser{}).Where("deleted_time IS NULL")

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if params.PageSize <= 0 {
		err := query.Find(&products).Error
		return products, total, err
	} else {

		// 设置默认值
		if params.Page < 1 {
			params.Page = 1
		}
		if params.PageSize < 1 || params.PageSize > 100 {
			params.PageSize = 10
		}
		offset := (params.Page - 1) * params.PageSize
		// 获取分页数据
		err := query.Offset(offset).
			Limit(params.PageSize).
			Find(&products).Error

		return products, total, err
	}

}

// 软删除商品（GORM 自动处理）
func (r *MpUserRepository) DeleteByID(id common.MyID) error {
	result := r.db.Model(&mp.MpUser{}).
		Where("id = ?", id).
		Update("deleted_time", time.Now())

	return result.Error
}