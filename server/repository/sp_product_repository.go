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
// 创建商品并返回完整的商品信息
func (r *SpProductRepository) Create(product *sp.SpProduct) (*sp.SpProduct, error) {
    // 执行创建操作
    if err := r.db.Create(product).Error; err != nil {
        return nil, err
    }
    // 返回创建后的完整对象
    return product, nil
}


// 更新商品
func (r *SpProductRepository) Update(product *sp.SpProduct) error {
	return r.db.Updates(product).Error
}

// 根据ID获取商品
func (r *SpProductRepository) FindByID(id uint) (*sp.SpProduct, error) {
	var product sp.SpProduct
	err := r.db.First(&product, id).Error
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
// 分页获取商品（带过滤选项）
func (r *SpProductRepository) ListWithPagination(params sp.ProductQueryParams) ([]sp.SpProduct, int64, error) {
	var products []sp.SpProduct
	var total int64

	// 设置默认值
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 10
	}
	if params.SortBy == "" {
		params.SortBy = "sort_num"
	}
	if params.SortOrder == "" {
		params.SortOrder = "ASC"
	}

	offset := (params.Page - 1) * params.PageSize

	// 构建查询
	query := r.db.Model(&sp.SpProduct{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	if params.CategoryID != 0 {
		query = query.Where("category_id = ?", params.CategoryID)
	}

	if params.State != 0 {
		query = query.Where("state = ?", params.State)
	}

	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}

	if params.Hot != nil {
		query = query.Where("hot = ?", *params.Hot)
	}

	// 设置排序
	orderClause := params.SortBy + " " + params.SortOrder
	query = query.Order(orderClause)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := query.Offset(offset).
		Limit(params.PageSize).
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

// 软删除商品（GORM 自动处理）
func (r *SpProductRepository) SoftDelete(id uint) error {
    return r.db.Delete(&sp.SpProduct{}, id).Error
}