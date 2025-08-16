package service

import (
	"errors"
	"server/models/shop"
)

type ShopTagMateService struct {
	*Service
}

func NewShopTagMateService(base *Service) *ShopTagMateService {
	return &ShopTagMateService{Service: base}
}

// CreateTagMate 创建标签元数据
func (s *ShopTagMateService) CreateTagMate(mate *shop.ShopTagMate) error {
	if mate.ID == 0 {
		return errors.New("标签ID不能为空")
	}
	
	return s.repoFactory.GetShopTagMateRepository().Create(mate)
}

// UpdateTagMate 更新标签元数据
func (s *ShopTagMateService) UpdateTagMate(mate *shop.ShopTagMate) error {
	if mate.ID == 0 {
		return errors.New("元数据ID不能为空")
	}
	
	// 检查元数据是否存在
	_, err := s.repoFactory.GetShopTagMateRepository().FindByID(mate.ID)
	if err != nil {
		return errors.New("元数据不存在")
	}
	
	return s.repoFactory.GetShopTagMateRepository().Update(mate)
}

// GetTagMateByID 根据ID获取元数据
func (s *ShopTagMateService) GetTagMateByID(id int) (*shop.ShopTagMate, error) {
	if id == 0 {
		return nil, errors.New("元数据ID不能为空")
	}
	return s.repoFactory.GetShopTagMateRepository().FindByID(id)
}

// UpdateTagSEO 更新SEO信息
func (s *ShopTagMateService) UpdateTagSEO(id int, title, keyword, description string) error {
	if id == 0 {
		return errors.New("元数据ID不能为空")
	}
	if title == "" && keyword == "" && description == "" {
		return errors.New("至少需要更新一个SEO字段")
	}
	return s.repoFactory.GetShopTagMateRepository().UpdateSEO(id, title, keyword, description)
}

// UpdateTagContent 更新标签内容
func (s *ShopTagMateService) UpdateTagContent(id int, content string) error {
	if id == 0 {
		return errors.New("元数据ID不能为空")
	}
	if content == "" {
		return errors.New("内容不能为空")
	}
	return s.repoFactory.GetShopTagMateRepository().UpdateContent(id, content)
}