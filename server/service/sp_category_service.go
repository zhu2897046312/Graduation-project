package service

import (
	"errors"
	"server/models/sp"
	"time"
)

type SpCategoryService struct {
	*Service
}

func NewSpCategoryService(base *Service) *SpCategoryService {
	return &SpCategoryService{Service: base}
}

// CreateCategory 创建分类
func (s *SpCategoryService) CreateCategory(category *sp.SpCategory) error {
	if category.Title == "" {
		return errors.New("分类名称不能为空")
	}
	
	if category.State != 0 && category.State != 1 {
		return errors.New("无效的状态值")
	}
	
	// 设置默认值
	if category.SortNum == 0 {
		category.SortNum = 99
	}
	
	category.CreatedTime = time.Now()
	category.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpCategoryRepository().Create(category)
}

// UpdateCategory 更新分类
func (s *SpCategoryService) UpdateCategory(category *sp.SpCategory) error {
	if category.Title == "" {
		return errors.New("分类名称不能为空")
	}
	
	if category.State != 0 && category.State != 1 {
		return errors.New("无效的状态值")
	}
	
	// 保留原始创建时间
	existing, err := s.repoFactory.GetSpCategoryRepository().FindByID(category.ID)
	if err != nil {
		return errors.New("分类不存在")
	}
	
	category.CreatedTime = existing.CreatedTime
	category.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpCategoryRepository().Update(category)
}

// GetCategoryByID 根据ID获取分类
func (s *SpCategoryService) GetCategoryByID(id uint) (*sp.SpCategory, error) {
	if id == 0 {
		return nil, errors.New("无效的分类ID")
	}
	return s.repoFactory.GetSpCategoryRepository().FindByID(id)
}

// GetAllCategories 获取所有分类
func (s *SpCategoryService) GetAllCategories() ([]*sp.SpCategory, error) {
	return s.repoFactory.GetSpCategoryRepository().FindAll()
}

// GetCategoriesByPid 根据父ID获取子分类
func (s *SpCategoryService) GetCategoriesByPid(pid uint) ([]*sp.SpCategory, error) {
	return s.repoFactory.GetSpCategoryRepository().FindByPid(pid)
}

// UpdateCategoryState 更新分类状态
func (s *SpCategoryService) UpdateCategoryState(id uint, state uint8) error {
	if state != 0 && state != 1 {
		return errors.New("无效的状态值")
	}
	return s.repoFactory.GetSpCategoryRepository().UpdateState(id, state)
}

// UpdateCategorySortNum 更新分类排序
func (s *SpCategoryService) UpdateCategorySortNum(id uint, sortNum uint16) error {
	if sortNum > 9999 {
		return errors.New("排序值不能超过9999")
	}
	return s.repoFactory.GetSpCategoryRepository().UpdateSortNum(id, sortNum)
}