package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpSkuRepository struct {
	*base.BaseRepository
}

func NewSpSkuRepository(DB *gorm.DB) *SpSkuRepository {
	return &SpSkuRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建SKU
func (r *SpSkuRepository) Create(sku *sp.SpSku) error {
	return r.DB.Create(sku).Error
}

// 更新SKU
func (r *SpSkuRepository) Update(sku *sp.SpSku) error {
	return r.DB.Save(sku).Error
}

// 根据商品ID获取SKU列表
func (r *SpSkuRepository) FindByProductID(productID uint) ([]sp.SpSku, error) {
	var skus []sp.SpSku
	err := r.DB.Where("product_id = ?", productID).Find(&skus).Error
	return skus, err
}

// 根据ID获取SKU
func (r *SpSkuRepository) FindByID(id uint) (*sp.SpSku, error) {
	var sku sp.SpSku
	err := r.DB.First(&sku, id).Error
	return &sku, err
}

// 根据SKU代码获取SKU
func (r *SpSkuRepository) FindBySkuCode(code string) (*sp.SpSku, error) {
	var sku sp.SpSku
	err := r.DB.Where("sku_code = ?", code).First(&sku).Error
	return &sku, err
}

// 更新SKU库存
func (r *SpSkuRepository) UpdateStock(id uint, stock int) error {
	return r.DB.Model(&sp.SpSku{}).
		Where("id = ?", id).
		Update("stock", stock).Error
}

// 减少SKU库存
func (r *SpSkuRepository) DecrementStock(id uint, quantity int) error {
	return r.DB.Model(&sp.SpSku{}).
		Where("id = ? AND stock >= ?", id, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

// 增加SKU库存
func (r *SpSkuRepository) IncrementStock(id uint, quantity int) error {
	return r.DB.Model(&sp.SpSku{}).
		Where("id = ?", id).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}

// 设置默认SKU
func (r *SpSkuRepository) SetDefaultSku(id uint, productID uint) error {
	// 先重置所有SKU的默认状态
	if err := r.DB.Model(&sp.SpSku{}).
		Where("product_id = ?", productID).
		Update("default_show", 0).Error; err != nil {
		return err
	}
	
	// 设置当前SKU为默认
	return r.DB.Model(&sp.SpSku{}).
		Where("id = ?", id).
		Update("default_show", 1).Error
}

func (r *SpSkuRepository) DeleteByProductID(productID uint) error {
	return r.DB.Where("product_id = ?", productID).
		Delete(&sp.SpSku{}).Error
}