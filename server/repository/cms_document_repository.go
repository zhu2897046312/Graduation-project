package repository

import (
	"fmt"
	"server/models/cms"

	"gorm.io/gorm"
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
// 分页获取文档（支持 title 模糊查询）
func (r *CmsDocumentRepository) ListWithPagination(page, pageSize int, title string) ([]cms.CmsDocument, int64, error) {
    var documents []cms.CmsDocument
    var total int64

    // 初始化查询构建器
    query := r.db.Model(&cms.CmsDocument{})

    // 如果 title 不为空，添加模糊查询条件
    if title != "" {
        query = query.Where("title LIKE ?", "%"+title+"%")
    }

    // 获取总数
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 如果 pageSize <= 0，返回全部数据
    if pageSize <= 0 {
        err := query.Order("created_time DESC").Find(&documents).Error
        return documents, total, err
    }

    // 否则返回分页数据
    offset := (page - 1) * pageSize
    err := query.Offset(offset).
        Limit(pageSize).
        Order("created_time DESC").
        Find(&documents).Error

    return documents, total, err
}


// cms_document_repository.go
// 在文件末尾添加以下方法

// 创建文档
func (r *CmsDocumentRepository) Create(document *cms.CmsDocument) error {
    fmt.Println(document.ID)
    err :=r.db.Create(document).Error
    fmt.Println(document.ID)
    return err
}

// 更新文档
func (r *CmsDocumentRepository) Update(document *cms.CmsDocument) error {
    return r.db.Updates(document).Error
}

// 根据ID查找文档
func (r *CmsDocumentRepository) FindByID(id int64) (*cms.CmsDocument, error) {
    var document cms.CmsDocument
    err := r.db.First(&document, id).Error
    if err != nil {
        return nil, err
    }
    return &document, nil
}

func (r *CmsDocumentRepository) DeleteByID(id int64) error {
    return r.db.Delete(&cms.CmsDocument{}, id).Error
}