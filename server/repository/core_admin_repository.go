package repository

import (
	"gorm.io/gorm"
	"server/models/core"
	
)

type CoreAdminRepository struct {
	*BaseRepository
}

func NewCoreAdminRepository(db *gorm.DB) *CoreAdminRepository {
	return &CoreAdminRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建管理员
func (r *CoreAdminRepository) Create(admin *core.CoreAdmin) error {
	return r.db.Create(admin).Error
}

// 更新管理员
func (r *CoreAdminRepository) Update(admin *core.CoreAdmin) error {
	return r.db.Updates(admin).Error
}

// 根据账号获取管理员
func (r *CoreAdminRepository) FindByAccount(account string) (*core.CoreAdmin, error) {
	var admin core.CoreAdmin
	err := r.db.Where("account = ?", account).First(&admin).Error
	return &admin, err
}

// 根据ID获取管理员
func (r *CoreAdminRepository) FindByID(id int64) (*core.CoreAdmin, error) {
	var admin core.CoreAdmin
	err := r.db.First(&admin, id).Error
	return &admin, err
}

// 根据部门ID获取管理员列表
func (r *CoreAdminRepository) FindByDeptID(deptID int64) ([]core.CoreAdmin, error) {
	var admins []core.CoreAdmin
	err := r.db.Where("dept_id = ?", deptID).Find(&admins).Error
	return admins, err
}

// 更新管理员状态
func (r *CoreAdminRepository) UpdateStatus(id int64, status int8) error {
	return r.db.Model(&core.CoreAdmin{}).
		Where("id = ?", id).
		Update("admin_status", status).Error
}

// 更新管理员密码
func (r *CoreAdminRepository) UpdatePassword(id int64, newPwd string) error {
	return r.db.Model(&core.CoreAdmin{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"pwd":       newPwd,
			"last_pwd":  gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
}

func (r *CoreAdminRepository) List(params core.CoreAdminQueryParam) ([]core.CoreAdmin,int64, error) {
	var coreAdmins []core.CoreAdmin
	var total int64

	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 10
	}
	offset := (params.Page - 1) * params.PageSize

	// 构建查询
	query := r.db.Model(&core.CoreAdmin{}).Where("deleted_time IS NULL")

	if params.Nickname != "" {
		query = query.Where("nickname like ?", "%"+params.Nickname+"%")
	}
	if params.Account != "" {
		query = query.Where("account like ?", "%"+params.Account+"%")
	}
	if params.AdminStatus != 0 {
		query = query.Where("admin_status = ?", params.AdminStatus)
	}	

	// 获取符合条件的总记录数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取匹配的退款记录列表
	err := query.Offset(offset).Limit(params.PageSize).Find(&coreAdmins).Error

	return coreAdmins, total, err
}