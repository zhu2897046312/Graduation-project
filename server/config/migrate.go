package config

import (
	"log"

	"server/models/core"
	"server/models/cms"
	"server/models/mp"
	"server/models/shop"
	"server/models/sp"
	"server/models/mypaypal"

	"gorm.io/gorm"
)

// Migrate 根据 models 自动迁移数据库表结构（只增不改删，与 GORM AutoMigrate 行为一致）
// 避免出现 Unknown column 'xxx' in 'where clause' 等表结构与 model 不一致的问题
func Migrate(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	log.Println("开始执行数据库迁移（AutoMigrate）...")

	err := db.AutoMigrate(
		// core
		&core.CoreAdmin{},
		&core.CoreDept{},
		&core.CoreRole{},
		&core.CorePermission{},
		&core.CoreConfig{},
		&core.CoreAdminRoleIndex{},
		// cms
		&cms.CmsDocument{},
		&cms.CmsDocumentArchive{},
		&cms.CmsRecommend{},
		&cms.CmsRecommendIndex{},
		// mp
		&mp.MpUser{},
		&mp.MpUserToken{},
		// shop
		&shop.ShopTag{},
		&shop.ShopTagIndex{},
		&shop.ShopTagMate{},
		// sp
		&sp.SpProduct{},
		&sp.SpCategory{},
		&sp.SpOrder{},
		&sp.SpOrderItem{},
		&sp.SpOrderReceiveAddress{},
		&sp.SpOrderRefund{},
		&sp.SpSku{},
		&sp.SpSkuIndex{},
		&sp.SpUserAddress{},
		&sp.SpUserCart{},
		&sp.SpProdAttributes{},
		&sp.SpProdAttributesValue{},
		&sp.SpProductContent{},
		&sp.SpProductProperty{},
		// paypal
		&mypaypal.PaypalOrderLogs{},
		&mypaypal.PaypalWebhookLogs{},
	)
	if err != nil {
		log.Printf("数据库迁移失败: %v", err)
		return err
	}
	log.Println("数据库迁移完成")
	return nil
}
