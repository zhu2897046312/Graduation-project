package service

import (
	"errors"
	"server/models/core"
	"server/models/common"
)

type CoreAdminRoleIndexService struct {
	*Service
}

func NewCoreAdminRoleIndexService(base *Service) *CoreAdminRoleIndexService {
	return &CoreAdminRoleIndexService{Service: base}
}

// CreateAdminRole 创建管理员-角色关联
func (s *CoreAdminRoleIndexService) CreateAdminRole(index *core.CoreAdminRoleIndex) error {
	if index.AdminID <= 0 {
		return errors.New("无效的管理员ID")
	}
	if index.RoleID <= 0 {
		return errors.New("无效的角色ID")
	}
	
	return s.repoFactory.GetCoreAdminRoleIndexRepository().Create(index)
}

// DeleteAdminRole 删除管理员-角色关联
func (s *CoreAdminRoleIndexService) DeleteAdminRole(adminID, roleID common.MyID) error {
	if adminID <= 0 {
		return errors.New("无效的管理员ID")
	}
	if roleID <= 0 {
		return errors.New("无效的角色ID")
	}
	
	return s.repoFactory.GetCoreAdminRoleIndexRepository().Delete(adminID, roleID)
}

// GetRolesByAdminID 根据管理员ID获取角色关联
func (s *CoreAdminRoleIndexService) GetRolesByAdminID(adminID common.MyID) ([]core.CoreAdminRoleIndex, error) {
	if adminID <= 0 {
		return nil, errors.New("无效的管理员ID")
	}
	return s.repoFactory.GetCoreAdminRoleIndexRepository().FindByAdminID(adminID)
}

// GetAdminsByRoleID 根据角色ID获取管理员关联
func (s *CoreAdminRoleIndexService) GetAdminsByRoleID(roleID common.MyID) ([]core.CoreAdminRoleIndex, error) {
	if roleID <= 0 {
		return nil, errors.New("无效的角色ID")
	}
	return s.repoFactory.GetCoreAdminRoleIndexRepository().FindByRoleID(roleID)
}

// DeleteAllRolesByAdminID 删除管理员的所有角色关联
func (s *CoreAdminRoleIndexService) DeleteAllRolesByAdminID(adminID common.MyID) error {
	if adminID <= 0 {
		return errors.New("无效的管理员ID")
	}
	return s.repoFactory.GetCoreAdminRoleIndexRepository().DeleteByAdminID(adminID)
}