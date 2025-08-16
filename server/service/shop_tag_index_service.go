package service

import (
	"errors"
	"server/models/shop"
)

type ShopTagIndexService struct {
	*Service
}

func NewShopTagIndexService(base *Service) *ShopTagIndexService {
	return &ShopTagIndexService{Service: base}
}

// CreateTagIndex 创建标签关联
func (s *ShopTagIndexService) CreateTagIndex(index *shop.ShopTagIndex) error {
	if index.ProductID == 0 {
		return errors.New("商品ID不能为空")
	}
	if index.TagID == 0 {
		return errors.New("标签ID不能为空")
	}
	
	return s.repoFactory.GetShopTagIndexRepository().Create(index)
}

// DeleteTagIndex 删除标签关联
func (s *ShopTagIndexService) DeleteTagIndex(productID, tagID int) error {
	if productID == 0 {
		return errors.New("商品ID不能为空")
	}
	if tagID == 0 {
		return errors.New("标签ID不能为空")
	}
	return s.repoFactory.GetShopTagIndexRepository().Delete(productID, tagID)
}

// GetTagIndicesByProductID 根据产品ID获取标签关联
func (s *ShopTagIndexService) GetTagIndicesByProductID(productID int) ([]shop.ShopTagIndex, error) {
	if productID == 0 {
		return nil, errors.New("商品ID不能为空")
	}
	return s.repoFactory.GetShopTagIndexRepository().FindByProductID(productID)
}

// GetTagIndicesByTagID 根据标签ID获取产品关联
func (s *ShopTagIndexService) GetTagIndicesByTagID(tagID int) ([]shop.ShopTagIndex, error) {
	if tagID == 0 {
		return nil, errors.New("标签ID不能为空")
	}
	return s.repoFactory.GetShopTagIndexRepository().FindByTagID(tagID)
}

// UpdateTagSortNum 更新标签关联排序
func (s *ShopTagIndexService) UpdateTagSortNum(id, sortNum int) error {
	if id == 0 {
		return errors.New("关联ID不能为空")
	}
	if sortNum < 0 {
		return errors.New("排序值不能为负数")
	}
	return s.repoFactory.GetShopTagIndexRepository().UpdateSortNum(id, sortNum)
}

// DeleteAllTagsByProductID 删除产品的所有标签关联
func (s *ShopTagIndexService) DeleteAllTagsByProductID(productID int) error {
	if productID == 0 {
		return errors.New("商品ID不能为空")
	}
	return s.repoFactory.GetShopTagIndexRepository().DeleteByProductID(productID)
}