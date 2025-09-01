package cms

import (
	"gorm.io/gorm"
	"time"
	"server/models/cms"
	"server/repository/base"
)

type CmsDocumentArchiveRepository struct {
	*base.BaseRepository
}

func NewCmsDocumentArchiveRepository(DB *gorm.DB) *CmsDocumentArchiveRepository {
	return &CmsDocumentArchiveRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建文档存档
func (r *CmsDocumentArchiveRepository) Create(archive *cms.CmsDocumentArchive) error {
	return r.DB.Create(archive).Error
}

// 更新文档存档
func (r *CmsDocumentArchiveRepository) Update(archive *cms.CmsDocumentArchive) error {
	return r.DB.Save(archive).Error
}

// 根据文档ID获取存档
func (r *CmsDocumentArchiveRepository) FindByDocumentID(documentID int64) (*cms.CmsDocumentArchive, error) {
	var archive cms.CmsDocumentArchive
	err := r.DB.Where("document_id = ?", documentID).First(&archive).Error
	return &archive, err
}
func (r *CmsDocumentArchiveRepository) DeleteByDocumentID(documentID int64) error {
    result := r.DB.Model(&cms.CmsDocumentArchive{}).
		Where("document_id = ?", documentID).
		Update("deleted_time", time.Now())

	return result.Error
}