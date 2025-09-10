// utils/order.go
package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// GenerateOrderSn 生成订单号
func GenerateOrderSn() string {
	now := time.Now()
	return fmt.Sprintf("%s%d", now.Format("20060102150405"), now.Nanosecond()/1000)
}

// GenerateUUID 生成UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// FormatPrice 格式化价格
func FormatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}