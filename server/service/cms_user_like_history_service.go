package service

import (
	"errors"
	"server/models/cms"
)

type CmsUserLikeHistoryService struct {
	*Service
}

func NewCmsUserLikeHistoryService(base *Service) *CmsUserLikeHistoryService {
	return &CmsUserLikeHistoryService{Service: base}
}

// CreateLikeHistory 创建点赞记录
func (s *CmsUserLikeHistoryService) CreateLikeHistory(history *cms.CmsUserLikeHistory) error {
	if history.UserID <= 0 {
		return errors.New("无效的用户ID")
	}
	if history.DocumentID <= 0 {
		return errors.New("无效的文档ID")
	}
	
	return s.repoFactory.GetCmsUserLikeHistoryRepository().Create(history)
}

// UpdateLikeHistory 更新点赞记录
func (s *CmsUserLikeHistoryService) UpdateLikeHistory(history *cms.CmsUserLikeHistory) error {
	if history.ID <= 0 {
		return errors.New("无效的记录ID")
	}
	
	return s.repoFactory.GetCmsUserLikeHistoryRepository().Update(history)
}

// GetLikeHistoryByUserID 根据用户ID获取点赞记录
func (s *CmsUserLikeHistoryService) GetLikeHistoryByUserID(userID int64) ([]cms.CmsUserLikeHistory, error) {
	if userID <= 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetCmsUserLikeHistoryRepository().FindByUserID(userID)
}

// GetLikeHistoryByDocumentID 根据文档ID获取点赞记录
func (s *CmsUserLikeHistoryService) GetLikeHistoryByDocumentID(documentID int64) ([]cms.CmsUserLikeHistory, error) {
	if documentID <= 0 {
		return nil, errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsUserLikeHistoryRepository().FindByDocumentID(documentID)
}

// CheckUserLikedDocument 检查用户是否点赞过文档
func (s *CmsUserLikeHistoryService) CheckUserLikedDocument(userID, documentID int64) (bool, error) {
	if userID <= 0 {
		return false, errors.New("无效的用户ID")
	}
	if documentID <= 0 {
		return false, errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsUserLikeHistoryRepository().HasLiked(userID, documentID)
}

// GetLikeCountByDocumentID 获取文档点赞数
func (s *CmsUserLikeHistoryService) GetLikeCountByDocumentID(documentID int64) (int64, error) {
	if documentID <= 0 {
		return 0, errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsUserLikeHistoryRepository().CountByDocumentID(documentID)
}