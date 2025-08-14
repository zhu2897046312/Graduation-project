package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
)

type SpProdAttributesRepository struct {
	*BaseRepository
}

func NewSpProdAttributesRepository(db *gorm.DB) *SpProdAttributesRepository {
	return &SpProdAttributesRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建商品属性
func (r *SpProdAttributesRepository) Create(attr *sp.SpProdAttributes) error {
	return r.db.Create(attr).Error
}

// 更新商品属性
func (r *SpProdAttributesRepository) Update(attr *sp.SpProdAttributes) error {
	return r.db.Save(attr).Error
}

// 根据ID获取属性
func (r *SpProdAttributesRepository) FindByID(id uint) (*sp.SpProdAttributes, error) {
	var attr sp.SpProdAttributes
	err := r.db.First(&attr, id).Error
	return &attr, err
}

// 获取所有属性
func (r *SpProdAttributesRepository) FindAll() ([]sp.SpProdAttributes, error) {
	var attrs []sp.SpProdAttributes
	err := r.db.Order("sort_num ASC").Find(&attrs).Error
	return attrs, err
}

// 更新属性排序
func (r *SpProdAttributesRepository) UpdateSortNum(id uint, sortNum uint16) error {
	return r.db.Model(&sp.SpProdAttributes{}).
		Where("id = ?", id).
		Update("sort_num", sortNum).Error
}

// 删除属性
func (r *SpProdAttributesRepository) Delete(id uint) error {
	return r.db.Delete(&sp.SpProdAttributes{}, id).Error
}