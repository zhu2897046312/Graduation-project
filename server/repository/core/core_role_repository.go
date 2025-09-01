package core

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/repository/base"
)

type CoreRoleRepository struct {
	*base.BaseRepository
}

func NewCoreRoleRepository(DB *gorm.DB) *CoreRoleRepository {
	return &CoreRoleRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建角色
func (r *CoreRoleRepository) Create(role *core.CoreRole) error {
	return r.DB.Create(role).Error
}

// 更新角色
func (r *CoreRoleRepository) Update(role *core.CoreRole) error {
	return r.DB.Save(role).Error
}

// 根据ID获取角色
func (r *CoreRoleRepository) FindByID(id int64) (*core.CoreRole, error) {
	var role core.CoreRole
	err := r.DB.First(&role, id).Error
	return &role, err
}

// 获取所有角色
func (r *CoreRoleRepository) FindAll() ([]core.CoreRole, error) {
	var roles []core.CoreRole
	err := r.DB.Find(&roles).Error
	return roles, err
}

// 更新角色状态
func (r *CoreRoleRepository) UpdateStatus(id int64, status int8) error {
	return r.DB.Model(&core.CoreRole{}).
		Where("id = ?", id).
		Update("role_status", status).Error
}

// 更新角色权限
func (r *CoreRoleRepository) UpdatePermissions(id int64, permissions []byte) error {
	return r.DB.Model(&core.CoreRole{}).
		Where("id = ?", id).
		Update("permission", permissions).Error
}