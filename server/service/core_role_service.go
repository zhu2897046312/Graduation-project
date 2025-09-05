package service

import (
	"errors"
	"server/models/core"
)

type CoreRoleService struct {
	*Service
}

func NewCoreRoleService(base *Service) *CoreRoleService {
	return &CoreRoleService{Service: base}
}

// CreateRole 创建角色
func (s *CoreRoleService) CreateRole(role *core.CoreRole) error {
	if role.RoleName == "" {
		return errors.New("角色名称不能为空")
	}
	
	return s.repoFactory.GetCoreRoleRepository().Create(role)
}

// UpdateRole 更新角色
func (s *CoreRoleService) UpdateRole(role *core.CoreRole) error {
	if role.ID <= 0 {
		return errors.New("无效的角色ID")
	}
	
	// 检查角色是否存在
	_, err := s.repoFactory.GetCoreRoleRepository().FindByID(role.ID)
	if err != nil {
		return errors.New("角色不存在")
	}
	
	return s.repoFactory.GetCoreRoleRepository().Update(role)
}

// GetRoleByID 根据ID获取角色
func (s *CoreRoleService) GetRoleByID(id int64) (*core.CoreRole, error) {
	if id <= 0 {
		return nil, errors.New("无效的角色ID")
	}
	return s.repoFactory.GetCoreRoleRepository().FindByID(id)
}

// GetAllRoles 获取所有角色
func (s *CoreRoleService) GetAllRoles() ([]core.CoreRole, error) {
	return s.repoFactory.GetCoreRoleRepository().FindAll()
}

// UpdateRoleStatus 更新角色状态
func (s *CoreRoleService) UpdateRoleStatus(id int64, status int8) error {
	if id <= 0 {
		return errors.New("无效的角色ID")
	}
	if status < 0 || status > 2 {
		return errors.New("无效的状态值")
	}
	
	return s.repoFactory.GetCoreRoleRepository().UpdateStatus(id, status)
}

// UpdateRolePermissions 更新角色权限
func (s *CoreRoleService) UpdateRolePermissions(id int64, permissions []byte) error {
	if id <= 0 {
		return errors.New("无效的角色ID")
	}
	if len(permissions) == 0 {
		return errors.New("权限数据不能为空")
	}
	
	return s.repoFactory.GetCoreRoleRepository().UpdatePermissions(id, permissions)
}

func (s *CoreRoleService) List(page int, pageSize int) ([]core.CoreRole, int64, error) {
	return s.repoFactory.GetCoreRoleRepository().List(page, pageSize)
}

func (s *CoreRoleService) GetAllRolesByAdminID(adminID int64) ([]core.CoreRole, error) {
	
 return s.repoFactory.GetCoreRoleRepository().FindByAdminID(adminID)
}

func (s *CoreRoleService) DeleteRole(ID int64) ( error) {
	if ID <= 0 {
		return errors.New("无效的角色ID")
	}
	_, err := s.repoFactory.GetCoreRoleRepository().FindByID(ID)
	if err != nil {
		return errors.New("角色不存在")
	}
	return s.repoFactory.GetCoreRoleRepository().Delete(ID)
}	