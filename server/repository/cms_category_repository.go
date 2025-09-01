package repository

import (
	"gorm.io/gorm"
	"time"
	"server/models/cms"
)

type CmsCategoryRepository struct {
	*BaseRepository
}

func NewCmsCategoryRepository(db *gorm.DB) *CmsCategoryRepository {
	return &CmsCategoryRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建分类
func (r *CmsCategoryRepository) Create(category *cms.CmsCategory) error {
	return r.db.Create(category).Error
}

// 更新分类
func (r *CmsCategoryRepository) Update(category *cms.CmsCategory) error {
	return r.db.Save(category).Error
}

// 根据ID获取分类
func (r *CmsCategoryRepository) FindByID(id int64) (*cms.CmsCategory, error) {
	var category cms.CmsCategory
	err := r.db.First(&category, id).Error
	return &category, err
}

// 根据父ID获取分类
func (r *CmsCategoryRepository) FindByParentID(parentID int64) ([]cms.CmsCategory, error) {
	var categories []cms.CmsCategory
	err := r.db.Where("pid = ?", parentID).
		Order("sort_num ASC").
		Find(&categories).Error
	return categories, err
}

// 根据类型获取分类
func (r *CmsCategoryRepository) FindByType(categoryType int8) ([]cms.CmsCategory, error) {
	var categories []cms.CmsCategory
	err := r.db.Where("category_type = ?", categoryType).
		Order("sort_num ASC").
		Find(&categories).Error
	return categories, err
}

// 获取所有分类
func (r *CmsCategoryRepository) FindAll() ([]cms.CmsCategory, error) {
	var categories []cms.CmsCategory
	err := r.db.Order("sort_num ASC").Find(&categories).Error
	return categories, err
}

// 分页获取分类
func (r *CmsCategoryRepository) ListWithPagination(page, pageSize int) ([]cms.CmsCategory, int64, error) {
	var categories []cms.CmsCategory
	var total int64
	
	offset := (page - 1) * pageSize

	// 获取总数
	if err := r.db.Model(&cms.CmsCategory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := r.db.Offset(offset).
		Limit(pageSize).
		Order("sort_num ASC").
		Find(&categories).Error

	return categories, total, err
}

// 更新分类排序
func (r *CmsCategoryRepository) UpdateSortNum(id int64, sortNum int) error {
	return r.db.Model(&cms.CmsCategory{}).
		Where("id = ?", id).
		Update("sort_num", sortNum).Error
}

// 删除分类
func (r *CmsCategoryRepository) Delete(id int64) error {
	result := r.db.Model(&cms.CmsCategory{}).
		Where("id = ?", id).
		Update("deleted_time", time.Now())

	return result.Error
}