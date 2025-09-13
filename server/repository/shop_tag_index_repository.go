package repository

import (
	"gorm.io/gorm"
	"server/models/shop"
	"server/models/common"
)

type ShopTagIndexRepository struct {
	*BaseRepository
}

func NewShopTagIndexRepository(db *gorm.DB) *ShopTagIndexRepository {
	return &ShopTagIndexRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建标签关联
func (r *ShopTagIndexRepository) Create(index *shop.ShopTagIndex) error {
	return r.db.Create(index).Error
}

// 删除标签关联
func (r *ShopTagIndexRepository) Delete(productID, tagID common.MyID) error {
	return r.db.Where("product_id = ? AND tag_id = ?", productID, tagID).
		Delete(&shop.ShopTagIndex{}).Error
}

// 根据产品ID获取标签关联
func (r *ShopTagIndexRepository) FindByProductID(productID common.MyID) ([]shop.ShopTagIndex, error) {
	var indices []shop.ShopTagIndex
	err := r.db.Where("product_id = ?", productID).
		Order("sort_num ASC").
		Find(&indices).Error
	return indices, err
}

// 根据标签ID获取产品关联
func (r *ShopTagIndexRepository) FindByTagID(tagID common.MyID) ([]shop.ShopTagIndex, error) {
	var indices []shop.ShopTagIndex
	err := r.db.Where("tag_id = ?", tagID).
		Order("sort_num ASC").
		Find(&indices).Error
	return indices, err
}

// 更新标签关联排序
func (r *ShopTagIndexRepository) UpdateSortNum(id, sortNum int) error {
	return r.db.Model(&shop.ShopTagIndex{}).
		Where("id = ?", id).
		Update("sort_num", sortNum).Error
}

// 删除产品的所有标签关联
func (r *ShopTagIndexRepository) DeleteByProductID(productID common.MyID) error {
	return r.db.Where("product_id = ?", productID).
		Delete(&shop.ShopTagIndex{}).Error
}