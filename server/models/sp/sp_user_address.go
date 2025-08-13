package sp

import (
	"time"
)
type SpUserAddress struct {
    ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID        uint       `gorm:"not null;default:0" json:"userId"`
    Title         string     `gorm:"size:200;not null" json:"title"`
    DefaultStatus *bool      `json:"defaultStatus"`
    FirstName     string     `gorm:"size:64;not null;default:''" json:"firstName"`
    LastName      string     `gorm:"size:64;not null;default:''" json:"lastName"`
    Email         string     `gorm:"size:128;not null;default:''" json:"email"`
    Phone         string     `gorm:"size:64;not null;default:''" json:"phone"`
    Province      string     `gorm:"size:64;not null;default:''" json:"province"`
    City          string     `gorm:"size:64;not null;default:''" json:"city"`
    Region        string     `gorm:"size:64;not null;default:''" json:"region"`
    DetailAddress string     `gorm:"size:200;not null;default:''" json:"detailAddress"`
    CreatedTime   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    Country       string     `gorm:"size:64;not null;default:''" json:"country"`
    PostalCode    string     `gorm:"size:64;not null;default:''" json:"postalCode"`
}