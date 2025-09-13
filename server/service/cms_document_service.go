package service

import (
	"errors"
	"server/models/cms"
    "server/models/common"
)

type CmsDocumentService struct {
	*Service
}

func NewCmsDocumentService(base *Service) *CmsDocumentService {
	return &CmsDocumentService{Service: base}
}

// GetDocumentsByCategoryID 根据分类ID获取文档
func (s *CmsDocumentService) GetDocumentsByCategoryID(categoryID common.MyID) ([]cms.CmsDocument, error) {
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
func (s *CmsDocumentService) ListDocuments(page, pageSize int,title string) ([]cms.CmsDocument, int64, error) {	
	return s.repoFactory.GetCmsDocumentRepository().ListWithPagination(page, pageSize,title)
}// cms_document_service.go
// 在文件末尾添加以下方法

// CreateDocument 创建文档
func (s *CmsDocumentService) CreateDocument(document *cms.CmsDocument) error {
    if document == nil {
        return errors.New("文档不能为空")
    }
    return s.repoFactory.GetCmsDocumentRepository().Create(document)
}

// UpdateDocument 更新文档
func (s *CmsDocumentService) UpdateDocument(document *cms.CmsDocument) error {
    if document == nil || document.ID <= 0 {
        return errors.New("无效的文档ID")
    }
    return s.repoFactory.GetCmsDocumentRepository().Update(document)
}

// GetDocumentByID 根据ID获取文档
func (s *CmsDocumentService) GetDocumentByID(id common.MyID) (*cms.CmsDocument, error) {
    if id <= 0 {
        return nil, errors.New("无效的文档ID")
    }
    return s.repoFactory.GetCmsDocumentRepository().FindByID(id)
}

func (s *CmsDocumentService) GetDocumentByCode(code string) (*cms.CmsDocument, error) {
    if code == "" {
        return nil, errors.New("无效的文档ID")
    }
    return s.repoFactory.GetCmsDocumentRepository().FindByDocumentCode(code)
}

func (s *CmsDocumentService) DeleteByID(id common.MyID) error {
    if id <= 0 {
        return errors.New("无效的文档ID")
    }
    return s.repoFactory.GetCmsDocumentRepository().DeleteByID(id)
}