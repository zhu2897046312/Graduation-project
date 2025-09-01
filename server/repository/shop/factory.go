package shop

import (
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	db                  *gorm.DB
	shop_tag            *ShopTagRepository
	shop_tag_index      *ShopTagIndexRepository
	shop_tag_mate       *ShopTagMateRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) GetDB() *gorm.DB {
	return f.db
}

func (f *RepositoryFactory) GetShopTagRepository() *ShopTagRepository {
	if f.shop_tag == nil {
		f.shop_tag = NewShopTagRepository(f.db)
	}
	return f.shop_tag
}

func (f *RepositoryFactory) GetShopTagIndexRepository() *ShopTagIndexRepository {
	if f.shop_tag_index == nil {
		f.shop_tag_index = NewShopTagIndexRepository(f.db)
	}
	return f.shop_tag_index
}

func (f *RepositoryFactory) GetShopTagMateRepository() *ShopTagMateRepository {
	if f.shop_tag_mate == nil {
		f.shop_tag_mate = NewShopTagMateRepository(f.db)
	}
	return f.shop_tag_mate
}