package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpCategoryRepository struct {
	*base.BaseRepository
}

func NewSpCategoryRepository(DB *gorm.DB) *SpCategoryRepository {
	return &SpCategoryRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建分类
func (r *SpCategoryRepository) Create(category *sp.SpCategory) error {
	return r.DB.Create(category).Error
}

// 更新分类
func (r *SpCategoryRepository) Update(category *sp.SpCategory) error {
	return r.DB.Updates(category).Error
}

// 根据ID获取分类
func (r *SpCategoryRepository) FindByID(id uint) (*sp.SpCategory, error) {
	var category sp.SpCategory
	err := r.DB.First(&category, id).Error
	return &category, err
}

// 获取所有分类
func (r *SpCategoryRepository) FindAll() ([]*sp.SpCategory, error) {
	var categories []*sp.SpCategory
	err := r.DB.Order("sort_num ASC").Find(&categories).Error
	return categories, err
}

// 根据父ID获取子分类
func (r *SpCategoryRepository) FindByPid(pid uint) ([]*sp.SpCategory, error) {
	var categories []*sp.SpCategory
	err := r.DB.Where("pid = ?", pid).
		Order("sort_num ASC").
		Find(&categories).Error
	return categories, err
}

// 更新分类状态
func (r *SpCategoryRepository) UpdateState(id uint, state uint8) error {
	return r.DB.Model(&sp.SpCategory{}).
		Where("id = ?", id).
		Update("state", state).Error
}

// 更新分类排序
func (r *SpCategoryRepository) UpdateSortNum(id uint, sortNum uint16) error {
	return r.DB.Model(&sp.SpCategory{}).
		Where("id = ?", id).
		Update("sort_num", sortNum).Error
}