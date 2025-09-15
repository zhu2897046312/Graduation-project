package service

import (
	"errors"
	"server/models/sp"
	"time"
)

type SpOrderOperateHistoryService struct {
	*Service
}

func NewSpOrderOperateHistoryService(base *Service) *SpOrderOperateHistoryService {
	return &SpOrderOperateHistoryService{Service: base}
}

// CreateHistory 创建操作历史
func (s *SpOrderOperateHistoryService) CreateHistory(history *sp.SpOrderOperateHistory) error {
	if history.OrderID == 0 {
		return errors.New("订单ID不能为空")
	}
	if history.OperateUser == "" {
		return errors.New("操作人不能为空")
	}
	if history.Remark == "" {
		return errors.New("操作类型不能为空")
	}
	
	history.CreatedTime = time.Now()
	return s.repoFactory.GetSpOrderOperateHistoryRepository().Create(history)
}

// GetHistoriesByOrderID 根据订单ID获取操作历史
func (s *SpOrderOperateHistoryService) GetHistoriesByOrderID(orderID uint) ([]sp.SpOrderOperateHistory, error) {
	if orderID == 0 {
		return nil, errors.New("无效的订单ID")
	}
	return s.repoFactory.GetSpOrderOperateHistoryRepository().FindByOrderID(orderID)
}

// GetHistoriesByOperateUser 根据操作人获取操作历史
func (s *SpOrderOperateHistoryService) GetHistoriesByOperateUser(user string) ([]sp.SpOrderOperateHistory, error) {
	if user == "" || len(user) > 50 {
		return nil, errors.New("操作人名称长度必须在1-50之间")
	}
	return s.repoFactory.GetSpOrderOperateHistoryRepository().FindByOperateUser(user)
}