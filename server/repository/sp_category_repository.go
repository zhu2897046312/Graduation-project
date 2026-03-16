package repository

import (
	"time"

	"gorm.io/gorm"
	"server/models/common"
	"server/models/sp"
)

type SpCategoryRepository struct {
	*BaseRepository
}

func NewSpCategoryRepository(db *gorm.DB) *SpCategoryRepository {
	return &SpCategoryRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建分类
func (r *SpCategoryRepository) Create(category *sp.SpCategory) error {
	return r.db.Create(category).Error
}

// 更新分类
func (r *SpCategoryRepository) Update(category *sp.SpCategory) error {
	return r.db.Updates(category).Error
}

// 根据ID获取分类
func (r *SpCategoryRepository) FindByID(id common.MyID) (*sp.SpCategory, error) {
	var category sp.SpCategory
	err := r.db.First(&category, id).Error
	return &category, err
}

// 获取所有分类
func (r *SpCategoryRepository) FindAll() ([]*sp.SpCategory, error) {
	var categories []*sp.SpCategory
	err := r.db.
		Where("deleted_time IS NULL").
		Order("sort_num ASC").
		Find(&categories).Error
	return categories, err
}

// 根据父ID获取子分类
func (r *SpCategoryRepository) FindByPid(pid common.MyID) ([]*sp.SpCategory, error) {
	var categories []*sp.SpCategory
	err := r.db.Where("pid = ?", pid).
		Order("sort_num ASC").
		Find(&categories).Error
	return categories, err
}

// 更新分类状态
func (r *SpCategoryRepository) UpdateState(id common.MyID, state uint8) error {
	return r.db.Model(&sp.SpCategory{}).
		Where("id = ?", id).
		Update("state", state).Error
}

// 更新分类排序
func (r *SpCategoryRepository) UpdateSortNum(id common.MyID, sortNum uint16) error {
	return r.db.Model(&sp.SpCategory{}).
		Where("id = ?", id).
		Update("sort_num", sortNum).Error
}

func (r *SpCategoryRepository) FindByCode(code string) (*sp.SpCategory, error) {
	var category sp.SpCategory
	err := r.db.Where("code = ?", code).First(&category).Error
	return &category, err
}

// SoftDelete 软删除分类（设置 deleted_time）
func (r *SpCategoryRepository) SoftDelete(id common.MyID) error {
	return r.db.Model(&sp.SpCategory{}).
		Where("id = ?", id).
		Update("deleted_time", time.Now()).Error
}

// SoftDeleteByIDs 批量软删除分类
func (r *SpCategoryRepository) SoftDeleteByIDs(ids []common.MyID) error {
	if len(ids) == 0 {
		return nil
	}
	return r.db.Model(&sp.SpCategory{}).
		Where("id IN ?", ids).
		Update("deleted_time", time.Now()).Error
}