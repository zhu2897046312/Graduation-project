package repository

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/models/common"
)

type CoreAdminRoleIndexRepository struct {
	*BaseRepository
}

func NewCoreAdminRoleIndexRepository(db *gorm.DB) *CoreAdminRoleIndexRepository {
	return &CoreAdminRoleIndexRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建管理员-角色关联
func (r *CoreAdminRoleIndexRepository) Create(index *core.CoreAdminRoleIndex) error {
	return r.db.Create(index).Error
}

// 删除管理员-角色关联
func (r *CoreAdminRoleIndexRepository) Delete(adminID, roleID common.MyID) error {
	return r.db.Where("admin_id = ? AND role_id = ?", adminID, roleID).
		Delete(&core.CoreAdminRoleIndex{}).Error
}

// 根据管理员ID获取角色关联
func (r *CoreAdminRoleIndexRepository) FindByAdminID(adminID common.MyID) ([]core.CoreAdminRoleIndex, error) {
	var indices []core.CoreAdminRoleIndex
	err := r.db.Where("admin_id = ?", adminID).Find(&indices).Error
	return indices, err
}

// 根据角色ID获取管理员关联
func (r *CoreAdminRoleIndexRepository) FindByRoleID(roleID common.MyID) ([]core.CoreAdminRoleIndex, error) {
	var indices []core.CoreAdminRoleIndex
	err := r.db.Where("role_id = ?", roleID).Find(&indices).Error
	return indices, err
}

// 删除管理员的所有角色关联
func (r *CoreAdminRoleIndexRepository) DeleteByAdminID(adminID common.MyID) error {
	return r.db.Where("admin_id = ?", adminID).
		Delete(&core.CoreAdminRoleIndex{}).Error
}