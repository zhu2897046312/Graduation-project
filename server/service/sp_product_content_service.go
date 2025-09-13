package service

import (
	"errors"
	"server/models/sp"
	"server/models/common"
)

type SpProductContentService struct {
	*Service
}

func NewSpProductContentService(base *Service) *SpProductContentService {
	return &SpProductContentService{Service: base}
}

// CreateContent 创建商品内容
func (s *SpProductContentService) CreateContent(content *sp.SpProductContent) error {
	if content.ProductID == 0 {
		return errors.New("商品ID不能为空")
	}
	return s.repoFactory.GetSpProductContentRepository().Create(content)
}

// UpdateContent 更新商品内容
func (s *SpProductContentService) UpdateContent(content *sp.SpProductContent) error {
	if content.ProductID == 0 {
		return errors.New("商品ID不能为空")
	}
	return s.repoFactory.GetSpProductContentRepository().Update(content)
}

// GetContentByProductID 根据商品ID获取内容
func (s *SpProductContentService) GetContentByProductID(productID common.MyID) (*sp.SpProductContent, error) {
	if productID == 0 {
		return nil, errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpProductContentRepository().FindByProductID(productID)
}

// UpdateSEO 更新SEO信息
func (s *SpProductContentService) UpdateSEO(productID common.MyID, title, keyword, description string) error {
	if productID == 0 {
		return errors.New("无效的商品ID")
	}
	if title == "" || keyword == "" || description == "" {
		return errors.New("SEO信息不能为空")
	}
	return s.repoFactory.GetSpProductContentRepository().UpdateSEO(productID, title, keyword, description)
}

// UpdateContentText 更新商品内容文本
func (s *SpProductContentService) UpdateContentText(productID common.MyID, content string) error {
	if productID == 0 {
		return errors.New("无效的商品ID")
	}
	if content == "" {
		return errors.New("内容不能为空")
	}
	return s.repoFactory.GetSpProductContentRepository().UpdateContent(productID, content)
}