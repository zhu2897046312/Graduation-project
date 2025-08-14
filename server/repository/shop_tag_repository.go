package repository

import (
	"gorm.io/gorm"
	"server/models/shop"
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

// 更新商品标签
func (r *ShopTagRepository) Update(tag *shop.ShopTag) error {
	return r.db.Save(tag).Error
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
func (r *ShopTagRepository) ListWithPagination(page, pageSize int) ([]shop.ShopTag, int64, error) {
	var tags []shop.ShopTag
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&shop.ShopTag{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).
		Limit(pageSize).
		Order("sort_num ASC").
		Find(&tags).Error

	return tags, total, err
}