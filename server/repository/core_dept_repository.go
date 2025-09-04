package repository

import (
	"gorm.io/gorm"
	"server/models/core"
)

type CoreDeptRepository struct {
	*BaseRepository
}

func NewCoreDeptRepository(db *gorm.DB) *CoreDeptRepository {
	return &CoreDeptRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建部门
func (r *CoreDeptRepository) Create(dept *core.CoreDept) error {
	return r.db.Create(dept).Error
}

// 更新部门
func (r *CoreDeptRepository) Update(dept *core.CoreDept) error {
	return r.db.Save(dept).Error
}

// 根据ID获取部门
func (r *CoreDeptRepository) FindByID(id int64) (*core.CoreDept, error) {
	var dept core.CoreDept
	err := r.db.First(&dept, id).Error
	return &dept, err
}

// // 获取所有部门树形结构
// func (r *CoreDeptRepository) FindTree() ([]*core.CoreDept, error) {
// 	var depts []*core.CoreDept
// 	err := r.db.Order("sort_num ASC").Find(&depts).Error
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	// 构建树形结构
// 	deptMap := make(map[int64]*core.CoreDept)
// 	for i := range depts {
// 		deptMap[depts[i].ID] = depts[i]
// 	}
	
// 	var tree []*core.CoreDept
// 	for i := range depts {
// 		if depts[i].Pid == 0 {
// 			tree = append(tree, depts[i])
// 		} else {
// 			if parent, ok := deptMap[depts[i].Pid]; ok {
// 				parent.Children = append(parent.Children, depts[i])
// 			}
// 		}
// 	}
	
// 	return tree, nil
// }

// 根据父ID获取子部门
func (r *CoreDeptRepository) FindByPid(pid int64) ([]*core.CoreDept, error) {
	var depts []*core.CoreDept
	err := r.db.Where("pid = ?", pid).Order("sort_num ASC").Find(&depts).Error
	return depts, err
}

// 删除部门
func (r *CoreDeptRepository) Delete(id int64) error {
	return r.db.Delete(&core.CoreDept{}, id).Error
}

func (r *CoreDeptRepository) List(pid int64) ([]core.CoreDept,int64, error) {
	var coreDepts []core.CoreDept
	var total int64

	// if params.Page < 1 {
	// 	params.Page = 1
	// }
	// if params.PageSize < 1 || params.PageSize > 100 {
	// 	params.PageSize = 10
	// }
	// offset := (params.Page - 1) * params.PageSize

	// 构建查询
	query := r.db.Model(&core.CoreDept{}).Where("deleted_time IS NULL")

	// if params.Nickname != "" {
	// 	query = query.Where("nickname like ?", "%"+params.Nickname+"%")
	// }
	// if params.Account != "" {
	// 	query = query.Where("account like ?", "%"+params.Account+"%")
	// }
	// if params.AdminStatus != 0 {
	// 	query = query.Where("admin_status = ?", params.AdminStatus)
	// }	

	// 获取符合条件的总记录数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取匹配的退款记录列表.Offset(offset).Limit(params.PageSize)
	err := query.Find(&coreDepts).Error

	return coreDepts, total, err
}

func (r *CoreDeptRepository) FindAll() ([]*core.CoreDept,int64, error) {
	var coreDepts []*core.CoreDept
	var total int64

	query := r.db.Model(&core.CoreDept{}).Where("deleted_time IS NULL")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Find(&coreDepts).Error
	return coreDepts,total, err
}