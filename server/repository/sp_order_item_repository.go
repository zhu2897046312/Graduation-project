package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
)

type SpOrderItemRepository struct {
	*BaseRepository
}

func NewSpOrderItemRepository(db *gorm.DB) *SpOrderItemRepository {
	return &SpOrderItemRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建订单项
func (r *SpOrderItemRepository) Create(item *sp.SpOrderItem) error {
	return r.db.Create(item).Error
}

// 批量创建订单项
func (r *SpOrderItemRepository) BatchCreate(items []*sp.SpOrderItem) error {
	return r.db.Create(&items).Error
}

// 根据订单ID获取订单项
func (r *SpOrderItemRepository) FindByOrderID(orderID uint) ([]sp.SpOrderItem, error) {
	var items []sp.SpOrderItem
	err := r.db.Where("order_id = ?", orderID).Find(&items).Error
	return items, err
}

// 根据产品ID获取订单项
func (r *SpOrderItemRepository) FindByProductID(productID uint) ([]sp.SpOrderItem, error) {
	var items []sp.SpOrderItem
	err := r.db.Where("product_id = ?", productID).Find(&items).Error
	return items, err
}

// 根据SKU ID获取订单项
func (r *SpOrderItemRepository) FindBySkuID(skuID uint) ([]sp.SpOrderItem, error) {
	var items []sp.SpOrderItem
	err := r.db.Where("sku_id = ?", skuID).Find(&items).Error
	return items, err
}

// 计算产品销售总量
func (r *SpOrderItemRepository) SumProductSales(productID uint) (int, error) {
	var total int
	err := r.db.Model(&sp.SpOrderItem{}).
		Select("SUM(quantity)").
		Where("product_id = ?", productID).
		Scan(&total).Error
	return total, err
}