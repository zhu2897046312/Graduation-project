package repository

import (
	"server/models/cms"
	"server/models/common"
	"time"

	"gorm.io/gorm"
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
	return r.db.Updates(index).Error
}

func (r *CmsRecommendIndexRepository) FindByID(id common.MyID) (*cms.CmsRecommendIndex, error) {
	var index cms.CmsRecommendIndex
	err := r.db.First(&index, id).Error
	return &index, err
}

func (r *CmsRecommendIndexRepository) Delete(id common.MyID) error {
	result := r.db.Model(&cms.CmsRecommendIndex{}).
		Where("id = ?", id).
		Update("deleted_time", time.Now())

	return result.Error
}

// 根据推荐ID获取索引项
func (r *CmsRecommendIndexRepository) FindByRecommendID(recommendID common.MyID) ([]cms.CmsRecommendIndex, error) {
	var indices []cms.CmsRecommendIndex
	query := r.db.Model(&cms.CmsRecommendIndex{}).Where("deleted_time IS NULL")
	err := query.Where("recommend_id = ?", recommendID).
		Order("sort_num ASC").
		Find(&indices).Error
	return indices, err
}

// 根据状态获取推荐索引
func (r *CmsRecommendIndexRepository) FindByState(state int8) ([]cms.CmsRecommendIndex, error) {
	var indices []cms.CmsRecommendIndex
	query := r.db.Model(&cms.CmsRecommendIndex{}).Where("deleted_time IS NULL")
	err := query.Where("state = ?", state).
		Order("sort_num ASC").
		Find(&indices).Error
	return indices, err
}

// 分页获取商品（带过滤选项）
func (r *CmsRecommendIndexRepository) ListWithPagination(params cms.RecommendIndexQueryParams) ([]cms.CmsRecommendIndex, int64, error) {
	var recommends []cms.CmsRecommendIndex
	var total int64

	offset := (params.Page - 1) * params.PageSize

	// 构建查询
	query := r.db.Model(&cms.CmsRecommendIndex{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	if params.RecommendID != 0 {
		query = query.Where("recommend_id = ?", params.RecommendID)
	}

	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}

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
