package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/models/common"
)

type SpOrderReceiveAddressRepository struct {
	*BaseRepository
}

func NewSpOrderReceiveAddressRepository(db *gorm.DB) *SpOrderReceiveAddressRepository {
	return &SpOrderReceiveAddressRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建收货地址
func (r *SpOrderReceiveAddressRepository) Create(address *sp.SpOrderReceiveAddress) error {
	return r.db.Create(address).Error
}

// 更新收货地址
func (r *SpOrderReceiveAddressRepository) Update(address *sp.SpOrderReceiveAddress) error {
	return r.db.Save(address).Error
}

// 根据订单ID获取收货地址
func (r *SpOrderReceiveAddressRepository) FindByOrderID(orderID common.MyID) (*sp.SpOrderReceiveAddress, error) {
	var address sp.SpOrderReceiveAddress
	err := r.db.Where("order_id = ?", orderID).First(&address).Error
	return &address, err
}

// 根据用户邮箱获取收货地址
func (r *SpOrderReceiveAddressRepository) FindByEmail(email string) ([]sp.SpOrderReceiveAddress, error) {
	var addresses []sp.SpOrderReceiveAddress
	err := r.db.Where("email = ?", email).Find(&addresses).Error
	return addresses, err
}