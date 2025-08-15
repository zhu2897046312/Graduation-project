package service

import (
	"errors"
	"server/models/cms"
)

type CmsDocumentService struct {
	*Service
}

func NewCmsDocumentService(base *Service) *CmsDocumentService {
	return &CmsDocumentService{Service: base}
}

// GetDocumentsByCategoryID 根据分类ID获取文档
func (s *CmsDocumentService) GetDocumentsByCategoryID(categoryID int64) ([]cms.CmsDocument, error) {
	if categoryID <= 0 {
		return nil, errors.New("无效的分类ID")
	}
	return s.repoFactory.GetCmsDocumentRepository().FindByCategoryID(categoryID)
}

// GetPopularDocuments 获取热门文档
func (s *CmsDocumentService) GetPopularDocuments(limit int) ([]cms.CmsDocument, error) {
	if limit <= 0 || limit > 100 {
		limit = 10 // 默认限制
	}
	return s.repoFactory.GetCmsDocumentRepository().FindPopular(limit)
}

// ListDocuments 分页获取文档
func (s *CmsDocumentService) ListDocuments(page, pageSize int) ([]cms.CmsDocument, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	
	return s.repoFactory.GetCmsDocumentRepository().ListWithPagination(page, pageSize)
}