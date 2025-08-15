package service

import (
	"errors"
	"server/models/cms"
)

type CmsScenicSpotService struct {
	*Service
}

func NewCmsScenicSpotService(base *Service) *CmsScenicSpotService {
	return &CmsScenicSpotService{Service: base}
}

// CreateScenicSpot 创建景点
func (s *CmsScenicSpotService) CreateScenicSpot(spot *cms.CmsScenicSpot) error {
	if spot.Title == "" {
		return errors.New("景点名称不能为空")
	}
	if spot.AssociatedPlaceID <= 0 {
		return errors.New("关联地点ID不能为空")
	}
	
	return s.repoFactory.GetCmsScenicSpotRepository().Create(spot)
}

// UpdateScenicSpot 更新景点
func (s *CmsScenicSpotService) UpdateScenicSpot(spot *cms.CmsScenicSpot) error {
	if spot.ID <= 0 {
		return errors.New("无效的景点ID")
	}
	if spot.Title == "" {
		return errors.New("景点名称不能为空")
	}
	
	// 检查景点是否存在
	existing, err := s.repoFactory.GetCmsScenicSpotRepository().FindByID(spot.ID)
	if err != nil {
		return errors.New("景点不存在")
	}
	
	// 保留原始创建时间
	spot.CreatedTime = existing.CreatedTime
	
	return s.repoFactory.GetCmsScenicSpotRepository().Update(spot)
}

// GetScenicSpotByID 根据ID获取景点
func (s *CmsScenicSpotService) GetScenicSpotByID(id int64) (*cms.CmsScenicSpot, error) {
	if id <= 0 {
		return nil, errors.New("无效的景点ID")
	}
	return s.repoFactory.GetCmsScenicSpotRepository().FindByID(id)
}

// GetScenicSpotsByPlaceID 根据地点ID获取景点
func (s *CmsScenicSpotService) GetScenicSpotsByPlaceID(placeID int64) ([]cms.CmsScenicSpot, error) {
	if placeID <= 0 {
		return nil, errors.New("无效的地点ID")
	}
	return s.repoFactory.GetCmsScenicSpotRepository().FindByPlaceID(placeID)
}

// IncrementReadNum 增加景点阅读量
func (s *CmsScenicSpotService) IncrementReadNum(id int64) error {
	if id <= 0 {
		return errors.New("无效的景点ID")
	}
	return s.repoFactory.GetCmsScenicSpotRepository().IncrementReadNum(id)
}

// ListScenicSpots 分页获取景点列表
func (s *CmsScenicSpotService) ListScenicSpots(page, pageSize int) ([]cms.CmsScenicSpot, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	
	return s.repoFactory.GetCmsScenicSpotRepository().ListWithPagination(page, pageSize)
}