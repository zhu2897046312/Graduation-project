package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpProductPropertyRepository struct {
	*base.BaseRepository
}

func NewSpProductPropertyRepository(DB *gorm.DB) *SpProductPropertyRepository {
	return &SpProductPropertyRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建商品属性
func (r *SpProductPropertyRepository) Create(property *sp.SpProductProperty) error {
	return r.DB.Create(property).Error
}

// 更新商品属性
func (r *SpProductPropertyRepository) Update(property *sp.SpProductProperty) error {
	return r.DB.Save(property).Error
}

// 根据商品ID获取属性
func (r *SpProductPropertyRepository) FindByProductID(productID uint) ([]sp.SpProductProperty, error) {
	var properties []sp.SpProductProperty
	err := r.DB.Where("product_id = ?", productID).
		Order("sort_num ASC").
		Find(&properties).Error
	return properties, err
}

// 批量创建商品属性
func (r *SpProductPropertyRepository) BatchCreate(properties []sp.SpProductProperty) error {
	return r.DB.Create(&properties).Error
}

// 删除商品的所有属性
func (r *SpProductPropertyRepository) DeleteByProductID(productID uint) error {
	return r.DB.Where("product_id = ?", productID).
		Delete(&sp.SpProductProperty{}).Error
}