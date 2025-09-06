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

func (s *SpCategoryService) GetCategoryByCode(code string) (*sp.SpCategory, error) {
	if code == "" {
		return nil, errors.New("分类代码不能为空")
	}
	return s.repoFactory.GetSpCategoryRepository().FindByCode(code)
}

// GetParents 根据分类代码获取所有父级分类（包括自身），按层级顺序从根分类到当前分类
func (s *SpCategoryService) GetParents(code string) ([]sp.SpCategory, error) {
	if code == "" {
		return nil, errors.New("分类代码不能为空")
	}

	// 1. 首先根据code获取当前分类
	currentCategory, err := s.repoFactory.GetSpCategoryRepository().FindByCode(code)
	if err != nil {
		return nil, err
	}

	// 如果找不到分类或分类已删除/禁用，返回空数组
	if currentCategory == nil || currentCategory.DeletedTime != nil || currentCategory.State != 1 {
		return []sp.SpCategory{}, nil
	}

	// 2. 创建结果数组并添加当前分类
	out := []sp.SpCategory{*currentCategory}

	// 3. 递归获取父级分类
	pid := currentCategory.Pid
	num := 10 // 防止无限循环的安全计数器

	for num >= 0 && pid > 0 {
		num--

		// 获取父级分类
		parentCategory, err := s.repoFactory.GetSpCategoryRepository().FindByID(pid)
		if err != nil {
			return nil, err
		}

		// 如果父级分类不存在或已删除/禁用，则停止循环
		if parentCategory == nil || parentCategory.DeletedTime != nil || parentCategory.State != 1 {
			break
		}

		// 添加父级分类到结果数组
		out = append(out, *parentCategory)
		
		// 继续向上查找
		pid = parentCategory.Pid
	}

	// 4. 反转数组，使顺序从根分类到当前分类
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out, nil
}