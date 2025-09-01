package paypal

import (
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	db                  *gorm.DB
	paypal_order_logs   *PaypalOrderLogsRepository
	paypal_webhook_logs *PaypalWebhookLogsRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) GetDB() *gorm.DB {
	return f.db
}

func (f *RepositoryFactory) GetPaypalOrderLogsRepository() *PaypalOrderLogsRepository {
	if f.paypal_order_logs == nil {
		f.paypal_order_logs = NewPaypalOrderLogsRepository(f.db)
	}
	return f.paypal_order_logs
}

func (f *RepositoryFactory) GetPaypalWebhookLogsRepository() *PaypalWebhookLogsRepository {
	if f.paypal_webhook_logs == nil {
		f.paypal_webhook_logs = NewPaypalWebhookLogsRepository(f.db)
	}
	return f.paypal_webhook_logs
}