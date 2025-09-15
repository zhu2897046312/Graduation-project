package service

import (
	"errors"
	"server/models/cms"
)

type CmsDocumentTagService struct {
	*Service
}

func NewCmsDocumentTagService(base *Service) *CmsDocumentTagService {
	return &CmsDocumentTagService{Service: base}
}

// CreateDocumentTag 创建文档标签关联
func (s *CmsDocumentTagService) CreateDocumentTag(docTag *cms.CmsDocumentTag) error {
	if docTag.DocumentID <= 0 {
		return errors.New("无效的文档ID")
	}
	if docTag.TagID <= 0 {
		return errors.New("无效的标签ID")
	}
	return s.repoFactory.GetCmsDocumentTagRepository().Create(docTag)
}

// DeleteDocumentTag 删除文档标签关联
func (s *CmsDocumentTagService) DeleteDocumentTag(documentID, tagID int64) error {
	if documentID <= 0 {
		return errors.New("无效的文档ID")
	}
	if tagID <= 0 {
		return errors.New("无效的标签ID")
	}
	return s.repoFactory.GetCmsDocumentTagRepository().Delete(documentID, tagID)
}

// GetTagsByDocumentID 根据文档ID获取标签
func (s *CmsDocumentTagService) GetTagsByDocumentID(documentID int64) ([]cms.CmsDocumentTag, error) {
	if documentID <= 0 {
		return nil, errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsDocumentTagRepository().FindByDocumentID(documentID)
}

// GetDocumentsByTagID 根据标签ID获取文档
func (s *CmsDocumentTagService) GetDocumentsByTagID(tagID int64) ([]cms.CmsDocumentTag, error) {
	if tagID <= 0 {
		return nil, errors.New("无效的标签ID")
	}
	return s.repoFactory.GetCmsDocumentTagRepository().FindByTagID(tagID)
}