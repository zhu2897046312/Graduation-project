package service

import (
	"errors"
	"server/models/cms"
	"time"
)

type CmsAssociatedPlaceService struct {
	*Service
}

func NewCmsAssociatedPlaceService(base *Service) *CmsAssociatedPlaceService {
	return &CmsAssociatedPlaceService{Service: base}
}

// CreatePlace 创建地点
func (s *CmsAssociatedPlaceService) CreatePlace(place *cms.CmsAssociatedPlace) error {
	// 验证必填字段
	if place.Title == "" {
		return errors.New("地点名称不能为空")
	}
	
	// 验证状态值
	if place.State != 0 && place.State != 1 && place.State != 2 {
		return errors.New("无效的状态值")
	}
	
	// 设置默认状态
	if place.State == 0 {
		place.State = 2 // 默认未发布
	}
	
	// 设置创建时间和更新时间
	now := time.Now()
	place.CreatedTime = now
	place.UpdatedTime = now
	
	return s.repoFactory.GetCmsPlaceRepository().Create(place)
}

// UpdatePlace 更新地点
func (s *CmsAssociatedPlaceService) UpdatePlace(place *cms.CmsAssociatedPlace) error {
	// 验证必填字段
	if place.Title == "" {
		return errors.New("地点名称不能为空")
	}
	
	// 验证状态值
	if place.State != 1 && place.State != 2 {
		return errors.New("无效的状态值")
	}
	
	// 获取现有记录以保留创建时间
	existing, err := s.repoFactory.GetCmsPlaceRepository().FindByID(place.ID)
	if err != nil {
		return errors.New("地点不存在")
	}
	
	// 保留原始创建时间
	place.CreatedTime = existing.CreatedTime
	// 设置更新时间
	place.UpdatedTime = time.Now()
	
	return s.repoFactory.GetCmsPlaceRepository().Update(place)
}

// GetPlaceByID 根据ID获取地点
func (s *CmsAssociatedPlaceService) GetPlaceByID(id int64) (*cms.CmsAssociatedPlace, error) {
	if id <= 0 {
		return nil, errors.New("无效的地点ID")
	}
	return s.repoFactory.GetCmsPlaceRepository().FindByID(id)
}

// ListPlaces 分页获取地点列表
func (s *CmsAssociatedPlaceService) ListPlaces(page, pageSize int) ([]cms.CmsAssociatedPlace, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	return s.repoFactory.GetCmsPlaceRepository().ListWithPagination(page, pageSize)
}

// SearchPlacesByInitial 根据拼音首字母搜索地点
func (s *CmsAssociatedPlaceService) SearchPlacesByInitial(initial string) ([]cms.CmsAssociatedPlace, error) {
	if len(initial) == 0 || len(initial) > 10 {
		return nil, errors.New("首字母长度必须在1-10之间")
	}
	
	return s.repoFactory.GetCmsPlaceRepository().SearchByInitialPinyin(initial)
}

// UpdatePlaceState 更新地点状态
func (s *CmsAssociatedPlaceService) UpdatePlaceState(id int64, state int8) error {
	// 验证状态值
	if state != 1 && state != 2 {
		return errors.New("无效的状态值")
	}
	
	// 获取现有记录
	place, err := s.GetPlaceByID(id)
	if err != nil {
		return err
	}
	
	// 更新状态
	place.State = state
	return s.UpdatePlace(place)
}

// DeletePlace 删除地点
func (s *CmsAssociatedPlaceService) DeletePlace(id int64) error {
	if id <= 0 {
		return errors.New("无效的地点ID")
	}
	
	// 检查地点是否存在
	_, err := s.GetPlaceByID(id)
	if err != nil {
		return err
	}
	
	// 执行删除操作
	// 注意: 实际项目中可能需要软删除或关联数据检查
	return s.repoFactory.GetCmsPlaceRepository().Delete(id)
}