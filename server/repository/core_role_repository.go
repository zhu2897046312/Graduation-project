package repository

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/models/common"
)

type CoreRoleRepository struct {
	*BaseRepository
}

func NewCoreRoleRepository(db *gorm.DB) *CoreRoleRepository {
	return &CoreRoleRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建角色
func (r *CoreRoleRepository) Create(role *core.CoreRole) error {
	return r.db.Create(role).Error
}

// 更新角色
func (r *CoreRoleRepository) Update(role *core.CoreRole) error {
	return r.db.Updates(role).Error
}

// 根据ID获取角色
func (r *CoreRoleRepository) FindByID(id common.MyID) (*core.CoreRole, error) {
	var role core.CoreRole
	err := r.db.First(&role, id).Error
	return &role, err
}

// 获取所有角色
func (r *CoreRoleRepository) FindAll() ([]core.CoreRole, error) {
	var roles []core.CoreRole
	err := r.db.Find(&roles).Error
	return roles, err
}

// 更新角色状态
func (r *CoreRoleRepository) UpdateStatus(id common.MyID, status int8) error {
	return r.db.Model(&core.CoreRole{}).
		Where("id = ?", id).
		Update("role_status", status).Error
}

// 更新角色权限
func (r *CoreRoleRepository) UpdatePermissions(id common.MyID, permissions []byte) error {
	return r.db.Model(&core.CoreRole{}).
		Where("id = ?", id).
		Update("permission", permissions).Error
}

// 分页获取商品（带过滤选项）
func (r *CoreRoleRepository) List(page int, pageSize int) ([]core.CoreRole, int64, error) {
	var products []core.CoreRole
	var total int64

	// 设置默认值
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
	pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 构建查询
	query := r.db.Model(&core.CoreRole{}).Where("deleted_time IS NULL")

	// // 应用过滤条件
	// if params.CategoryID != 0 {
	// 	query = query.Where("category_id = ?", params.CategoryID)
	// }

	// if params.State != 0 {
	// 	query = query.Where("state = ?", params.State)
	// }

	// if params.Title != "" {
	// 	query = query.Where("title LIKE ?", "%"+params.Title+"%")
	// }

	// if params.Hot != nil {
	// 	query = query.Where("hot = ?", *params.Hot)
	// }

	// // 设置排序
	// orderClause := params.SortBy + " " + params.SortOrder
	// query = query.Order(orderClause)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := query.Offset(offset).
		Limit(pageSize).
		Find(&products).Error

	return products, total, err
}

func (r *CoreRoleRepository) FindByAdminID(id common.MyID) ([]core.CoreRole, error) {
	var products []core.CoreRole
	// 构建查询
	query := r.db.Model(&core.CoreRole{}).Where("deleted_time IS NULL")

	if id != 0 {
		query = query.Where("admin_id = ?", id)
	}
	// 获取分页数据
	err := query.Find(&products).Error

	return products, err
}

func (r *CoreRoleRepository) Delete(id common.MyID) error {
	return r.db.Model(&core.CoreRole{}).
		Where("id = ?", id).
		Update("deleted_time", gorm.Expr("NOW()")).Error
}