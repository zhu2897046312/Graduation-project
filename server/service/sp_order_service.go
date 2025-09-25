package service

import (
	"errors"
	"server/models/sp"
	"server/models/common"
	"time"
)

type SpOrderService struct {
	*Service
}

func NewSpOrderService(base *Service) *SpOrderService {
	return &SpOrderService{Service: base}
}

// CreateOrder 创建订单
func (s *SpOrderService) CreateOrder(order *sp.SpOrder) error {
	if order.UserID == 0 && order.Email == ""  {
		return errors.New("用户ID和邮箱不能同时为空")
	}
	if order.Code == "" {
		return errors.New("订单号不能为空")
	}
	if order.TotalAmount <= 0 {
		return errors.New("订单总金额必须大于0")
	}
	if order.State == 0 {
		order.State = 1 // 默认待支付
	}
	
	order.CreatedTime = time.Now()
	order.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpOrderRepository().Create(order)
}

// UpdateOrder 更新订单
func (s *SpOrderService) UpdateOrder(order *sp.SpOrder) error {
	if order.ID == 0 {
		return errors.New("订单ID不能为空")
	}
	if order.TotalAmount <= 0 {
		return errors.New("订单总金额必须大于0")
	}
	if order.State != 1 && order.State != 2 && order.State != 3 && order.State != 4 && order.State != 5 {
		return errors.New("无效的订单状态")
	}
	
	// 保留原始创建时间
	existing, err := s.repoFactory.GetSpOrderRepository().FindByID(order.ID)
	if err != nil {
		return errors.New("订单不存在")
	}
	
	order.CreatedTime = existing.CreatedTime
	order.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpOrderRepository().Update(order)
}

// GetOrderByID 根据ID获取订单
func (s *SpOrderService) GetOrderByID(id common.MyID) (*sp.SpOrder, error) {
	if id == 0 {
		return nil, errors.New("无效的订单ID")
	}
	return s.repoFactory.GetSpOrderRepository().FindByID(id)
}

// GetOrderByCode 根据订单号获取订单
func (s *SpOrderService) GetOrderByCode(code string) (*sp.SpOrder, error) {
	if code == "" {
		return nil, errors.New("订单号不能为空")
	}
	return s.repoFactory.GetSpOrderRepository().FindByCode(code)
}

func (s *SpOrderService) GetByVisitorQueryCode(code string) (*sp.SpOrder, error) {
	if code == "" {
		return nil, errors.New("订单号不能为空")
	}
	return s.repoFactory.GetSpOrderRepository().FindByVisitorQueryCode(code)
}

// GetOrdersByUserID 根据用户ID获取订单列表
func (s *SpOrderService) GetOrdersByUserID(userID common.MyID) ([]sp.SpOrder, error) {
	if userID == 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetSpOrderRepository().FindByUserID(userID)
}

// GetOrdersByState 根据状态获取订单列表
func (s *SpOrderService) GetOrdersByState(state uint8) ([]sp.SpOrder, error) {
	if state != 1 && state != 2 && state != 3 && state != 4 && state != 5 {
		return nil, errors.New("无效的订单状态")
	}
	return s.repoFactory.GetSpOrderRepository().FindByState(state)
}

func (s *SpOrderService) List(param sp.ListOrdersQueryParam) ([]sp.SpOrder,int64, error) {
	return s.repoFactory.GetSpOrderRepository().ListWithPagination(param)
}
// UpdateOrderState 更新订单状态
func (s *SpOrderService) UpdateOrderState(code string, state common.MyState,remark string) error {
	if code == "" {
		return errors.New("订单ID不能为空")
	}
	_,err :=  s.repoFactory.GetSpOrderRepository().FindByVisitorQueryCode(code)
	if err != nil {
		return errors.New("订单不存在")
	}

	return s.repoFactory.GetSpOrderRepository().UpdateState(code, state, remark)
}



func (s *SpOrderService) UpdateOrderStateByOrderID(id common.MyID, state uint8,remark string) error {
	if id == 0 {
		return errors.New("订单ID不能为空")
	}
	_,err :=  s.repoFactory.GetSpOrderRepository().FindByID(id)
	if err != nil {
		return errors.New("订单不存在")
	}

	return s.repoFactory.GetSpOrderRepository().UpdateStateByID(id, state, remark)
}

// UpdateDeliveryInfo 更新物流信息
func (s *SpOrderService) UpdateDeliveryInfo(id common.MyID, company, sn string) error {
	if id  == 0 {
		return errors.New("订单ID不能为空")
	}
	if company == "" {
		return errors.New("物流公司不能为空")
	}
	if sn == "" {
		return errors.New("物流单号不能为空")
	}
	_,err :=  s.repoFactory.GetSpOrderRepository().FindByID(id)
	if err != nil {
		return errors.New("订单不存在")
	}
	return s.repoFactory.GetSpOrderRepository().UpdateDeliveryInfo(id, company, sn)
}

func (s *SpOrderService) DeleteOrder(id common.MyID) error {
	if id == 0 {
		return errors.New("订单ID不能为空")
	}
	_,err :=  s.repoFactory.GetSpOrderRepository().FindByID(id)
	if err != nil {
		return errors.New("订单不存在")
	}
	return s.repoFactory.GetSpOrderRepository().Delete(id)
}

func (s *SpOrderService) DeleteOrderByVisitorQueryCode(code string) error {
	if code == "" {
		return errors.New("订单ID不能为空")
	}
	_,err :=  s.repoFactory.GetSpOrderRepository().FindByVisitorQueryCode(code)
	if err != nil {
		return errors.New("订单不存在")
	}
	return s.repoFactory.GetSpOrderRepository().DeleteOrderByVisitorQueryCode(code)
}