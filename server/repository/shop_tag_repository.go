package repository

import (
	"gorm.io/gorm"
	"server/models/shop"
	"time"
)

type ShopTagRepository struct {
	*BaseRepository
}

func NewShopTagRepository(db *gorm.DB) *ShopTagRepository {
	return &ShopTagRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建商品标签
func (r *ShopTagRepository) Create(tag *shop.ShopTag) error {
	return r.db.Create(tag).Error
}

// 更新商品标签shop
func (r *ShopTagRepository) Update(tag *shop.ShopTag) error {
	return r.db.Updates(tag).Error
}

// 根据ID获取标签
func (r *ShopTagRepository) FindByID(id int) (*shop.ShopTag, error) {
	var tag shop.ShopTag
	err := r.db.First(&tag, id).Error
	return &tag, err
}

// 根据状态获取标签列表
func (r *ShopTagRepository) FindByState(state int8) ([]shop.ShopTag, error) {
	var tags []shop.ShopTag
	err := r.db.Where("state = ?", state).
		Order("sort_num ASC").
		Find(&tags).Error
	return tags, err
}

// 根据匹配词获取标签
func (r *ShopTagRepository) FindByMatchWord(matchWord string) ([]shop.ShopTag, error) {
	var tags []shop.ShopTag
	err := r.db.Where("match_word LIKE ?", "%"+matchWord+"%").
		Order("sort_num ASC").
		Find(&tags).Error
	return tags, err
}

// 增加标签阅读量
func (r *ShopTagRepository) IncrementReadNum(id int) error {
	return r.db.Model(&shop.ShopTag{}).
		Where("id = ?", id).
		Update("read_num", gorm.Expr("read_num + ?", 1)).Error
}

// 获取所有标签（分页）
func (r *ShopTagRepository) ListWithPagination(params shop.TagQueryParams) ([]shop.ShopTag, int64, error) {
	var tags []shop.ShopTag
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
	query := r.db.Model(&shop.ShopTag{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	if params.State != 0 {
		query = query.Where("state = ?", params.State)
	}

	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
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
		Find(&tags).Error

	return tags, total, err
}

func (r *ShopTagRepository) DeleteByID(id int) error {
	result := r.db.Model(&shop.ShopTag{}).
		Where("id = ?", id).
		Update("deleted_time", time.Now())

	return result.Error
}