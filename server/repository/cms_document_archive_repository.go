package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsDocumentArchiveRepository struct {
	*BaseRepository
}

func NewCmsDocumentArchiveRepository(db *gorm.DB) *CmsDocumentArchiveRepository {
	return &CmsDocumentArchiveRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建文档存档
func (r *CmsDocumentArchiveRepository) Create(archive *cms.CmsDocumentArchive) error {
	return r.db.Create(archive).Error
}

// 更新文档存档
func (r *CmsDocumentArchiveRepository) Update(archive *cms.CmsDocumentArchive) error {
	return r.db.Save(archive).Error
}

// 根据文档ID获取存档
func (r *CmsDocumentArchiveRepository) FindByDocumentID(documentID int64) (*cms.CmsDocumentArchive, error) {
	var archive cms.CmsDocumentArchive
	err := r.db.Where("document_id = ?", documentID).First(&archive).Error
	return &archive, err
}