package repository

import (
	"gorm.io/gorm"
	"server/models/core"
)

type CoreRoleRepository struct {
	*BaseRepository
}

func NewCoreRoleRepository(db *gorm.DB) *CoreRoleRepository {
	return &CoreRoleRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建角色
func (r *CoreRoleRepository) Create(role *core.CoreRole) error {
	return r.db.Create(role).Error
}

// 更新角色
func (r *CoreRoleRepository) Update(role *core.CoreRole) error {
	return r.db.Save(role).Error
}

// 根据ID获取角色
func (r *CoreRoleRepository) FindByID(id int64) (*core.CoreRole, error) {
	var role core.CoreRole
	err := r.db.First(&role, id).Error
	return &role, err
}

// 获取所有角色
func (r *CoreRoleRepository) FindAll() ([]core.CoreRole, error) {
	var roles []core.CoreRole
	err := r.db.Find(&roles).Error
	return roles, err
}

// 更新角色状态
func (r *CoreRoleRepository) UpdateStatus(id int64, status int8) error {
	return r.db.Model(&core.CoreRole{}).
		Where("id = ?", id).
		Update("role_status", status).Error
}

// 更新角色权限
func (r *CoreRoleRepository) UpdatePermissions(id int64, permissions []byte) error {
	return r.db.Model(&core.CoreRole{}).
		Where("id = ?", id).
		Update("permission", permissions).Error
}