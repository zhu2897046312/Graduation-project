package repository

import (
	"gorm.io/gorm"
	"server/models/shop"
)

type ShopTagMateRepository struct {
	*BaseRepository
}

func NewShopTagMateRepository(db *gorm.DB) *ShopTagMateRepository {
	return &ShopTagMateRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建标签元数据
func (r *ShopTagMateRepository) Create(mate *shop.ShopTagMate) error {
	return r.db.Create(mate).Error
}

// 更新标签元数据
func (r *ShopTagMateRepository) Update(mate *shop.ShopTagMate) error {
	return r.db.Save(mate).Error
}

// 根据ID获取元数据
func (r *ShopTagMateRepository) FindByID(id int) (*shop.ShopTagMate, error) {
	var mate shop.ShopTagMate
	err := r.db.First(&mate, id).Error
	return &mate, err
}

// 更新SEO信息
func (r *ShopTagMateRepository) UpdateSEO(id int, title, keyword, description string) error {
	return r.db.Model(&shop.ShopTagMate{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"seo_title":       title,
			"seo_keyword":     keyword,
			"seo_description": description,
		}).Error
}

// 更新标签内容
func (r *ShopTagMateRepository) UpdateContent(id int, content string) error {
	return r.db.Model(&shop.ShopTagMate{}).
		Where("id = ?", id).
		Update("content", content).Error
}