package repository

import (
	"gorm.io/gorm"
	"server/models/mp"
)

type MpPayConfigRepository struct {
	*BaseRepository
}

func NewMpPayConfigRepository(db *gorm.DB) *MpPayConfigRepository {
	return &MpPayConfigRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建支付配置
func (r *MpPayConfigRepository) Create(config *mp.MpPayConfig) error {
	return r.db.Create(config).Error
}

// 更新支付配置
func (r *MpPayConfigRepository) Update(config *mp.MpPayConfig) error {
	return r.db.Save(config).Error
}

// 根据ID获取支付配置
func (r *MpPayConfigRepository) FindByID(id int64) (*mp.MpPayConfig, error) {
	var config mp.MpPayConfig
	err := r.db.First(&config, id).Error
	return &config, err
}

// 获取所有启用的支付配置
func (r *MpPayConfigRepository) FindActive() ([]mp.MpPayConfig, error) {
	var configs []mp.MpPayConfig
	err := r.db.Where("state = 1").
		Order("sort_num ASC").
		Find(&configs).Error
	return configs, err
}

// 根据Code获取支付配置
func (r *MpPayConfigRepository) FindByCode(code string) (*mp.MpPayConfig, error) {
	var config mp.MpPayConfig
	err := r.db.Where("code = ?", code).First(&config).Error
	return &config, err
}

// 更新支付配置状态
func (r *MpPayConfigRepository) UpdateState(id int64, state int8) error {
	return r.db.Model(&mp.MpPayConfig{}).
		Where("id = ?", id).
		Update("state", state).Error
}