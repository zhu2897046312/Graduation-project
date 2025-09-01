package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpSkuIndexRepository struct {
	*base.BaseRepository
}

func NewSpSkuIndexRepository(DB *gorm.DB) *SpSkuIndexRepository {
	return &SpSkuIndexRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建SKU索引
func (r *SpSkuIndexRepository) Create(index *sp.SpSkuIndex) error {
	return r.DB.Create(index).Error
}

// 批量创建SKU索引
func (r *SpSkuIndexRepository) BatchCreate(indices []sp.SpSkuIndex) error {
	return r.DB.Create(&indices).Error
}

// 根据SKU ID获取索引
func (r *SpSkuIndexRepository) FindBySkuID(skuID uint) ([]sp.SpSkuIndex, error) {
	var indices []sp.SpSkuIndex
	err := r.DB.Where("sku_id = ?", skuID).Find(&indices).Error
	return indices, err
}

// 根据产品ID获取索引
func (r *SpSkuIndexRepository) FindByProductID(productID uint) ([]sp.SpSkuIndex, error) {
	var indices []sp.SpSkuIndex
	err := r.DB.Where("product_id = ?", productID).Find(&indices).Error
	return indices, err
}

// 根据属性值ID获取索引
func (r *SpSkuIndexRepository) FindByAttributeValueID(valueID uint) ([]sp.SpSkuIndex, error) {
	var indices []sp.SpSkuIndex
	err := r.DB.Where("prod_attributes_value_id = ?", valueID).Find(&indices).Error
	return indices, err
}

// 删除SKU的所有索引
func (r *SpSkuIndexRepository) DeleteBySkuID(skuID uint) error {
	return r.DB.Where("sku_id = ?", skuID).
		Delete(&sp.SpSkuIndex{}).Error
}

// 删除产品的所有索引
func (r *SpSkuIndexRepository) DeleteByProductID(productID uint) error {
	return r.DB.Where("product_id = ?", productID).
		Delete(&sp.SpSkuIndex{}).Error
}