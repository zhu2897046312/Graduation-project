package repository

import (
	"server/models/sp"
	"server/models/common"
	"gorm.io/gorm"
)

type SpUserAddressRepository struct {
	*BaseRepository
}

func NewSpUserAddressRepository(db *gorm.DB) *SpUserAddressRepository {
	return &SpUserAddressRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *SpUserAddressRepository) FindByID(id common.MyID) (*sp.SpUserAddress, error) {
	var address sp.SpUserAddress
	err := r.db.Model(&sp.SpUserAddress{}).Where("id = ?", id).First(&address).Error
	return &address, err
}

func (r *SpUserAddressRepository) Create(address *sp.SpUserAddress) error {
	return r.db.Model(&sp.SpUserAddress{}).Create(address).Error
}

func (r *SpUserAddressRepository) Update(address *sp.SpUserAddress) error {
	return r.db.Model(&sp.SpUserAddress{}).Updates(address).Error
}

func (r *SpUserAddressRepository) Delete(id common.MyID) error {
	return r.db.Model(&sp.SpUserAddress{}).Where("id = ?", id).Delete(&sp.SpUserAddress{}).Error
}

func (r *SpUserAddressRepository) ListAddress(params *sp.SpUserAddressListParam) ([]*sp.SpUserAddress, int64, error) {
	var addresses []*sp.SpUserAddress
	var total int64

	// 构建基础查询
	query := r.db.Model(&sp.SpUserAddress{}).Where("deleted_time IS NULL")

	if params.UserID != 0 {
		query = query.Where("user_id = ?", params.UserID)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 处理分页逻辑
	if params.PageSize > 0 {
		// 确保页码有效
		if params.Page < 1 {
			params.Page = 1
		}
		
		// 限制最大分页大小
		if params.PageSize > 100 {
			params.PageSize = 100
		}

		offset := (params.Page - 1) * params.PageSize
		query = query.Offset(offset).Limit(params.PageSize)
	}

	// 执行查询
	err := query.Find(&addresses).Error
	return addresses, total, err
}
