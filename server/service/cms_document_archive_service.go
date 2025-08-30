package service

import (
	"errors"
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsDocumentArchiveService struct {
	*Service
}

func NewCmsDocumentArchiveService(base *Service) *CmsDocumentArchiveService {
	return &CmsDocumentArchiveService{Service: base}
}

// CreateArchive 创建文档存档
func (s *CmsDocumentArchiveService) CreateArchive(archive *cms.CmsDocumentArchive) error {
	if archive.DocumentID <= 0 {
		return errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsDocumentArchiveRepository().Create(archive)
}

// UpdateArchive 更新文档存档
func (s *CmsDocumentArchiveService) UpdateArchive(archive *cms.CmsDocumentArchive) error {
	if archive.DocumentID <= 0 {
		return errors.New("无效的文档ID")
	}
	
	// 检查存档是否存在
	_, err := s.repoFactory.GetCmsDocumentArchiveRepository().FindByDocumentID(archive.DocumentID)
	if err != nil {
		return errors.New("存档不存在")
	}
	
	// 保留原始创建时间（如果模型中有该字段）
	// archive.CreatedAt = existing.CreatedAt
	
	return s.repoFactory.GetCmsDocumentArchiveRepository().Update(archive)
}

// GetArchiveByDocumentID 根据文档ID获取存档
func (s *CmsDocumentArchiveService) GetArchiveByDocumentID(documentID int64) (*cms.CmsDocumentArchive, error) {
	if documentID <= 0 {
		return nil, errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsDocumentArchiveRepository().FindByDocumentID(documentID)
}

func (s *CmsDocumentArchiveService) DeleteByDocumetnID(id int64) error {
	_,err := s.repoFactory.GetCmsDocumentArchiveRepository().FindByDocumentID(id)
	if err == gorm.ErrRecordNotFound {
		return errors.New("存档不存在")
	}
    return s.repoFactory.GetCmsDocumentArchiveRepository().DeleteByDocumentID(id)
}