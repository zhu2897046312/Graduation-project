package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpUserAddressRepository struct {
	*base.BaseRepository
}

func NewSpUserAddressRepository(DB *gorm.DB) *SpUserAddressRepository {
	return &SpUserAddressRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建用户地址
func (r *SpUserAddressRepository) Create(address *sp.SpUserAddress) error {
	return r.DB.Create(address).Error
}

// 更新用户地址
func (r *SpUserAddressRepository) Update(address *sp.SpUserAddress) error {
	return r.DB.Save(address).Error
}

// 根据ID获取地址
func (r *SpUserAddressRepository) FindByID(id uint) (*sp.SpUserAddress, error) {
	var address sp.SpUserAddress
	err := r.DB.First(&address, id).Error
	return &address, err
}

// 根据用户ID获取地址列表
func (r *SpUserAddressRepository) FindByUserID(userID uint) ([]sp.SpUserAddress, error) {
	var addresses []sp.SpUserAddress
	err := r.DB.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&addresses).Error
	return addresses, err
}

// 获取用户的默认地址
func (r *SpUserAddressRepository) FindDefaultByUserID(userID uint) (*sp.SpUserAddress, error) {
	var address sp.SpUserAddress
	err := r.DB.Where("user_id = ? AND default_status = 1", userID).
		First(&address).Error
	return &address, err
}

// 设置默认地址
func (r *SpUserAddressRepository) SetDefaultAddress(id uint, userID uint) error {
	// 先重置用户的所有默认地址状态
	if err := r.DB.Model(&sp.SpUserAddress{}).
		Where("user_id = ?", userID).
		Update("default_status", false).Error; err != nil {
		return err
	}
	
	// 设置当前地址为默认
	return r.DB.Model(&sp.SpUserAddress{}).
		Where("id = ?", id).
		Update("default_status", true).Error
}

// 删除地址
func (r *SpUserAddressRepository) Delete(id uint) error {
	return r.DB.Delete(&sp.SpUserAddress{}, id).Error
}