package service

import (
	"errors"
	"server/models/cms"
)

type CmsRecommendIndexService struct {
	*Service
}

func NewCmsRecommendIndexService(base *Service) *CmsRecommendIndexService {
	return &CmsRecommendIndexService{Service: base}
}

// CreateRecommendIndex 创建推荐索引
func (s *CmsRecommendIndexService) CreateRecommendIndex(index *cms.CmsRecommendIndex) error {
	if index.RecommendID <= 0 {
		return errors.New("无效的推荐ID")
	}
	if index.RecommendID <= 0 {
		return errors.New("无效的内容ID")
	}
	
	return s.repoFactory.GetCmsRecommendIndexRepository().Create(index)
}

// UpdateRecommendIndex 更新推荐索引
func (s *CmsRecommendIndexService) UpdateRecommendIndex(index *cms.CmsRecommendIndex) error {
	if index.ID <= 0 {
		return errors.New("无效的索引ID")
	}
	
	// 检查索引是否存在
	existing, err := s.repoFactory.GetCmsRecommendIndexRepository().FindByRecommendID(index.RecommendID)
	if err != nil || len(existing) == 0 {
		return errors.New("推荐索引不存在")
	}
	
	return s.repoFactory.GetCmsRecommendIndexRepository().Update(index)
}

// GetIndicesByRecommendID 根据推荐ID获取索引
func (s *CmsRecommendIndexService) GetIndicesByRecommendID(recommendID int) ([]cms.CmsRecommendIndex, error) {
	if recommendID <= 0 {
		return nil, errors.New("无效的推荐ID")
	}
	return s.repoFactory.GetCmsRecommendIndexRepository().FindByRecommendID(recommendID)
}

// GetIndicesByState 根据状态获取推荐索引
func (s *CmsRecommendIndexService) GetIndicesByState(state int8) ([]cms.CmsRecommendIndex, error) {
	if state < 0 || state > 2 {
		return nil, errors.New("无效的状态值")
	}
	return s.repoFactory.GetCmsRecommendIndexRepository().FindByState(state)
}

// DeleteRecommendIndex 删除推荐索引
func (s *CmsRecommendIndexService) DeleteRecommendIndex(id int) error {
	if id <= 0 {
		return errors.New("无效的索引ID")
	}
	return s.repoFactory.GetCmsRecommendIndexRepository().Delete(id)
}