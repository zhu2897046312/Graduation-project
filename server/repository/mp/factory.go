package mp

import (
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	db                  *gorm.DB
	mp_order            *MpOrderRepository
	mp_pay_config       *MpPayConfigRepository
	mp_product          *MpProductRepository
	mp_reset_pwd        *MpResetPwdTokensRepository
	mp_user             *MpUserRepository
	mp_user_token       *MpUserTokenRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) GetDB() *gorm.DB {
	return f.db
}

func (f *RepositoryFactory) GetMpOrderRepository() *MpOrderRepository {
	if f.mp_order == nil {
		f.mp_order = NewMpOrderRepository(f.db)
	}
	return f.mp_order
}

func (f *RepositoryFactory) GetMpPayConfigRepository() *MpPayConfigRepository {
	if f.mp_pay_config == nil {
		f.mp_pay_config = NewMpPayConfigRepository(f.db)
	}
	return f.mp_pay_config
}

func (f *RepositoryFactory) GetMpProductRepository() *MpProductRepository {
	if f.mp_product == nil {
		f.mp_product = NewMpProductRepository(f.db)
	}
	return f.mp_product
}

func (f *RepositoryFactory) GetMpResetPwdTokensRepository() *MpResetPwdTokensRepository {
	if f.mp_reset_pwd == nil {
		f.mp_reset_pwd = NewMpResetPwdTokensRepository(f.db)
	}
	return f.mp_reset_pwd
}

func (f *RepositoryFactory) GetMpUserRepository() *MpUserRepository {
	if f.mp_user == nil {
		f.mp_user = NewMpUserRepository(f.db)
	}
	return f.mp_user
}

func (f *RepositoryFactory) GetMpUserTokenRepository() *MpUserTokenRepository {
	if f.mp_user_token == nil {
		f.mp_user_token = NewMpUserTokenRepository(f.db)
	}
	return f.mp_user_token
}