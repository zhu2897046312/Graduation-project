package service

import (
	"errors"
	"server/models/cms"
)

type CmsCategoryService struct {
	*Service
}

func NewCmsCategoryService(base *Service) *CmsCategoryService {
	return &CmsCategoryService{Service: base}
}

// CreateCategory 创建分类
func (s *CmsCategoryService) CreateCategory(category *cms.CmsCategory) error {
	// 验证必填字段
	if category.Title == "" {
		return errors.New("分类名称不能为空")
	}
	
	// 设置默认排序值
	if category.SortNum == 0 {
		category.SortNum = 999 // 默认排序值
	}
	
	return s.repoFactory.GetCmsCategoryRepository().Create(category)
}

// UpdateCategory 更新分类
func (s *CmsCategoryService) UpdateCategory(category *cms.CmsCategory) error {
	// 验证必填字段
	if category.Title == "" {
		return errors.New("分类名称不能为空")
	}
	
	// 检查分类是否存在
	_, err := s.repoFactory.GetCmsCategoryRepository().FindByID(category.ID)
	if err != nil {
		return errors.New("分类不存在")
	}
	
	// 保留原始创建时间（如果模型中有该字段）
	// category.CreatedAt = existing.CreatedAt
	
	return s.repoFactory.GetCmsCategoryRepository().Update(category)
}

// GetCategoryByID 根据ID获取分类
func (s *CmsCategoryService) GetCategoryByID(id int64) (*cms.CmsCategory, error) {
	if id <= 0 {
		return nil, errors.New("无效的分类ID")
	}
	return s.repoFactory.GetCmsCategoryRepository().FindByID(id)
}

// GetCategoriesByParentID 根据父ID获取子分类
func (s *CmsCategoryService) GetCategoriesByParentID(parentID int64) ([]cms.CmsCategory, error) {
	if parentID < 0 {
		return nil, errors.New("无效的父分类ID")
	}
	return s.repoFactory.GetCmsCategoryRepository().FindByParentID(parentID)
}

// GetCategoriesByType 根据类型获取分类
func (s *CmsCategoryService) GetCategoriesByType(categoryType int8) ([]cms.CmsCategory, error) {
	if categoryType < 0 {
		return nil, errors.New("无效的分类类型")
	}
	return s.repoFactory.GetCmsCategoryRepository().FindByType(categoryType)
}

// GetAllCategories 获取所有分类
func (s *CmsCategoryService) GetAllCategories() ([]cms.CmsCategory, error) {
	return s.repoFactory.GetCmsCategoryRepository().FindAll()
}

// ListCategories 分页获取分类
func (s *CmsCategoryService) ListCategories(page, pageSize int) ([]cms.CmsCategory, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	
	return s.repoFactory.GetCmsCategoryRepository().ListWithPagination(page, pageSize)
}

// UpdateCategorySort 更新分类排序
func (s *CmsCategoryService) UpdateCategorySort(id int64, sortNum int) error {
	if id <= 0 {
		return errors.New("无效的分类ID")
	}
	
	if sortNum < 0 {
		return errors.New("排序值不能为负数")
	}
	
	return s.repoFactory.GetCmsCategoryRepository().UpdateSortNum(id, sortNum)
}

// DeleteCategory 删除分类
func (s *CmsCategoryService) DeleteCategory(id int64) error {
	if id <= 0 {
		return errors.New("无效的分类ID")
	}
	
	// 检查分类是否存在
	_, err := s.GetCategoryByID(id)
	if err != nil {
		return err
	}
	
	// 检查是否有子分类（可选）
	children, err := s.GetCategoriesByParentID(id)
	if err == nil && len(children) > 0 {
		return errors.New("请先删除子分类")
	}
	
	return s.repoFactory.GetCmsCategoryRepository().Delete(id)
}