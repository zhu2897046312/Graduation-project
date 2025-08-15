package service

import (
	"errors"
	"server/models/cms"
)

type CmsDocumentVideoService struct {
	*Service
}

func NewCmsDocumentVideoService(base *Service) *CmsDocumentVideoService {
	return &CmsDocumentVideoService{Service: base}
}

// CreateDocumentVideo 创建文档视频
func (s *CmsDocumentVideoService) CreateDocumentVideo(video *cms.CmsDocumentVideo) error {
	if video.DocumentID <= 0 {
		return errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsDocumentVideoRepository().Create(video)
}

// UpdateDocumentVideo 更新文档视频
func (s *CmsDocumentVideoService) UpdateDocumentVideo(video *cms.CmsDocumentVideo) error {
	if video.DocumentID <= 0 {
		return errors.New("无效的文档ID")
	}
	
	// 检查视频是否存在
	_, err := s.repoFactory.GetCmsDocumentVideoRepository().FindByDocumentID(video.DocumentID)
	if err != nil {
		return errors.New("视频不存在")
	}
	
	// 保留原始创建时间（如果模型中有该字段）
	// video.CreatedAt = existing.CreatedAt
	
	return s.repoFactory.GetCmsDocumentVideoRepository().Update(video)
}

// GetVideoByDocumentID 根据文档ID获取视频
func (s *CmsDocumentVideoService) GetVideoByDocumentID(documentID int64) (*cms.CmsDocumentVideo, error) {
	if documentID <= 0 {
		return nil, errors.New("无效的文档ID")
	}
	return s.repoFactory.GetCmsDocumentVideoRepository().FindByDocumentID(documentID)
}