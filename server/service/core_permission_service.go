package service

import (
	"errors"
	"server/models/core"
)

type CorePermissionService struct {
	*Service
}

func NewCorePermissionService(base *Service) *CorePermissionService {
	return &CorePermissionService{Service: base}
}

// CreatePermission 创建权限
func (s *CorePermissionService) CreatePermission(permission *core.CorePermission) error {
	if permission.Title == "" {
		return errors.New("权限名称不能为空")
	}
	if permission.Code == "" {
		return errors.New("权限代码不能为空")
	}
	
	// 检查权限代码是否已存在
	existing, _ := s.repoFactory.GetCorePermissionRepository().FindByCode(permission.Code)
	if existing != nil {
		return errors.New("权限代码已存在")
	}
	
	return s.repoFactory.GetCorePermissionRepository().Create(permission)
}

// UpdatePermission 更新权限
func (s *CorePermissionService) UpdatePermission(permission *core.CorePermission) error {
	if permission.ID <= 0 {
		return errors.New("无效的权限ID")
	}
	if permission.Title == "" {
		return errors.New("权限名称不能为空")
	}
	if permission.Code == "" {
		return errors.New("权限代码不能为空")
	}
	
	// 检查权限是否存在
	existing, err := s.repoFactory.GetCorePermissionRepository().FindByID(permission.ID)
	if err != nil {
		return errors.New("权限不存在")
	}
	
	// 检查代码冲突
	if existing.Code != permission.Code {
		codeCheck, _ := s.repoFactory.GetCorePermissionRepository().FindByCode(permission.Code)
		if codeCheck != nil {
			return errors.New("权限代码已存在")
		}
	}
	
	return s.repoFactory.GetCorePermissionRepository().Update(permission)
}

// GetPermissionByID 根据ID获取权限
func (s *CorePermissionService) GetPermissionByID(id int64) (*core.CorePermission, error) {
	if id <= 0 {
		return nil, errors.New("无效的权限ID")
	}
	return s.repoFactory.GetCorePermissionRepository().FindByID(id)
}

// GetPermissionByCode 根据Code获取权限
func (s *CorePermissionService) GetPermissionByCode(code string) (*core.CorePermission, error) {
	if code == "" {
		return nil, errors.New("权限代码不能为空")
	}
	return s.repoFactory.GetCorePermissionRepository().FindByCode(code)
}