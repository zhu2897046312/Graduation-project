package shop

import (
	"gorm.io/gorm"
	"server/models/shop"
	"server/repository/base"
	"time"
)

type ShopTagMateRepository struct {
	*base.BaseRepository
}

func NewShopTagMateRepository(DB *gorm.DB) *ShopTagMateRepository {
	return &ShopTagMateRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建标签元数据
func (r *ShopTagMateRepository) Create(mate *shop.ShopTagMate) error {
	return r.DB.Create(mate).Error
}

// 更新标签元数据
func (r *ShopTagMateRepository) Update(mate *shop.ShopTagMate) error {
	return r.DB.Updates(mate).Error
}

// 根据ID获取元数据
func (r *ShopTagMateRepository) FindByID(id int) (*shop.ShopTagMate, error) {
	var mate shop.ShopTagMate
	err := r.DB.First(&mate, id).Error
	return &mate, err
}

// 更新SEO信息
func (r *ShopTagMateRepository) UpdateSEO(id int, title, keyword, description string) error {
	return r.DB.Model(&shop.ShopTagMate{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"seo_title":       title,
			"seo_keyword":     keyword,
			"seo_description": description,
		}).Error
}

// 更新标签内容
func (r *ShopTagMateRepository) UpdateContent(id int, content string) error {
	return r.DB.Model(&shop.ShopTagMate{}).
		Where("id = ?", id).
		Update("content", content).Error
}

func (r *ShopTagMateRepository) DeleteByID(id int) error{
	//return r.DB.Where("id = ?", id).Delete(&shop.ShopTagMate{}).Error
	result := r.DB.Model(&shop.ShopTagMate{}).
		Where("id = ?", id).
		Update("deleted_time", time.Now())

	return result.Error
	
}