package service

import (
	"errors"
	"server/models/sp"
	"server/models/common"
	"time"
)

type SpOrderReceiveAddressService struct {
	*Service
}

func NewSpOrderReceiveAddressService(base *Service) *SpOrderReceiveAddressService {
	return &SpOrderReceiveAddressService{Service: base}
}

// CreateAddress 创建收货地址
func (s *SpOrderReceiveAddressService) CreateAddress(address *sp.SpOrderReceiveAddress) error {
	if address.OrderID == 0 {
		return errors.New("订单ID不能为空")
	}
	if address.FirstName == "" {
		return errors.New("收货人姓名不能为空")
	}
	if address.Phone == "" {
		return errors.New("收货人电话不能为空")
	}
	if address.Province == "" || address.City == "" || address.DetailAddress == "" || address.PostalCode == "" {
		return errors.New("收货地址不完整")
	}
	
	address.CreatedTime = time.Now()
	address.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpOrderReceiveAddressRepository().Create(address)
}

// UpdateAddress 更新收货地址
func (s *SpOrderReceiveAddressService) UpdateAddress(address *sp.SpOrderReceiveAddress) error {
	if address.ID == 0 {
		return errors.New("地址ID不能为空")
	}
	if address.FirstName == "" {
		return errors.New("收货人姓名不能为空")
	}
	if address.Phone == "" {
		return errors.New("收货人电话不能为空")
	}
	if address.Province == "" || address.City == "" || address.DetailAddress == "" || address.PostalCode == "" {
		return errors.New("收货地址不完整")
	}
	
	// 保留原始创建时间
	existing, err := s.repoFactory.GetSpOrderReceiveAddressRepository().FindByOrderID(address.ID)
	if err != nil {
		return errors.New("地址不存在")
	}
	
	address.CreatedTime = existing.CreatedTime
	address.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpOrderReceiveAddressRepository().Update(address)
}

// GetAddressByOrderID 根据订单ID获取收货地址
func (s *SpOrderReceiveAddressService) GetAddressByOrderID(orderID common.MyID) (*sp.SpOrderReceiveAddress, error) {
	if orderID == 0 {
		return nil, errors.New("无效的订单ID")
	}
	return s.repoFactory.GetSpOrderReceiveAddressRepository().FindByOrderID(orderID)
}

// GetAddressesByEmail 根据邮箱获取收货地址
func (s *SpOrderReceiveAddressService) GetAddressesByEmail(email string) ([]sp.SpOrderReceiveAddress, error) {
	if email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	return s.repoFactory.GetSpOrderReceiveAddressRepository().FindByEmail(email)
}