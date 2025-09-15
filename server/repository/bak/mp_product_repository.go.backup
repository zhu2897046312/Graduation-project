package repository

import (
	"gorm.io/gorm"
	"server/models/mp"
	"server/models/common"
)

type MpProductRepository struct {
	*BaseRepository
}

func NewMpProductRepository(db *gorm.DB) *MpProductRepository {
	return &MpProductRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建产品
func (r *MpProductRepository) Create(product *mp.MpProduct) error {
	return r.db.Create(product).Error
}

// 更新产品
func (r *MpProductRepository) Update(product *mp.MpProduct) error {
	return r.db.Save(product).Error
}

// 根据ID获取产品
func (r *MpProductRepository) FindByID(id common.MyID) (*mp.MpProduct, error) {
	var product mp.MpProduct
	err := r.db.First(&product, id).Error
	return &product, err
}

// 根据产品类型获取产品列表
func (r *MpProductRepository) FindByType(productType int8) ([]mp.MpProduct, error) {
	var products []mp.MpProduct
	err := r.db.Where("product_type = ? AND state = 1", productType).
		Order("sort_num ASC").
		Find(&products).Error
	return products, err
}

// 根据终端类型获取产品列表
func (r *MpProductRepository) FindByTerminal(terminalType int8) ([]mp.MpProduct, error) {
	var products []mp.MpProduct
	err := r.db.Where("terminal_type = ? AND state = 1", terminalType).
		Order("sort_num ASC").
		Find(&products).Error
	return products, err
}

// 根据产品代码获取产品
func (r *MpProductRepository) FindByCode(code string) (*mp.MpProduct, error) {
	var product mp.MpProduct
	err := r.db.Where("code = ?", code).First(&product).Error
	return &product, err
}

// 更新产品状态
func (r *MpProductRepository) UpdateState(id common.MyID, state int) error {
	return r.db.Model(&mp.MpProduct{}).
		Where("id = ?", id).
		Update("state", state).Error
}