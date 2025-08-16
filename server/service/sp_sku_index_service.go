package service

import (
	"errors"
	"server/models/sp"
)

type SpSkuIndexService struct {
	*Service
}

func NewSpSkuIndexService(base *Service) *SpSkuIndexService {
	return &SpSkuIndexService{Service: base}
}

// CreateIndex 创建SKU索引
func (s *SpSkuIndexService) CreateIndex(index *sp.SpSkuIndex) error {
	if index.SkuID == 0 {
		return errors.New("SKU ID不能为空")
	}
	if index.ProductID == 0 {
		return errors.New("商品ID不能为空")
	}
	if index.ProdAttributesValueID == 0 {
		return errors.New("属性值ID不能为空")
	}
	
	return s.repoFactory.GetSpSkuIndexRepository().Create(index)
}

// BatchCreateIndices 批量创建SKU索引
func (s *SpSkuIndexService) BatchCreateIndices(indices []sp.SpSkuIndex) error {
	if len(indices) == 0 {
		return errors.New("索引列表不能为空")
	}
	for _, idx := range indices {
		if idx.SkuID == 0 || idx.ProductID == 0 || idx.ProdAttributesValueID == 0 {
			return errors.New("索引数据不完整")
		}
	}
	return s.repoFactory.GetSpSkuIndexRepository().BatchCreate(indices)
}

// GetIndicesBySkuID 根据SKU ID获取索引
func (s *SpSkuIndexService) GetIndicesBySkuID(skuID uint) ([]sp.SpSkuIndex, error) {
	if skuID == 0 {
		return nil, errors.New("无效的SKU ID")
	}
	return s.repoFactory.GetSpSkuIndexRepository().FindBySkuID(skuID)
}

// GetIndicesByProductID 根据产品ID获取索引
func (s *SpSkuIndexService) GetIndicesByProductID(productID uint) ([]sp.SpSkuIndex, error) {
	if productID == 0 {
		return nil, errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpSkuIndexRepository().FindByProductID(productID)
}

// GetIndicesByAttributeValueID 根据属性值ID获取索引
func (s *SpSkuIndexService) GetIndicesByAttributeValueID(valueID uint) ([]sp.SpSkuIndex, error) {
	if valueID == 0 {
		return nil, errors.New("无效的属性值ID")
	}
	return s.repoFactory.GetSpSkuIndexRepository().FindByAttributeValueID(valueID)
}

// DeleteBySkuID 删除SKU的所有索引
func (s *SpSkuIndexService) DeleteBySkuID(skuID uint) error {
	if skuID == 0 {
		return errors.New("无效的SKU ID")
	}
	return s.repoFactory.GetSpSkuIndexRepository().DeleteBySkuID(skuID)
}

// DeleteByProductID 删除产品的所有索引
func (s *SpSkuIndexService) DeleteByProductID(productID uint) error {
	if productID == 0 {
		return errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpSkuIndexRepository().DeleteByProductID(productID)
}