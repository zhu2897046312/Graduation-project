package service

import (
	"errors"
	"server/models/sp"
	"time"
)

type SpProdAttributesService struct {
	*Service
}

func NewSpProdAttributesService(base *Service) *SpProdAttributesService {
	return &SpProdAttributesService{Service: base}
}

// CreateAttribute 创建商品属性
func (s *SpProdAttributesService) CreateAttribute(attr *sp.SpProdAttributes) error {
	if attr.Title == "" {
		return errors.New("属性名称不能为空")
	}
	
	if attr.SortNum == 0 {
		attr.SortNum = 99
	}
	
	attr.CreatedTime = time.Now()
	attr.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpProdAttributesRepository().Create(attr)
}

// UpdateAttribute 更新商品属性
func (s *SpProdAttributesService) UpdateAttribute(attr *sp.SpProdAttributes) error {
	if attr.ID == 0 {
		return errors.New("属性ID不能为空")
	}
	if attr.Title == "" {
		return errors.New("属性名称不能为空")
	}
	
	// 保留原始创建时间
	existing, err := s.repoFactory.GetSpProdAttributesRepository().FindByID(attr.ID)
	if err != nil {
		return errors.New("属性不存在")
	}
	
	attr.CreatedTime = existing.CreatedTime
	attr.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpProdAttributesRepository().Update(attr)
}

// GetAttributeByID 根据ID获取属性
func (s *SpProdAttributesService) GetAttributeByID(id uint) (*sp.SpProdAttributes, error) {
	if id == 0 {
		return nil, errors.New("无效的属性ID")
	}
	return s.repoFactory.GetSpProdAttributesRepository().FindByID(id)
}

// GetAllAttributes 获取所有属性
func (s *SpProdAttributesService) GetAllAttributes() ([]sp.SpProdAttributes, error) {
	return s.repoFactory.GetSpProdAttributesRepository().FindAll()
}

// GetAttributesByPage 分页获取属性
func (s *SpProdAttributesService) GetAttributesByPage(title string, page, pageSize int) ([]sp.SpProdAttributes, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20 // 默认每页20条
	}
	
	return s.repoFactory.GetSpProdAttributesRepository().FindByPage(title, page, pageSize)
}

// UpdateAttributeSortNum 更新属性排序
func (s *SpProdAttributesService) UpdateAttributeSortNum(id uint, sortNum uint16) error {
	if sortNum > 9999 {
		return errors.New("排序值不能超过9999")
	}
	return s.repoFactory.GetSpProdAttributesRepository().UpdateSortNum(id, sortNum)
}

// DeleteAttribute 删除属性
func (s *SpProdAttributesService) DeleteAttribute(id uint) error {
	if id == 0 {
		return errors.New("无效的属性ID")
	}
	return s.repoFactory.GetSpProdAttributesRepository().Delete(id)
}