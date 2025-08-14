package repository

import (
    "gorm.io/gorm"
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

// 根据父ID获取分类
func (r *CmsCategoryRepository) FindByParentID(parentID int64) ([]cms.CmsCategory, error) {
    var categories []cms.CmsCategory
    err := r.db.Where("pid = ?", parentID).Find(&categories).Error
    return categories, err
}

// 根据类型获取分类
func (r *CmsCategoryRepository) FindByType(categoryType int8) ([]cms.CmsCategory, error) {
    var categories []cms.CmsCategory
    err := r.db.Where("category_type = ?", categoryType).Find(&categories).Error
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