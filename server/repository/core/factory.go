package core

import (
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	db                  *gorm.DB
	core_admin          *CoreAdminRepository
	core_admin_role     *CoreAdminRoleIndexRepository
	core_config         *CoreConfigRepository
	core_dept           *CoreDeptRepository
	core_permission     *CorePermissionRepository
	core_request_log    *CoreRequestLogRepository
	core_role           *CoreRoleRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) GetDB() *gorm.DB {
	return f.db
}

func (f *RepositoryFactory) GetCoreAdminRepository() *CoreAdminRepository {
	if f.core_admin == nil {
		f.core_admin = NewCoreAdminRepository(f.db)
	}
	return f.core_admin
}

func (f *RepositoryFactory) GetCoreAdminRoleIndexRepository() *CoreAdminRoleIndexRepository {
	if f.core_admin_role == nil {
		f.core_admin_role = NewCoreAdminRoleIndexRepository(f.db)
	}
	return f.core_admin_role
}

func (f *RepositoryFactory) GetCoreConfigRepository() *CoreConfigRepository {
	if f.core_config == nil {
		f.core_config = NewCoreConfigRepository(f.db)
	}
	return f.core_config
}

func (f *RepositoryFactory) GetCoreDeptRepository() *CoreDeptRepository {
	if f.core_dept == nil {
		f.core_dept = NewCoreDeptRepository(f.db)
	}
	return f.core_dept
}

func (f *RepositoryFactory) GetCorePermissionRepository() *CorePermissionRepository {
	if f.core_permission == nil {
		f.core_permission = NewCorePermissionRepository(f.db)
	}
	return f.core_permission
}

func (f *RepositoryFactory) GetCoreRequestLogRepository() *CoreRequestLogRepository {
	if f.core_request_log == nil {
		f.core_request_log = NewCoreRequestLogRepository(f.db)
	}
	return f.core_request_log
}

func (f *RepositoryFactory) GetCoreRoleRepository() *CoreRoleRepository {
	if f.core_role == nil {
		f.core_role = NewCoreRoleRepository(f.db)
	}
	return f.core_role
}