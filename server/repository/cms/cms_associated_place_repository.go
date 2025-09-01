package cms

import (
	"gorm.io/gorm"
	"server/models/cms"
	"server/repository/base"
)

type CmsAssociatedPlaceRepository struct {
	*base.BaseRepository
}

func NewCmsAssociatedPlaceRepository(DB *gorm.DB) *CmsAssociatedPlaceRepository {
	return &CmsAssociatedPlaceRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建地点
func (r *CmsAssociatedPlaceRepository) Create(place *cms.CmsAssociatedPlace) error {
	return r.DB.Create(place).Error
}

// 更新地点
func (r *CmsAssociatedPlaceRepository) Update(place *cms.CmsAssociatedPlace) error {
	return r.DB.Save(place).Error
}

// 根据ID获取地点
func (r *CmsAssociatedPlaceRepository) FindByID(id int64) (*cms.CmsAssociatedPlace, error) {
	var place cms.CmsAssociatedPlace
	err := r.DB.First(&place, id).Error
	return &place, err
}

// 分页获取地点列表
func (r *CmsAssociatedPlaceRepository) ListWithPagination(page, pageSize int) ([]cms.CmsAssociatedPlace, int64, error) {
	var places []cms.CmsAssociatedPlace
	var total int64

	offset := (page - 1) * pageSize

	if err := r.DB.Model(&cms.CmsAssociatedPlace{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.DB.Offset(offset).
		Limit(pageSize).
		Order("created_time DESC").
		Find(&places).Error

	return places, total, err
}

// 根据拼音首字母搜索地点
func (r *CmsAssociatedPlaceRepository) SearchByInitialPinyin(initial string) ([]cms.CmsAssociatedPlace, error) {
	var places []cms.CmsAssociatedPlace
	err := r.DB.Where("initial_pinyin LIKE ?", initial+"%").
		Order("created_time DESC").
		Find(&places).Error
	return places, err
}