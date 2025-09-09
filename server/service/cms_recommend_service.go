package service

import (
	"errors"
	"server/models/cms"
)

type CmsRecommendService struct {
	*Service
}

func NewCmsRecommendService(base *Service) *CmsRecommendService {
	return &CmsRecommendService{Service: base}
}

// CreateRecommend 创建推荐内容
func (s *CmsRecommendService) CreateRecommend(recommend *cms.CmsRecommend) error {
	if recommend.Title == "" {
		return errors.New("推荐标题不能为空")
	}
	if recommend.Description == "" {
		return errors.New("推荐位置不能为空")
	}
	
	// 设置默认状态
	if recommend.State == 0 {
		recommend.State = 1 // 默认启用
	}
	
	return s.repoFactory.GetCmsRecommendRepository().Create(recommend)
}

// UpdateRecommend 更新推荐内容
func (s *CmsRecommendService) UpdateRecommend(recommend *cms.CmsRecommend) error {
	if recommend.ID <= 0 {
		return errors.New("无效的推荐ID")
	}
	if recommend.Title == "" {
		return errors.New("推荐标题不能为空")
	}
	
	// 检查推荐是否存在
	_, err := s.repoFactory.GetCmsRecommendRepository().FindByID(recommend.ID)
	if err != nil {
		return errors.New("推荐内容不存在")
	}
	
	return s.repoFactory.GetCmsRecommendRepository().Update(recommend)
}

// GetActiveRecommends 获取所有有效推荐
func (s *CmsRecommendService) GetActiveRecommends() ([]cms.CmsRecommend, error) {
	return s.repoFactory.GetCmsRecommendRepository().FindAll()
}

// GetRecommendsByState 根据状态获取推荐
func (s *CmsRecommendService) GetRecommendsByState(state int8) ([]cms.CmsRecommend, error) {
	if state < 0 || state > 2 {
		return nil, errors.New("无效的状态值")
	}
	return s.repoFactory.GetCmsRecommendRepository().FindByState(state)
}

func (s *CmsRecommendService) ListRecommends(prarams cms.RecommendQueryParams) ([]cms.CmsRecommend, int64, error) {
	return s.repoFactory.GetCmsRecommendRepository().ListWithPagination(prarams)
}

func (s *CmsRecommendService) DeleteRecommendByID(id int) error {
	if id <= 0 {
		return errors.New("无效的推荐ID")
	}
	return s.repoFactory.GetCmsRecommendRepository().Delete(id)
}

func (s *CmsRecommendService) GetRecommendByID(id int) (*cms.CmsRecommend, error) {
	if id <= 0 {
		return nil, errors.New("无效的推荐ID")
	}
	return s.repoFactory.GetCmsRecommendRepository().FindByID(id)
}

func (s *CmsRecommendService) GetRecommendByCode(recommendCode string) (*cms.CmsRecommend, error) {
	if recommendCode == "" {
		return nil, errors.New("无效的推荐ID")
	}
	return s.repoFactory.GetCmsRecommendRepository().FindByDocumentCode(recommendCode)
}