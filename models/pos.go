package models

import "time"

type Category struct {
	CategoryID   int    `gorm:"primaryKey"`
	CategoryName string `gorm:"type:varchar(255);not null"`
}

type Item struct {
	ItemID         int    `gorm:"primaryKey"`
	ItemName       string `gorm:"type:varchar(255);not null"`
	ItemPrice      int    `gorm:"not null"`
	ItemStock      int    `gorm:"not null"`
	ItemCategoryID int    `gorm:"not null"`
	CreatedDate    time.Time
	UpdateDate     time.Time
}

type SellData struct {
	SellDataID         int    `gorm:"primaryKey"`
	SellItemName       string `gorm:"type:varchar(255);not null"`
	SellItemPrice      int    `gorm:"not null"`
	SellItemCategoryID int    `gorm:"not null"`
	SellItemID         int    `gorm:"not null"`
	SellCreatedDate    time.Time
}

type SellDataSummary struct {
	SellDataSummaryID          int       `gorm:"primaryKey"`
	SellItemSummaryPrice       int       `gorm:"not null"`
	SellItemSummaryCreatedDate time.Time `gorm:"not null"`
}

type User struct {
	UserId       uint   `gorm:"primaryKey"`
	UserUsername string `gorm:"type:varchar(255);not null"`
	UserPassword string `gorm:"type:varchar(255);not null"`
}
