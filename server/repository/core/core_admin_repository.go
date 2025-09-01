package core

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/repository/base"
)

type CoreAdminRepository struct {
	*base.BaseRepository
}

func NewCoreAdminRepository(DB *gorm.DB) *CoreAdminRepository {
	return &CoreAdminRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建管理员
func (r *CoreAdminRepository) Create(admin *core.CoreAdmin) error {
	return r.DB.Create(admin).Error
}

// 更新管理员
func (r *CoreAdminRepository) Update(admin *core.CoreAdmin) error {
	return r.DB.Save(admin).Error
}

// 根据账号获取管理员
func (r *CoreAdminRepository) FindByAccount(account string) (*core.CoreAdmin, error) {
	var admin core.CoreAdmin
	err := r.DB.Where("account = ?", account).First(&admin).Error
	return &admin, err
}

// 根据ID获取管理员
func (r *CoreAdminRepository) FindByID(id int64) (*core.CoreAdmin, error) {
	var admin core.CoreAdmin
	err := r.DB.First(&admin, id).Error
	return &admin, err
}

// 根据部门ID获取管理员列表
func (r *CoreAdminRepository) FindByDeptID(deptID int64) ([]core.CoreAdmin, error) {
	var admins []core.CoreAdmin
	err := r.DB.Where("dept_id = ?", deptID).Find(&admins).Error
	return admins, err
}

// 更新管理员状态
func (r *CoreAdminRepository) UpdateStatus(id int64, status int8) error {
	return r.DB.Model(&core.CoreAdmin{}).
		Where("id = ?", id).
		Update("admin_status", status).Error
}

// 更新管理员密码
func (r *CoreAdminRepository) UpdatePassword(id int64, newPwd string) error {
	return r.DB.Model(&core.CoreAdmin{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"pwd":       newPwd,
			"last_pwd":  gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
}