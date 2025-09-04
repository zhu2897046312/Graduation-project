package service

import (
	"errors"
	"server/models/core"
)

type CoreDeptService struct {
	*Service
}

func NewCoreDeptService(base *Service) *CoreDeptService {
	return &CoreDeptService{Service: base}
}

// CreateDept 创建部门
func (s *CoreDeptService) CreateDept(dept *core.CoreDept) error {
	if dept.Thumb == "" {
		return errors.New("部门名称不能为空")
	}
	
	// 设置默认排序值
	if dept.SortNum == 0 {
		dept.SortNum = 999
	}
	
	return s.repoFactory.GetCoreDeptRepository().Create(dept)
}

// UpdateDept 更新部门
func (s *CoreDeptService) UpdateDept(dept *core.CoreDept) error {
	if dept.ID <= 0 {
		return errors.New("无效的部门ID")
	}
	if dept.Thumb == "" {
		return errors.New("部门名称不能为空")
	}
	
	// 检查部门是否存在
	existing, err := s.repoFactory.GetCoreDeptRepository().FindByID(dept.ID)
	if err != nil {
		return errors.New("部门不存在")
	}
	
	// 保留原始创建时间
	dept.CreatedTime = existing.CreatedTime
	
	return s.repoFactory.GetCoreDeptRepository().Update(dept)
}

// GetDeptByID 根据ID获取部门
func (s *CoreDeptService) GetDeptByID(id int64) (*core.CoreDept, error) {
	if id <= 0 {
		return nil, errors.New("无效的部门ID")
	}
	return s.repoFactory.GetCoreDeptRepository().FindByID(id)
}

// GetSubDepts 根据父ID获取子部门
func (s *CoreDeptService) GetSubDepts(pid int64) ([]*core.CoreDept, error) {
	if pid < 0 {
		return nil, errors.New("无效的父部门ID")
	}
	return s.repoFactory.GetCoreDeptRepository().FindByPid(pid)
}

// DeleteDept 删除部门
func (s *CoreDeptService) DeleteDept(id int64) error {
	if id <= 0 {
		return errors.New("无效的部门ID")
	}
	
	// 检查部门下是否有管理员
	admins, err := s.repoFactory.GetCoreAdminRepository().FindByDeptID(id)
	if err == nil && len(admins) > 0 {
		return errors.New("请先移除部门下的管理员")
	}
	
	// 检查是否有子部门
	subDepts, err := s.GetSubDepts(id)
	if err == nil && len(subDepts) > 0 {
		return errors.New("请先删除子部门")
	}
	
	return s.repoFactory.GetCoreDeptRepository().Delete(id)
}

func (s *CoreDeptService) List(pid int64) ([]core.CoreDept, int64, error) {
	return s.repoFactory.GetCoreDeptRepository().List(pid)
}

func (s *CoreDeptService) GetAll() ([]*core.CoreDept,int64, error) {
	return s.repoFactory.GetCoreDeptRepository().FindAll()
}