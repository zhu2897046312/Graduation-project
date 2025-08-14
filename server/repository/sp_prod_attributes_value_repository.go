package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
)

type SpProdAttributesValueRepository struct {
	*BaseRepository
}

func NewSpProdAttributesValueRepository(db *gorm.DB) *SpProdAttributesValueRepository {
	return &SpProdAttributesValueRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建属性值
func (r *SpProdAttributesValueRepository) Create(value *sp.SpProdAttributesValue) error {
	return r.db.Create(value).Error
}

// 更新属性值
func (r *SpProdAttributesValueRepository) Update(value *sp.SpProdAttributesValue) error {
	return r.db.Save(value).Error
}

// 根据属性ID获取值列表
func (r *SpProdAttributesValueRepository) FindByAttributeID(attrID uint) ([]sp.SpProdAttributesValue, error) {
	var values []sp.SpProdAttributesValue
	err := r.db.Where("prod_attributes_id = ?", attrID).
		Order("sort_num ASC").
		Find(&values).Error
	return values, err
}

// 根据ID获取属性值
func (r *SpProdAttributesValueRepository) FindByID(id uint) (*sp.SpProdAttributesValue, error) {
	var value sp.SpProdAttributesValue
	err := r.db.First(&value, id).Error
	return &value, err
}

// 批量创建属性值
func (r *SpProdAttributesValueRepository) BatchCreate(values []sp.SpProdAttributesValue) error {
	return r.db.Create(&values).Error
}

// 删除属性的所有值
func (r *SpProdAttributesValueRepository) DeleteByAttributeID(attrID uint) error {
	return r.db.Where("prod_attributes_id = ?", attrID).
		Delete(&sp.SpProdAttributesValue{}).Error
}