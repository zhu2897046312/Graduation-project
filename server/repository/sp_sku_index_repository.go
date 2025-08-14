package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
)

type SpSkuIndexRepository struct {
	*BaseRepository
}

func NewSpSkuIndexRepository(db *gorm.DB) *SpSkuIndexRepository {
	return &SpSkuIndexRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建SKU索引
func (r *SpSkuIndexRepository) Create(index *sp.SpSkuIndex) error {
	return r.db.Create(index).Error
}

// 批量创建SKU索引
func (r *SpSkuIndexRepository) BatchCreate(indices []sp.SpSkuIndex) error {
	return r.db.Create(&indices).Error
}

// 根据SKU ID获取索引
func (r *SpSkuIndexRepository) FindBySkuID(skuID uint) ([]sp.SpSkuIndex, error) {
	var indices []sp.SpSkuIndex
	err := r.db.Where("sku_id = ?", skuID).Find(&indices).Error
	return indices, err
}

// 根据产品ID获取索引
func (r *SpSkuIndexRepository) FindByProductID(productID uint) ([]sp.SpSkuIndex, error) {
	var indices []sp.SpSkuIndex
	err := r.db.Where("product_id = ?", productID).Find(&indices).Error
	return indices, err
}

// 根据属性值ID获取索引
func (r *SpSkuIndexRepository) FindByAttributeValueID(valueID uint) ([]sp.SpSkuIndex, error) {
	var indices []sp.SpSkuIndex
	err := r.db.Where("prod_attributes_value_id = ?", valueID).Find(&indices).Error
	return indices, err
}

// 删除SKU的所有索引
func (r *SpSkuIndexRepository) DeleteBySkuID(skuID uint) error {
	return r.db.Where("sku_id = ?", skuID).
		Delete(&sp.SpSkuIndex{}).Error
}

// 删除产品的所有索引
func (r *SpSkuIndexRepository) DeleteByProductID(productID uint) error {
	return r.db.Where("product_id = ?", productID).
		Delete(&sp.SpSkuIndex{}).Error
}