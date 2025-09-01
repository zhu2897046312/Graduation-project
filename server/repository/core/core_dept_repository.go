package core

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/repository/base"
)

type CoreDeptRepository struct {
	*base.BaseRepository
}

func NewCoreDeptRepository(DB *gorm.DB) *CoreDeptRepository {
	return &CoreDeptRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建部门
func (r *CoreDeptRepository) Create(dept *core.CoreDept) error {
	return r.DB.Create(dept).Error
}

// 更新部门
func (r *CoreDeptRepository) Update(dept *core.CoreDept) error {
	return r.DB.Save(dept).Error
}

// 根据ID获取部门
func (r *CoreDeptRepository) FindByID(id int64) (*core.CoreDept, error) {
	var dept core.CoreDept
	err := r.DB.First(&dept, id).Error
	return &dept, err
}

// // 获取所有部门树形结构
// func (r *CoreDeptRepository) FindTree() ([]*core.CoreDept, error) {
// 	var depts []*core.CoreDept
// 	err := r.DB.Order("sort_num ASC").Find(&depts).Error
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
	err := r.DB.Where("pid = ?", pid).Order("sort_num ASC").Find(&depts).Error
	return depts, err
}

// 删除部门
func (r *CoreDeptRepository) Delete(id int64) error {
	return r.DB.Delete(&core.CoreDept{}, id).Error
}