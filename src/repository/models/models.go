package models

import (
	"gorm.io/gorm"
)

type InvoiceAddress struct {
	gorm.Model
	InvoiceID   uint32 `gorm:"column:invoice_id;index"`
	Network     string `gorm:"column:network"`
	Address     string `gorm:"column:address;index"`
	SeedVersion string `gorm:"column:seed_version;"`
}

type UserAddress struct {
	gorm.Model
	UserID      uint32 `gorm:"column:user_id;index"`
	Network     string `gorm:"column:network"`
	Address     string `gorm:"column:address;uniqueIndex"`
	SeedVersion string `gorm:"column:seed_version;"`
}
