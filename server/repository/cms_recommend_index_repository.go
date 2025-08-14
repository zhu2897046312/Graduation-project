package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsRecommendIndexRepository struct {
	*BaseRepository
}

func NewCmsRecommendIndexRepository(db *gorm.DB) *CmsRecommendIndexRepository {
	return &CmsRecommendIndexRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建推荐索引
func (r *CmsRecommendIndexRepository) Create(index *cms.CmsRecommendIndex) error {
	return r.db.Create(index).Error
}

// 更新推荐索引
func (r *CmsRecommendIndexRepository) Update(index *cms.CmsRecommendIndex) error {
	return r.db.Save(index).Error
}

// 根据推荐ID获取索引项
func (r *CmsRecommendIndexRepository) FindByRecommendID(recommendID int) ([]cms.CmsRecommendIndex, error) {
	var indices []cms.CmsRecommendIndex
	err := r.db.Where("recommend_id = ?", recommendID).
		Order("sort_num ASC").
		Find(&indices).Error
	return indices, err
}

// 根据状态获取推荐索引
func (r *CmsRecommendIndexRepository) FindByState(state int8) ([]cms.CmsRecommendIndex, error) {
	var indices []cms.CmsRecommendIndex
	err := r.db.Where("state = ?", state).
		Order("sort_num ASC").
		Find(&indices).Error
	return indices, err
}

// 删除推荐索引
func (r *CmsRecommendIndexRepository) Delete(id int) error {
	return r.db.Delete(&cms.CmsRecommendIndex{}, id).Error
}