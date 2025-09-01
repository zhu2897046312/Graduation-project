package core

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/repository/base"
)

type CorePermissionRepository struct {
	*base.BaseRepository
}

func NewCorePermissionRepository(DB *gorm.DB) *CorePermissionRepository {
	return &CorePermissionRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建权限
func (r *CorePermissionRepository) Create(permission *core.CorePermission) error {
	return r.DB.Create(permission).Error
}

// 更新权限
func (r *CorePermissionRepository) Update(permission *core.CorePermission) error {
	return r.DB.Save(permission).Error
}

// 根据ID获取权限
func (r *CorePermissionRepository) FindByID(id int64) (*core.CorePermission, error) {
	var permission core.CorePermission
	err := r.DB.First(&permission, id).Error
	return &permission, err
}

// 根据Code获取权限
func (r *CorePermissionRepository) FindByCode(code string) (*core.CorePermission, error) {
	var permission core.CorePermission
	err := r.DB.Where("code = ?", code).First(&permission).Error
	return &permission, err
}

// // 获取权限树
// func (r *CorePermissionRepository) FindTree() ([]*core.CorePermission, error) {
// 	var permissions []*core.CorePermission
// 	err := r.DB.Order("id ASC").Find(&permissions).Error
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	// 构建树形结构
// 	permMap := make(map[int64]*core.CorePermission)
// 	for i := range permissions {
// 		permMap[permissions[i].ID] = permissions[i]
// 	}
	
// 	var tree []*core.CorePermission
// 	for i := range permissions {
// 		if permissions[i].Pid == 0 {
// 			tree = append(tree, permissions[i])
// 		} else {
// 			if parent, ok := permMap[permissions[i].Pid]; ok {
// 				parent.Children = append(parent.Children, permissions[i])
// 			}
// 		}
// 	}
	
// 	return tree, nil
// }