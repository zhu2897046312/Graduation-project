package cms

import (
	"gorm.io/gorm"
	"server/models/cms"
	"server/repository/base"
)

type CmsDocumentVideoRepository struct {
	*base.BaseRepository
}

func NewCmsDocumentVideoRepository(DB *gorm.DB) *CmsDocumentVideoRepository {
	return &CmsDocumentVideoRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建文档视频
func (r *CmsDocumentVideoRepository) Create(video *cms.CmsDocumentVideo) error {
	return r.DB.Create(video).Error
}

// 更新文档视频
func (r *CmsDocumentVideoRepository) Update(video *cms.CmsDocumentVideo) error {
	return r.DB.Save(video).Error
}

// 根据文档ID获取视频
func (r *CmsDocumentVideoRepository) FindByDocumentID(documentID int64) (*cms.CmsDocumentVideo, error) {
	var video cms.CmsDocumentVideo
	err := r.DB.Where("document_id = ?", documentID).First(&video).Error
	return &video, err
}