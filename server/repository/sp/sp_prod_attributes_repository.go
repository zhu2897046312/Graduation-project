package sp

import (
	"gorm.io/gorm"
	"time"
	"server/models/sp"
	"server/repository/base"
)

type SpProdAttributesRepository struct {
	*base.BaseRepository
}

func NewSpProdAttributesRepository(DB *gorm.DB) *SpProdAttributesRepository {
	return &SpProdAttributesRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建商品属性
func (r *SpProdAttributesRepository) Create(attr *sp.SpProdAttributes) error {
	return r.DB.Create(attr).Error
}

// 更新商品属性
func (r *SpProdAttributesRepository) Update(attr *sp.SpProdAttributes) error {
	return r.DB.Updates(attr).Error
}

// 根据ID获取属性
func (r *SpProdAttributesRepository) FindByID(id uint) (*sp.SpProdAttributes, error) {
	var attr sp.SpProdAttributes
	err := r.DB.First(&attr, id).Error
	return &attr, err
}

// 获取所有属性
func (r *SpProdAttributesRepository) FindAll() ([]sp.SpProdAttributes, error) {
	var attrs []sp.SpProdAttributes
	err := r.DB.Order("sort_num ASC").Find(&attrs).Error
	return attrs, err
}

// 分页查询属性
func (r *SpProdAttributesRepository) FindByPage(title string, page, pageSize int) ([]sp.SpProdAttributes, int64, error) {
	var attrs []sp.SpProdAttributes
	var total int64
	
	query := r.DB.Model(&sp.SpProdAttributes{}).Where("deleted_time IS NULL")
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	
	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Order("sort_num ASC, id DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&attrs).Error
		
	return attrs, total, err
}

// 更新属性排序
func (r *SpProdAttributesRepository) UpdateSortNum(id uint, sortNum uint16) error {
	return r.DB.Model(&sp.SpProdAttributes{}).
		Where("id = ?", id).
		Update("sort_num", sortNum).Error
}

// 删除属性
func (r *SpProdAttributesRepository) Delete(id uint) error {
	result := r.DB.Model(&sp.SpProdAttributes{}).
		Where("id = ?", id).
		Update("deleted_time", time.Now())

	return result.Error
}