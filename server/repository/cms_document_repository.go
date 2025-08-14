package repository

import (
    "gorm.io/gorm"
    "server/models/cms"
)

type CmsDocumentRepository struct {
    *BaseRepository
}

func NewCmsDocumentRepository(db *gorm.DB) *CmsDocumentRepository {
    return &CmsDocumentRepository{
        BaseRepository: NewBaseRepository(db),
    }
}

// 根据分类ID获取文档
func (r *CmsDocumentRepository) FindByCategoryID(categoryID int64) ([]cms.CmsDocument, error) {
    var documents []cms.CmsDocument
    err := r.db.Where("category_id = ?", categoryID).Find(&documents).Error
    return documents, err
}

// 获取热门文档
func (r *CmsDocumentRepository) FindPopular(limit int) ([]cms.CmsDocument, error) {
    var documents []cms.CmsDocument
    err := r.db.Order("read_num DESC").Limit(limit).Find(&documents).Error
    return documents, err
}

// 分页获取文档
func (r *CmsDocumentRepository) ListWithPagination(page, pageSize int) ([]cms.CmsDocument, int64, error) {
    var documents []cms.CmsDocument
    var total int64
    
    offset := (page - 1) * pageSize

    // 获取总数
    if err := r.db.Model(&cms.CmsDocument{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 获取分页数据
    err := r.db.Offset(offset).
        Limit(pageSize).
        Order("created_time DESC").
        Find(&documents).Error

    return documents, total, err
}