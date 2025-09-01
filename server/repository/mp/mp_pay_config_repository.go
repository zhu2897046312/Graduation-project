package mp

import (
	"gorm.io/gorm"
	"server/models/mp"
	"server/repository/base"
)

type MpPayConfigRepository struct {
	*base.BaseRepository
}

func NewMpPayConfigRepository(DB *gorm.DB) *MpPayConfigRepository {
	return &MpPayConfigRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建支付配置
func (r *MpPayConfigRepository) Create(config *mp.MpPayConfig) error {
	return r.DB.Create(config).Error
}

// 更新支付配置
func (r *MpPayConfigRepository) Update(config *mp.MpPayConfig) error {
	return r.DB.Save(config).Error
}

// 根据ID获取支付配置
func (r *MpPayConfigRepository) FindByID(id int64) (*mp.MpPayConfig, error) {
	var config mp.MpPayConfig
	err := r.DB.First(&config, id).Error
	return &config, err
}

// 获取所有启用的支付配置
func (r *MpPayConfigRepository) FindActive() ([]mp.MpPayConfig, error) {
	var configs []mp.MpPayConfig
	err := r.DB.Where("state = 1").
		Order("sort_num ASC").
		Find(&configs).Error
	return configs, err
}

// 根据Code获取支付配置
func (r *MpPayConfigRepository) FindByCode(code string) (*mp.MpPayConfig, error) {
	var config mp.MpPayConfig
	err := r.DB.Where("code = ?", code).First(&config).Error
	return &config, err
}

// 更新支付配置状态
func (r *MpPayConfigRepository) UpdateState(id int64, state int8) error {
	return r.DB.Model(&mp.MpPayConfig{}).
		Where("id = ?", id).
		Update("state", state).Error
}