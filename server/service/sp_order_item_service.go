package service

import (
	"errors"
	"server/models/sp"
)

type SpOrderItemService struct {
	*Service
}

func NewSpOrderItemService(base *Service) *SpOrderItemService {
	return &SpOrderItemService{Service: base}
}

// CreateOrderItem 创建订单项
func (s *SpOrderItemService) CreateOrderItem(item *sp.SpOrderItem) error {
	if item.OrderID == 0 {
		return errors.New("订单ID不能为空")
	}
	if item.ProductID == 0 {
		return errors.New("产品ID不能为空")
	}
	if item.Quantity <= 0 {
		return errors.New("数量必须大于0")
	}
	return s.repoFactory.GetSpOrderItemRepository().Create(item)
}

// BatchCreateOrderItems 批量创建订单项
func (s *SpOrderItemService) BatchCreateOrderItems(items []*sp.SpOrderItem) error {
	if len(items) == 0 {
		return errors.New("订单项列表不能为空")
	}
	
	for _, item := range items {
		if item.OrderID == 0 || item.ProductID == 0 || item.Quantity <= 0 {
			return errors.New("订单项数据不完整")
		}
	}
	
	return s.repoFactory.GetSpOrderItemRepository().BatchCreate(items)
}

// GetItemsByOrderID 根据订单ID获取订单项
func (s *SpOrderItemService) GetItemsByOrderID(orderID uint) ([]sp.SpOrderItem, error) {
	if orderID == 0 {
		return nil, errors.New("无效的订单ID")
	}
	return s.repoFactory.GetSpOrderItemRepository().FindByOrderID(orderID)
}

// GetItemsByProductID 根据产品ID获取订单项
func (s *SpOrderItemService) GetItemsByProductID(productID uint) ([]sp.SpOrderItem, error) {
	if productID == 0 {
		return nil, errors.New("无效的产品ID")
	}
	return s.repoFactory.GetSpOrderItemRepository().FindByProductID(productID)
}

// GetItemsBySkuID 根据SKU ID获取订单项
func (s *SpOrderItemService) GetItemsBySkuID(skuID uint) ([]sp.SpOrderItem, error) {
	if skuID == 0 {
		return nil, errors.New("无效的SKU ID")
	}
	return s.repoFactory.GetSpOrderItemRepository().FindBySkuID(skuID)
}

// CalculateProductSales 计算产品销售总量
func (s *SpOrderItemService) CalculateProductSales(productID uint) (int, error) {
	if productID == 0 {
		return 0, errors.New("无效的产品ID")
	}
	return s.repoFactory.GetSpOrderItemRepository().SumProductSales(productID)
}