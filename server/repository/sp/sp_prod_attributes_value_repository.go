package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpProdAttributesValueRepository struct {
	*base.BaseRepository
}

func NewSpProdAttributesValueRepository(DB *gorm.DB) *SpProdAttributesValueRepository {
	return &SpProdAttributesValueRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建属性值
func (r *SpProdAttributesValueRepository) Create(value *sp.SpProdAttributesValue) error {
	return r.DB.Create(value).Error
}

// 更新属性值
func (r *SpProdAttributesValueRepository) Update(value *sp.SpProdAttributesValue) error {
	return r.DB.Save(value).Error
}

// 根据属性ID获取值列表
func (r *SpProdAttributesValueRepository) FindByAttributeID(attrID uint) ([]sp.SpProdAttributesValue, error) {
	var values []sp.SpProdAttributesValue
	err := r.DB.Where("prod_attributes_id = ?", attrID).
		Order("sort_num ASC").
		Find(&values).Error
	return values, err
}

// 根据ID获取属性值
func (r *SpProdAttributesValueRepository) FindByID(id uint) (*sp.SpProdAttributesValue, error) {
	var value sp.SpProdAttributesValue
	err := r.DB.First(&value, id).Error
	return &value, err
}

func (r *SpProdAttributesValueRepository) FindByIDList(attrID uint) ([]sp.SpProdAttributesValue, error) {
	var values []sp.SpProdAttributesValue
	err := r.DB.Where("id = ?", attrID).
		Order("sort_num ASC").
		Find(&values).Error
	return values, err
}

// 批量创建属性值
func (r *SpProdAttributesValueRepository) BatchCreate(values []sp.SpProdAttributesValue) error {
	return r.DB.Create(&values).Error
}

// 删除属性的所有值
func (r *SpProdAttributesValueRepository) DeleteByAttributeID(attrID uint) error {
	return r.DB.Where("prod_attributes_id = ?", attrID).
		Delete(&sp.SpProdAttributesValue{}).Error
}

func (r *SpProdAttributesValueRepository) ListWithPagination(params sp.SpProdAttributesQueryParams) ([]sp.SpProdAttributesValue, int64, error) {
	var products []sp.SpProdAttributesValue
	var total int64

	// 设置默认值
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 10
	}

	offset := (params.Page - 1) * params.PageSize

	// 构建查询
	query := r.DB.Model(&sp.SpProdAttributesValue{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	if params.ProdAttributesID != 0 {
		query = query.Where("prod_attributes_id = ?", params.ProdAttributesID)
	}
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := query.Offset(offset).
		Limit(params.PageSize).
		Find(&products).Error

	return products, total, err
}

func (r *SpProdAttributesValueRepository) List(params sp.SpProdAttributesQueryParams) ([]sp.SpProdAttributesValue,int64, error) {
	var products []sp.SpProdAttributesValue
	var total int64
	// 构建查询
	query := r.DB.Model(&sp.SpProdAttributesValue{})

	// 应用过滤条件
	if params.ProdAttributesID != 0 {
		query = query.Where("prod_attributes_id = ?", params.ProdAttributesID)
	}
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 获取数据
	err := query.Find(&products).Error
	return products,total, err
}
