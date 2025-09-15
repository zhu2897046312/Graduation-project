package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsDocumentTagRepository struct {
	*BaseRepository
}

func NewCmsDocumentTagRepository(db *gorm.DB) *CmsDocumentTagRepository {
	return &CmsDocumentTagRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建文档标签关联
func (r *CmsDocumentTagRepository) Create(docTag *cms.CmsDocumentTag) error {
	return r.db.Create(docTag).Error
}

// 删除文档标签关联
func (r *CmsDocumentTagRepository) Delete(documentID, tagID int64) error {
	return r.db.Where("document_id = ? AND tag_id = ?", documentID, tagID).
		Delete(&cms.CmsDocumentTag{}).Error
}

// 根据文档ID获取标签关联
func (r *CmsDocumentTagRepository) FindByDocumentID(documentID int64) ([]cms.CmsDocumentTag, error) {
	var tags []cms.CmsDocumentTag
	err := r.db.Where("document_id = ?", documentID).Find(&tags).Error
	return tags, err
}

// 根据标签ID获取文档关联
func (r *CmsDocumentTagRepository) FindByTagID(tagID int64) ([]cms.CmsDocumentTag, error) {
	var tags []cms.CmsDocumentTag
	err := r.db.Where("tag_id = ?", tagID).Find(&tags).Error
	return tags, err
}