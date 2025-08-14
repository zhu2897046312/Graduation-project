package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsDocumentVideoRepository struct {
	*BaseRepository
}

func NewCmsDocumentVideoRepository(db *gorm.DB) *CmsDocumentVideoRepository {
	return &CmsDocumentVideoRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建文档视频
func (r *CmsDocumentVideoRepository) Create(video *cms.CmsDocumentVideo) error {
	return r.db.Create(video).Error
}

// 更新文档视频
func (r *CmsDocumentVideoRepository) Update(video *cms.CmsDocumentVideo) error {
	return r.db.Save(video).Error
}

// 根据文档ID获取视频
func (r *CmsDocumentVideoRepository) FindByDocumentID(documentID int64) (*cms.CmsDocumentVideo, error) {
	var video cms.CmsDocumentVideo
	err := r.db.Where("document_id = ?", documentID).First(&video).Error
	return &video, err
}