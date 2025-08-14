package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
)

type SpProductRepository struct {
	*BaseRepository
}

func NewSpProductRepository(db *gorm.DB) *SpProductRepository {
	return &SpProductRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建商品
func (r *SpProductRepository) Create(product *sp.SpProduct) error {
	return r.db.Create(product).Error
}

// 更新商品
func (r *SpProductRepository) Update(product *sp.SpProduct) error {
	return r.db.Save(product).Error
}

// 根据ID获取商品
func (r *SpProductRepository) FindByID(id uint) (*sp.SpProduct, error) {
	var product sp.SpProduct
	err := r.db.Preload("Content").
		Preload("Properties").
		Preload("Skus").
		First(&product, id).Error
	return &product, err
}

// 根据分类ID获取商品
func (r *SpProductRepository) FindByCategoryID(categoryID uint) ([]sp.SpProduct, error) {
	var products []sp.SpProduct
	err := r.db.Where("category_id = ?", categoryID).
		Order("sort_num ASC").
		Find(&products).Error
	return products, err
}

// 获取热门商品
func (r *SpProductRepository) FindHotProducts(limit int) ([]sp.SpProduct, error) {
	var products []sp.SpProduct
	err := r.db.Where("hot = 1 AND state = 1").
		Order("sort_num ASC").
		Limit(limit).
		Find(&products).Error
	return products, err
}

// 分页获取商品
func (r *SpProductRepository) ListWithPagination(page, pageSize int) ([]sp.SpProduct, int64, error) {
	var products []sp.SpProduct
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&sp.SpProduct{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).
		Limit(pageSize).
		Order("sort_num ASC").
		Find(&products).Error

	return products, total, err
}

// 更新商品库存
func (r *SpProductRepository) UpdateStock(id uint, stock int) error {
	return r.db.Model(&sp.SpProduct{}).
		Where("id = ?", id).
		Update("stock", stock).Error
}

// 更新商品状态
func (r *SpProductRepository) UpdateState(id uint, state uint8) error {
	return r.db.Model(&sp.SpProduct{}).
		Where("id = ?", id).
		Update("state", state).Error
}

// 增加销量
func (r *SpProductRepository) IncrementSoldNum(id uint, num uint16) error {
	return r.db.Model(&sp.SpProduct{}).
		Where("id = ?", id).
		Update("sold_num", gorm.Expr("sold_num + ?", num)).Error
}