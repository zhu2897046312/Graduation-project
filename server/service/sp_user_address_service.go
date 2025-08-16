package service

import (
	"errors"
	"server/models/sp"
	"time"
)

type SpUserAddressService struct {
	*Service
}

func NewSpUserAddressService(base *Service) *SpUserAddressService {
	return &SpUserAddressService{Service: base}
}

// CreateAddress 创建用户地址
func (s *SpUserAddressService) CreateAddress(address *sp.SpUserAddress) error {
	if address.UserID == 0 {
		return errors.New("用户ID不能为空")
	}
	if address.Title == "" {
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
	
	return s.repoFactory.GetSpUserAddressRepository().Create(address)
}

// UpdateAddress 更新用户地址
func (s *SpUserAddressService) UpdateAddress(address *sp.SpUserAddress) error {
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
	existing, err := s.repoFactory.GetSpUserAddressRepository().FindByID(address.ID)
	if err != nil {
		return errors.New("地址不存在")
	}
	
	address.CreatedTime = existing.CreatedTime
	address.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpUserAddressRepository().Update(address)
}

// GetAddressByID 根据ID获取地址
func (s *SpUserAddressService) GetAddressByID(id uint) (*sp.SpUserAddress, error) {
	if id == 0 {
		return nil, errors.New("无效的地址ID")
	}
	return s.repoFactory.GetSpUserAddressRepository().FindByID(id)
}

// GetAddressesByUserID 根据用户ID获取地址列表
func (s *SpUserAddressService) GetAddressesByUserID(userID uint) ([]sp.SpUserAddress, error) {
	if userID == 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetSpUserAddressRepository().FindByUserID(userID)
}

// GetDefaultAddress 获取用户的默认地址
func (s *SpUserAddressService) GetDefaultAddress(userID uint) (*sp.SpUserAddress, error) {
	if userID == 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetSpUserAddressRepository().FindDefaultByUserID(userID)
}

// SetDefaultAddress 设置默认地址
func (s *SpUserAddressService) SetDefaultAddress(id uint, userID uint) error {
	if id == 0 {
		return errors.New("无效的地址ID")
	}
	if userID == 0 {
		return errors.New("无效的用户ID")
	}
	return s.repoFactory.GetSpUserAddressRepository().SetDefaultAddress(id, userID)
}

// DeleteAddress 删除地址
func (s *SpUserAddressService) DeleteAddress(id uint) error {
	if id == 0 {
		return errors.New("无效的地址ID")
	}
	return s.repoFactory.GetSpUserAddressRepository().Delete(id)
}