package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsRecommendRepository struct {
	*BaseRepository
}

func NewCmsRecommendRepository(db *gorm.DB) *CmsRecommendRepository {
	return &CmsRecommendRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建推荐内容
func (r *CmsRecommendRepository) Create(recommend *cms.CmsRecommend) error {
	return r.db.Create(recommend).Error
}

// 更新推荐内容
func (r *CmsRecommendRepository) Update(recommend *cms.CmsRecommend) error {
	return r.db.Save(recommend).Error
}

// 获取推荐列表
func (r *CmsRecommendRepository) FindAll() ([]cms.CmsRecommend, error) {
	var recommends []cms.CmsRecommend
	err := r.db.Where("state = 1").Order("created_time DESC").Find(&recommends).Error
	return recommends, err
}

// 根据状态获取推荐内容
func (r *CmsRecommendRepository) FindByState(state int8) ([]cms.CmsRecommend, error) {
	var recommends []cms.CmsRecommend
	err := r.db.Where("state = ?", state).Order("created_time DESC").Find(&recommends).Error
	return recommends, err
}