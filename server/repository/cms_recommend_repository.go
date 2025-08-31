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

// 分页获取商品（带过滤选项）
func (r *CmsRecommendRepository) ListWithPagination(params cms.RecommendQueryParams) ([]cms.CmsRecommend, int64, error) {
	var recommends []cms.CmsRecommend
	var total int64

	offset := (params.Page - 1) * params.PageSize

	// 构建查询
	query := r.db.Model(&cms.CmsRecommend{})

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	 // 如果 pageSize <= 0，返回全部数据
	 if params.PageSize <= 0 {
        err := query.Order("created_time DESC").Find(&recommends).Error
        return recommends, total, err
    }

	// 获取分页数据
	err := query.Offset(offset).
		Limit(params.PageSize).
		Find(&recommends).Error

	return recommends, total, err
}