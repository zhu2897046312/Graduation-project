package service

import (
	"errors"
	"server/models/sp"
	"server/models/common"
)

type SpSkuService struct {
	*Service
}

func NewSpSkuService(base *Service) *SpSkuService {
	return &SpSkuService{Service: base}
}

// CreateSku 创建SKU
func (s *SpSkuService) CreateSku(sku *sp.SpSku) error {
	if sku.ProductID == 0 {
		return errors.New("商品ID不能为空")
	}
	if sku.SkuCode == "" {
		return errors.New("SKU编码不能为空")
	}
	if sku.Price <= 0 {
		return errors.New("SKU价格必须大于0")
	}
	
	return s.repoFactory.GetSpSkuRepository().Create(sku)
}

// UpdateSku 更新SKU
func (s *SpSkuService) UpdateSku(sku *sp.SpSku) error {
	if sku.ID == 0 {
		return errors.New("SKU ID不能为空")
	}
	if sku.Price <= 0 {
		return errors.New("SKU价格必须大于0")
	}
	
	return s.repoFactory.GetSpSkuRepository().Update(sku)
}

// GetSkusByProductID 根据商品ID获取SKU列表
func (s *SpSkuService) GetSkusByProductID(productID common.MyID) ([]sp.SpSku, error) {
	if productID == 0 {
		return nil, errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpSkuRepository().FindByProductID(productID)
}

// GetSkuByID 根据ID获取SKU
func (s *SpSkuService) GetSkuByID(id common.MyID) (*sp.SpSku, error) {
	if id == 0 {
		return nil, errors.New("无效的SKU ID")
	}
	return s.repoFactory.GetSpSkuRepository().FindByID(id)
}

// GetSkuByCode 根据SKU代码获取SKU
func (s *SpSkuService) GetSkuByCode(code string) (*sp.SpSku, error) {
	if code == "" {
		return nil, errors.New("SKU编码不能为空")
	}
	return s.repoFactory.GetSpSkuRepository().FindBySkuCode(code)
}

// UpdateStock 更新SKU库存
func (s *SpSkuService) UpdateStock(id common.MyID, stock int) error {
	if stock < 0 {
		return errors.New("库存不能为负数")
	}
	return s.repoFactory.GetSpSkuRepository().UpdateStock(id, stock)
}

// DecrementStock 减少SKU库存
func (s *SpSkuService) DecrementStock(id common.MyID, quantity int) error {
	if quantity <= 0 {
		return errors.New("减少数量必须大于0")
	}
	return s.repoFactory.GetSpSkuRepository().DecrementStock(id, quantity)
}

// IncrementStock 增加SKU库存
func (s *SpSkuService) IncrementStock(id common.MyID, quantity int) error {
	if quantity <= 0 {
		return errors.New("增加数量必须大于0")
	}
	return s.repoFactory.GetSpSkuRepository().IncrementStock(id, quantity)
}

// SetDefaultSku 设置默认SKU
func (s *SpSkuService) SetDefaultSku(id common.MyID, productID common.MyID) error {
	if id == 0 {
		return errors.New("无效的SKU ID")
	}
	if productID == 0 {
		return errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpSkuRepository().SetDefaultSku(id, productID)
}

func (s *SpSkuService) DeleteByProductID(productID common.MyID) error  {
	if productID == 0 {
		return errors.New("无效的SKU ID")
	}
	sku , _ := s.repoFactory.GetSpSkuRepository().FindByProductID(productID)
	if sku != nil {
		s.repoFactory.GetSpSkuRepository().DeleteByProductID(productID)
	}
	return nil
}