package service

import (
	"errors"
	"server/models/sp"
)

type SpProdAttributesValueService struct {
	*Service
}

func NewSpProdAttributesValueService(base *Service) *SpProdAttributesValueService {
	return &SpProdAttributesValueService{Service: base}
}

// CreateAttributeValue 创建属性值
func (s *SpProdAttributesValueService) CreateAttributeValue(value *sp.SpProdAttributesValue) error {
	if value.ProdAttributesID == 0 {
		return errors.New("属性ID不能为空")
	}
	if value.Title == "" {
		return errors.New("属性值不能为空")
	}
	
	if value.SortNum == 0 {
		value.SortNum = 99
	}
	
	return s.repoFactory.GetSpProdAttributesValueRepository().Create(value)
}

// UpdateAttributeValue 更新属性值
func (s *SpProdAttributesValueService) UpdateAttributeValue(value *sp.SpProdAttributesValue) error {
	if value.ID == 0 {
		return errors.New("属性值ID不能为空")
	}
	if value.ProdAttributesID == 0 {
		return errors.New("属性ID不能为空")
	}
	if value.Title == "" {
		return errors.New("属性值不能为空")
	}
	
	return s.repoFactory.GetSpProdAttributesValueRepository().Update(value)
}

// GetValuesByAttributeID 根据属性ID获取值列表
func (s *SpProdAttributesValueService) GetValuesByAttributeID(attrID uint) ([]sp.SpProdAttributesValue, error) {
	if attrID == 0 {
		return nil, errors.New("无效的属性ID")
	}
	return s.repoFactory.GetSpProdAttributesValueRepository().FindByAttributeID(attrID)
}

// GetValueByID 根据ID获取属性值
func (s *SpProdAttributesValueService) GetValueByID(id uint) (*sp.SpProdAttributesValue, error) {
	if id == 0 {
		return nil, errors.New("无效的属性值ID")
	}
	return s.repoFactory.GetSpProdAttributesValueRepository().FindByID(id)
}

// BatchCreateAttributeValues 批量创建属性值
func (s *SpProdAttributesValueService) BatchCreateAttributeValues(values []sp.SpProdAttributesValue) error {
	if len(values) == 0 {
		return errors.New("属性值列表不能为空")
	}
	for _, value := range values {
		if value.ProdAttributesID == 0 || value.Title == "" {
			return errors.New("属性值数据不完整")
		}
	}
	return s.repoFactory.GetSpProdAttributesValueRepository().BatchCreate(values)
}

// DeleteValuesByAttributeID 删除属性的所有值
func (s *SpProdAttributesValueService) DeleteValuesByAttributeID(attrID uint) error {
	if attrID == 0 {
		return errors.New("无效的属性ID")
	}
	return s.repoFactory.GetSpProdAttributesValueRepository().DeleteByAttributeID(attrID)
}