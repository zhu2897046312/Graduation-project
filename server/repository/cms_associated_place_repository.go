package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsAssociatedPlaceRepository struct {
	*BaseRepository
}

func NewCmsAssociatedPlaceRepository(db *gorm.DB) *CmsAssociatedPlaceRepository {
	return &CmsAssociatedPlaceRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建地点
func (r *CmsAssociatedPlaceRepository) Create(place *cms.CmsAssociatedPlace) error {
	return r.db.Create(place).Error
}

// 更新地点
func (r *CmsAssociatedPlaceRepository) Update(place *cms.CmsAssociatedPlace) error {
	return r.db.Save(place).Error
}

// 根据ID获取地点
func (r *CmsAssociatedPlaceRepository) FindByID(id int64) (*cms.CmsAssociatedPlace, error) {
	var place cms.CmsAssociatedPlace
	err := r.db.First(&place, id).Error
	return &place, err
}

// 分页获取地点列表
func (r *CmsAssociatedPlaceRepository) ListWithPagination(page, pageSize int) ([]cms.CmsAssociatedPlace, int64, error) {
	var places []cms.CmsAssociatedPlace
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&cms.CmsAssociatedPlace{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).
		Limit(pageSize).
		Order("created_time DESC").
		Find(&places).Error

	return places, total, err
}

// 根据拼音首字母搜索地点
func (r *CmsAssociatedPlaceRepository) SearchByInitialPinyin(initial string) ([]cms.CmsAssociatedPlace, error) {
	var places []cms.CmsAssociatedPlace
	err := r.db.Where("initial_pinyin LIKE ?", initial+"%").
		Order("created_time DESC").
		Find(&places).Error
	return places, err
}