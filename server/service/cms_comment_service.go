package service

import (
	"errors"
	"server/models/cms"
)

type CmsCommentService struct {
	*Service
}

func NewCmsCommentService(base *Service) *CmsCommentService {
	return &CmsCommentService{Service: base}
}

// GetCommentsByDocumentID 根据文档ID获取评论
func (s *CmsCommentService) GetCommentsByDocumentID(documentID int64) ([]cms.CmsComment, error) {
	if documentID <= 0 {
		return nil, errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsCommentRepository().FindByDocumentID(documentID)
}

// GetCommentsByUserID 根据用户ID获取评论
func (s *CmsCommentService) GetCommentsByUserID(userID int64) ([]cms.CmsComment, error) {
	if userID <= 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetCmsCommentRepository().FindByUserID(userID)
}

// GetTopLevelComments 获取顶级评论
func (s *CmsCommentService) GetTopLevelComments() ([]cms.CmsComment, error) {
	return s.repoFactory.GetCmsCommentRepository().FindTopLevelComments()
}

// GetCommentReplies 获取评论回复
func (s *CmsCommentService) GetCommentReplies(commentID int64) ([]cms.CmsComment, error) {
	if commentID <= 0 {
		return nil, errors.New("无效的评论ID")
	}
	return s.repoFactory.GetCmsCommentRepository().FindReplies(commentID)
}

// ListComments 分页获取评论
func (s *CmsCommentService) ListComments(page, pageSize int) ([]cms.CmsComment, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	
	return s.repoFactory.GetCmsCommentRepository().ListWithPagination(page, pageSize)
}