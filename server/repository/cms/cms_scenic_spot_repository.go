package cms

import (
	"gorm.io/gorm"
	"server/models/cms"
	"server/repository/base"
)

type CmsScenicSpotRepository struct {
	*base.BaseRepository
}

func NewCmsScenicSpotRepository(DB *gorm.DB) *CmsScenicSpotRepository {
	return &CmsScenicSpotRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建景点
func (r *CmsScenicSpotRepository) Create(spot *cms.CmsScenicSpot) error {
	return r.DB.Create(spot).Error
}

// 更新景点
func (r *CmsScenicSpotRepository) Update(spot *cms.CmsScenicSpot) error {
	return r.DB.Save(spot).Error
}

// 根据ID获取景点
func (r *CmsScenicSpotRepository) FindByID(id int64) (*cms.CmsScenicSpot, error) {
	var spot cms.CmsScenicSpot
	err := r.DB.First(&spot, id).Error
	return &spot, err
}

// 根据地点ID获取景点
func (r *CmsScenicSpotRepository) FindByPlaceID(placeID int64) ([]cms.CmsScenicSpot, error) {
	var spots []cms.CmsScenicSpot
	err := r.DB.Where("associated_place_id = ?", placeID).
		Order("created_time DESC").
		Find(&spots).Error
	return spots, err
}

// 增加景点阅读量
func (r *CmsScenicSpotRepository) IncrementReadNum(id int64) error {
	return r.DB.Model(&cms.CmsScenicSpot{}).
		Where("id = ?", id).
		Update("read_num", gorm.Expr("read_num + ?", 1)).Error
}

// 分页获取景点列表
func (r *CmsScenicSpotRepository) ListWithPagination(page, pageSize int) ([]cms.CmsScenicSpot, int64, error) {
	var spots []cms.CmsScenicSpot
	var total int64

	offset := (page - 1) * pageSize

	if err := r.DB.Model(&cms.CmsScenicSpot{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.DB.Offset(offset).
		Limit(pageSize).
		Order("created_time DESC").
		Find(&spots).Error

	return spots, total, err
}