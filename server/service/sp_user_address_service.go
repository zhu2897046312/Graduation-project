package service

import (
	"errors"
	"server/models/sp"
	"server/models/common"
	"time"
)

type SpUserAddressService struct {
	*Service
}

func NewSpUserAddressService(base *Service) *SpUserAddressService {
	return &SpUserAddressService{Service: base}
}

func (s *SpUserAddressService) Create(address *sp.SpUserAddress) error {
	if address.UserID == 0 {
		return errors.New("用户ID不能为空")
	}
	if address.Title == "" {
		return errors.New("地址标题不能为空")
	}
	if address.FirstName == "" {
		return errors.New("姓不能为空")
	}
	if address.LastName == "" {
		return errors.New("名不能为空")
	}
	if address.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if address.Phone == "" {
		return errors.New("电话不能为空")
	}
	if address.Province == "" {
		return errors.New("省不能为空")
	}
	if address.City == "" {
		return errors.New("市不能为空")
	}
	if address.Region == "" {
		return errors.New("区不能为空")
	}
	if address.DetailAddress == "" {
		return errors.New("详细地址不能为空")
	}
	if address.Country == "" {
		return errors.New("国家不能为空")
	}
	if address.PostalCode == "" {
		return errors.New("邮政编码不能为空")
	}		
	
	address.CreatedTime = time.Now()
	address.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpUserAddressRepository().Create(address)
}

func (s *SpUserAddressService) Update(address *sp.SpUserAddress) error {
	if address.ID == 0 {
		return errors.New("地址ID不能为空")
	}
	_,err := s.repoFactory.GetSpUserAddressRepository().FindByID(address.ID)
	if err == nil{
		return errors.New("该id地址已经存在")
	}
	address.UpdatedTime = time.Now()
	return s.repoFactory.GetSpUserAddressRepository().Update(address)
}

func (s *SpUserAddressService) Delete(id common.MyID) error {
	_,err := s.repoFactory.GetSpUserAddressRepository().FindByID(id)
	if err != nil{
		return errors.New("该id地址不存在存在")
	}
	return s.repoFactory.GetSpUserAddressRepository().Delete(id)
}

func (s *SpUserAddressService) ListAddress(params *sp.SpUserAddressListParam) ([]*sp.SpUserAddress, int64, error) {
	return s.repoFactory.GetSpUserAddressRepository().ListAddress(params)
}

func (s *SpUserAddressService) GetAddressByID(id common.MyID) (*sp.SpUserAddress, error) {
	return s.repoFactory.GetSpUserAddressRepository().FindByID(id)
}