package service

import (
	"errors"
	"server/models/mp"
	"server/models/common"
)

type MpOrderService struct {
	*Service
}

func NewMpOrderService(base *Service) *MpOrderService {
	return &MpOrderService{Service: base}
}

// CreateOrder 创建订单
func (s *MpOrderService) CreateOrder(order *mp.MpOrder) error {
	if order.UserID <= 0 {
		return errors.New("无效的用户ID")
	}
	if order.ProductID <= 0 {
		return errors.New("无效的产品ID")
	}
	if order.PayPrice <= 0 {
		return errors.New("订单金额必须大于0")
	}
	
	// 设置默认状态
	if order.State == 0 {
		order.State = 0 // 默认待支付
	}
	
	return s.repoFactory.GetMpOrderRepository().Create(order)
}

// UpdateOrder 更新订单
func (s *MpOrderService) UpdateOrder(order *mp.MpOrder) error {
	if order.ID == 0 {
		return errors.New("无效的订单ID")
	}
	
	// 检查订单是否存在
	existing, err := s.repoFactory.GetMpOrderRepository().FindByID(order.ID)
	if err != nil {
		return errors.New("订单不存在")
	}
	
	// 保留原始创建时间
	order.CreatedTime = existing.CreatedTime
	
	return s.repoFactory.GetMpOrderRepository().Update(order)
}

// GetOrderByID 根据订单ID获取订单
func (s *MpOrderService) GetOrderByID(id common.MyID) (*mp.MpOrder, error) {
	if id == 0 {
		return nil, errors.New("订单ID不能为空")
	}
	return s.repoFactory.GetMpOrderRepository().FindByID(id)
}

// GetOrdersByUserID 根据用户ID获取订单列表
func (s *MpOrderService) GetOrdersByUserID(userID common.MyID) ([]mp.MpOrder, error) {
	if userID <= 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetMpOrderRepository().FindByUserID(userID)
}

// GetOrdersByState 根据状态获取订单列表
func (s *MpOrderService) GetOrdersByState(state int8) ([]mp.MpOrder, error) {
	if state < 0 || state > 3 {
		return nil, errors.New("无效的订单状态")
	}
	return s.repoFactory.GetMpOrderRepository().FindByState(state)
}

// UpdateOrderState 更新订单状态
func (s *MpOrderService) UpdateOrderState(id common.MyID, state int8) error {
	if id == 0 {
		return errors.New("订单ID不能为空")
	}
	if state < 0 || state > 3 {
		return errors.New("无效的订单状态")
	}
	
	return s.repoFactory.GetMpOrderRepository().UpdateState(id, state)
}

// GetOrderByThirdID 根据第三方支付ID获取订单
func (s *MpOrderService) GetOrderByThirdID(thirdID common.MyID) (*mp.MpOrder, error) {
	if thirdID == 0 {
		return nil, errors.New("第三方支付ID不能为空")
	}
	return s.repoFactory.GetMpOrderRepository().FindByThirdID(thirdID)
}