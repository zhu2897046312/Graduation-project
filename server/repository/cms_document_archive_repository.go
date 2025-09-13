package repository

import (
	"gorm.io/gorm"
	"time"
	"server/models/cms"
	"server/models/common"
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
func (r *CmsDocumentArchiveRepository) FindByDocumentID(documentID common.MyID) (*cms.CmsDocumentArchive, error) {
	var archive cms.CmsDocumentArchive
	err := r.db.Where("document_id = ?", documentID).First(&archive).Error
	return &archive, err
}
func (r *CmsDocumentArchiveRepository) DeleteByDocumentID(documentID common.MyID) error {
    result := r.db.Model(&cms.CmsDocumentArchive{}).
		Where("document_id = ?", documentID).
		Update("deleted_time", time.Now())

	return result.Error
}

func (r *CmsDocumentArchiveRepository) FindByDocumentCode(documentCode common.MyID) (*cms.CmsDocumentArchive, error) {
	var archive cms.CmsDocumentArchive
	err := r.db.Where("document_id = ?", documentCode).First(&archive).Error
	return &archive, err
}