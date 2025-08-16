package service

import (
	"errors"
	"server/models/core"
)

type CoreAdminService struct {
	*Service
}

func NewCoreAdminService(base *Service) *CoreAdminService {
	return &CoreAdminService{Service: base}
}

// CreateAdmin 创建管理员
func (s *CoreAdminService) CreateAdmin(admin *core.CoreAdmin) error {
	// 验证必填字段
	if admin.Account == "" {
		return errors.New("管理员账号不能为空")
	}
	if admin.Pwd == "" {
		return errors.New("管理员密码不能为空")
	}
	
	// 检查账号是否已存在
	existing, _ := s.repoFactory.GetCoreAdminRepository().FindByAccount(admin.Account)
	if existing != nil {
		return errors.New("管理员账号已存在")
	}
	
	return s.repoFactory.GetCoreAdminRepository().Create(admin)
}

// UpdateAdmin 更新管理员
func (s *CoreAdminService) UpdateAdmin(admin *core.CoreAdmin) error {
	if admin.ID <= 0 {
		return errors.New("无效的管理员ID")
	}
	
	// 检查管理员是否存在
	existing, err := s.repoFactory.GetCoreAdminRepository().FindByID(admin.ID)
	if err != nil {
		return errors.New("管理员不存在")
	}
	
	// 保留原始密码（如果未修改）
	if admin.Pwd == "" {
		admin.Pwd = existing.Pwd
	}
	
	return s.repoFactory.GetCoreAdminRepository().Update(admin)
}

// GetAdminByAccount 根据账号获取管理员
func (s *CoreAdminService) GetAdminByAccount(account string) (*core.CoreAdmin, error) {
	if account == "" {
		return nil, errors.New("账号不能为空")
	}
	return s.repoFactory.GetCoreAdminRepository().FindByAccount(account)
}

// GetAdminByID 根据ID获取管理员
func (s *CoreAdminService) GetAdminByID(id int64) (*core.CoreAdmin, error) {
	if id <= 0 {
		return nil, errors.New("无效的管理员ID")
	}
	return s.repoFactory.GetCoreAdminRepository().FindByID(id)
}

// GetAdminsByDeptID 根据部门ID获取管理员列表
func (s *CoreAdminService) GetAdminsByDeptID(deptID int64) ([]core.CoreAdmin, error) {
	if deptID <= 0 {
		return nil, errors.New("无效的部门ID")
	}
	return s.repoFactory.GetCoreAdminRepository().FindByDeptID(deptID)
}

// UpdateAdminStatus 更新管理员状态
func (s *CoreAdminService) UpdateAdminStatus(id int64, status int8) error {
	if id <= 0 {
		return errors.New("无效的管理员ID")
	}
	if status < 0 || status > 2 {
		return errors.New("无效的状态值")
	}
	
	return s.repoFactory.GetCoreAdminRepository().UpdateStatus(id, status)
}

// UpdateAdminPassword 更新管理员密码
func (s *CoreAdminService) UpdateAdminPassword(id int64, newPwd string) error {
	if id <= 0 {
		return errors.New("无效的管理员ID")
	}
	if len(newPwd) < 6 {
		return errors.New("密码长度不能少于6位")
	}
	
	return s.repoFactory.GetCoreAdminRepository().UpdatePassword(id, newPwd)
}