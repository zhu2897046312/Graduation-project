package service

import (
	"errors"
	"server/models/cms"
)

type CmsTagService struct {
	*Service
}

func NewCmsTagService(base *Service) *CmsTagService {
	return &CmsTagService{Service: base}
}

// CreateTag 创建标签
func (s *CmsTagService) CreateTag(tag *cms.CmsTag) error {
	if tag.Title == "" {
		return errors.New("标签名称不能为空")
	}
	
	// 设置默认排序值
	if tag.SortNum == 0 {
		tag.SortNum = 999
	}
	
	return s.repoFactory.GetCmsTagRepository().Create(tag)
}

// UpdateTag 更新标签
func (s *CmsTagService) UpdateTag(tag *cms.CmsTag) error {
	if tag.ID <= 0 {
		return errors.New("无效的标签ID")
	}
	if tag.Title == "" {
		return errors.New("标签名称不能为空")
	}
	
	// 检查标签是否存在
	existing, err := s.repoFactory.GetCmsTagRepository().FindByID(tag.ID)
	if err != nil {
		return errors.New("标签不存在")
	}
	
	// 保留原始创建时间
	tag.CreatedTime = existing.CreatedTime
	
	return s.repoFactory.GetCmsTagRepository().Update(tag)
}

// GetTagByID 根据ID获取标签
func (s *CmsTagService) GetTagByID(id int64) (*cms.CmsTag, error) {
	if id <= 0 {
		return nil, errors.New("无效的标签ID")
	}
	return s.repoFactory.GetCmsTagRepository().FindByID(id)
}

// GetTagsByState 根据状态获取标签
func (s *CmsTagService) GetTagsByState(state int8) ([]cms.CmsTag, error) {
	if state < 0 || state > 2 {
		return nil, errors.New("无效的状态值")
	}
	return s.repoFactory.GetCmsTagRepository().FindByState(state)
}

// IncrementReadNum 增加标签阅读量
func (s *CmsTagService) IncrementReadNum(id int64) error {
	if id <= 0 {
		return errors.New("无效的标签ID")
	}
	return s.repoFactory.GetCmsTagRepository().IncrementReadNum(id)
}

// SearchTags 搜索标签
func (s *CmsTagService) SearchTags(keyword string) ([]cms.CmsTag, error) {
	if keyword == "" || len(keyword) > 50 {
		return nil, errors.New("关键词长度必须在1-50之间")
	}
	return s.repoFactory.GetCmsTagRepository().Search(keyword)
}