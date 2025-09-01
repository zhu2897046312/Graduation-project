package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpOrderReceiveAddressRepository struct {
	*base.BaseRepository
}

func NewSpOrderReceiveAddressRepository(DB *gorm.DB) *SpOrderReceiveAddressRepository {
	return &SpOrderReceiveAddressRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建收货地址
func (r *SpOrderReceiveAddressRepository) Create(address *sp.SpOrderReceiveAddress) error {
	return r.DB.Create(address).Error
}

// 更新收货地址
func (r *SpOrderReceiveAddressRepository) Update(address *sp.SpOrderReceiveAddress) error {
	return r.DB.Save(address).Error
}

// 根据订单ID获取收货地址
func (r *SpOrderReceiveAddressRepository) FindByOrderID(orderID uint) (*sp.SpOrderReceiveAddress, error) {
	var address sp.SpOrderReceiveAddress
	err := r.DB.Where("order_id = ?", orderID).First(&address).Error
	return &address, err
}

// 根据用户邮箱获取收货地址
func (r *SpOrderReceiveAddressRepository) FindByEmail(email string) ([]sp.SpOrderReceiveAddress, error) {
	var addresses []sp.SpOrderReceiveAddress
	err := r.DB.Where("email = ?", email).Find(&addresses).Error
	return addresses, err
}