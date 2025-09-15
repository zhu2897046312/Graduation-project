package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsScenicSpotRepository struct {
	*BaseRepository
}

func NewCmsScenicSpotRepository(db *gorm.DB) *CmsScenicSpotRepository {
	return &CmsScenicSpotRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建景点
func (r *CmsScenicSpotRepository) Create(spot *cms.CmsScenicSpot) error {
	return r.db.Create(spot).Error
}

// 更新景点
func (r *CmsScenicSpotRepository) Update(spot *cms.CmsScenicSpot) error {
	return r.db.Save(spot).Error
}

// 根据ID获取景点
func (r *CmsScenicSpotRepository) FindByID(id int64) (*cms.CmsScenicSpot, error) {
	var spot cms.CmsScenicSpot
	err := r.db.First(&spot, id).Error
	return &spot, err
}

// 根据地点ID获取景点
func (r *CmsScenicSpotRepository) FindByPlaceID(placeID int64) ([]cms.CmsScenicSpot, error) {
	var spots []cms.CmsScenicSpot
	err := r.db.Where("associated_place_id = ?", placeID).
		Order("created_time DESC").
		Find(&spots).Error
	return spots, err
}

// 增加景点阅读量
func (r *CmsScenicSpotRepository) IncrementReadNum(id int64) error {
	return r.db.Model(&cms.CmsScenicSpot{}).
		Where("id = ?", id).
		Update("read_num", gorm.Expr("read_num + ?", 1)).Error
}

// 分页获取景点列表
func (r *CmsScenicSpotRepository) ListWithPagination(page, pageSize int) ([]cms.CmsScenicSpot, int64, error) {
	var spots []cms.CmsScenicSpot
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&cms.CmsScenicSpot{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).
		Limit(pageSize).
		Order("created_time DESC").
		Find(&spots).Error

	return spots, total, err
}