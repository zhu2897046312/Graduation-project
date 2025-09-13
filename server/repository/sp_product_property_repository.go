package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/models/common"
)

type SpProductPropertyRepository struct {
	*BaseRepository
}

func NewSpProductPropertyRepository(db *gorm.DB) *SpProductPropertyRepository {
	return &SpProductPropertyRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建商品属性
func (r *SpProductPropertyRepository) Create(property *sp.SpProductProperty) error {
	return r.db.Create(property).Error
}

// 更新商品属性
func (r *SpProductPropertyRepository) Update(property *sp.SpProductProperty) error {
	return r.db.Save(property).Error
}

// 根据商品ID获取属性
func (r *SpProductPropertyRepository) FindByProductID(productID common.MyID) ([]sp.SpProductProperty, error) {
	var properties []sp.SpProductProperty
	err := r.db.Where("product_id = ?", productID).
		Order("sort_num ASC").
		Find(&properties).Error
	return properties, err
}

// 批量创建商品属性
func (r *SpProductPropertyRepository) BatchCreate(properties []sp.SpProductProperty) error {
	return r.db.Create(&properties).Error
}

// 删除商品的所有属性
func (r *SpProductPropertyRepository) DeleteByProductID(productID common.MyID) error {
	return r.db.Where("product_id = ?", productID).
		Delete(&sp.SpProductProperty{}).Error
}