package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpProductContentRepository struct {
	*base.BaseRepository
}

func NewSpProductContentRepository(DB *gorm.DB) *SpProductContentRepository {
	return &SpProductContentRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建商品内容
func (r *SpProductContentRepository) Create(content *sp.SpProductContent) error {
	return r.DB.Create(content).Error
}

// 更新商品内容
func (r *SpProductContentRepository) Update(content *sp.SpProductContent) error {
	return r.DB.Where("product_id = ?", content.ProductID).Updates(content).Error
}

// 根据商品ID获取内容
func (r *SpProductContentRepository) FindByProductID(productID uint) (*sp.SpProductContent, error) {
	var content sp.SpProductContent
	err := r.DB.Where("product_id = ?", productID).First(&content).Error
	return &content, err
}

// 更新SEO信息
func (r *SpProductContentRepository) UpdateSEO(productID uint, title, keyword, description string) error {
	return r.DB.Model(&sp.SpProductContent{}).
		Where("product_id = ?", productID).
		Updates(map[string]interface{}{
			"seo_title":       title,
			"seo_keyword":     keyword,
			"seo_description": description,
		}).Error
}

// 更新商品内容
func (r *SpProductContentRepository) UpdateContent(productID uint, content string) error {
	return r.DB.Model(&sp.SpProductContent{}).
		Where("product_id = ?", productID).
		Update("content", content).Error
}