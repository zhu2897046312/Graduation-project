package service

import (
	"errors"
	"server/models/sp"
	"server/models/common"
)

type SpProductPropertyService struct {
	*Service
}

func NewSpProductPropertyService(base *Service) *SpProductPropertyService {
	return &SpProductPropertyService{Service: base}
}

// CreateProperty 创建商品属性
func (s *SpProductPropertyService) CreateProperty(property *sp.SpProductProperty) error {
	if property.ProductID == 0 {
		return errors.New("商品ID不能为空")
	}
	if property.Title == "" {
		return errors.New("属性名称不能为空")
	}
	if property.Value == "" {
		return errors.New("属性值不能为空")
	}
	
	if property.SortNum == 0 {
		property.SortNum = 99
	}
	
	return s.repoFactory.GetSpProductPropertyRepository().Create(property)
}

// UpdateProperty 更新商品属性
func (s *SpProductPropertyService) UpdateProperty(property *sp.SpProductProperty) error {
	if property.ID == 0 {
		return errors.New("属性ID不能为空")
	}
	if property.Title == "" {
		return errors.New("属性名称不能为空")
	}
	if property.Value == "" {
		return errors.New("属性值不能为空")
	}
	
	return s.repoFactory.GetSpProductPropertyRepository().Update(property)
}

// GetPropertiesByProductID 根据商品ID获取属性
func (s *SpProductPropertyService) GetPropertiesByProductID(productID common.MyID) ([]sp.SpProductProperty, error) {
	if productID == 0 {
		return nil, errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpProductPropertyRepository().FindByProductID(productID)
}

// BatchCreateProperties 批量创建商品属性
func (s *SpProductPropertyService) BatchCreateProperties(properties []sp.SpProductProperty) error {
	if len(properties) == 0 {
		return errors.New("属性列表不能为空")
	}
	for _, prop := range properties {
		if prop.ProductID == 0 || prop.Title == "" || prop.Value == "" {
			return errors.New("属性数据不完整")
		}
	}
	return s.repoFactory.GetSpProductPropertyRepository().BatchCreate(properties)
}

// DeletePropertiesByProductID 删除商品的所有属性
func (s *SpProductPropertyService) DeletePropertiesByProductID(productID common.MyID) error {
	if productID == 0 {
		return errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpProductPropertyRepository().DeleteByProductID(productID)
}