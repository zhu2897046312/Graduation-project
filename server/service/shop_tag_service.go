package service

import (
	"errors"
	"server/models/shop"
)

type ShopTagService struct {
	*Service
}

func NewShopTagService(base *Service) *ShopTagService {
	return &ShopTagService{Service: base}
}

// CreateTag 创建商品标签
func (s *ShopTagService) CreateTag(tag *shop.ShopTag) error {
	if tag.Title == "" {
		return errors.New("标签名称不能为空")
	}
	
	// 设置默认排序值
	if tag.SortNum == 0 {
		tag.SortNum = 999
	}
	
	return s.repoFactory.GetShopTagRepository().Create(tag)
}

// UpdateTag 更新商品标签
func (s *ShopTagService) UpdateTag(tag *shop.ShopTag) error {
	if tag.ID == 0 {
		return errors.New("标签ID不能为空")
	}
	if tag.Title == "" {
		return errors.New("标签名称不能为空")
	}
	
	// 检查标签是否存在
	existing, err := s.repoFactory.GetShopTagRepository().FindByID(tag.ID)
	if err != nil {
		return errors.New("标签不存在")
	}
	
	// 保留原始创建时间
	tag.CreatedTime = existing.CreatedTime
	
	return s.repoFactory.GetShopTagRepository().Update(tag)
}

// GetTagByID 根据ID获取标签
func (s *ShopTagService) GetTagByID(id int) (*shop.ShopTag, error) {
	if id == 0 {
		return nil, errors.New("标签ID不能为空")
	}
	return s.repoFactory.GetShopTagRepository().FindByID(id)
}

// GetTagsByState 根据状态获取标签列表
func (s *ShopTagService) GetTagsByState(state int8) ([]shop.ShopTag, error) {
	if state < 0 || state > 2 {
		return nil, errors.New("无效的状态值")
	}
	return s.repoFactory.GetShopTagRepository().FindByState(state)
}

// SearchTagsByMatchWord 根据匹配词获取标签
func (s *ShopTagService) SearchTagsByMatchWord(matchWord string) ([]shop.ShopTag, error) {
	if matchWord == "" || len(matchWord) > 50 {
		return nil, errors.New("匹配词长度必须在1-50之间")
	}
	return s.repoFactory.GetShopTagRepository().FindByMatchWord(matchWord)
}

// IncrementTagReadNum 增加标签阅读量
func (s *ShopTagService) IncrementTagReadNum(id int) error {
	if id == 0 {
		return errors.New("标签ID不能为空")
	}
	return s.repoFactory.GetShopTagRepository().IncrementReadNum(id)
}

// ListTags 获取所有标签（分页）
func (s *ShopTagService) ListTags(params shop.TagQueryParams) ([]shop.ShopTag, int64, error) {
	// 验证分页参数
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 10
	}
	
	return s.repoFactory.GetShopTagRepository().ListWithPagination(params)
}