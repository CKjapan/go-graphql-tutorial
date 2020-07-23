package models

// Link is
type Link struct {
	ID      string `gorm:"column:ID;primary_key"`
	Title   string `gorm:"column:Title"`
	Address string `gorm:"column:Address"`
}

// User is
type User struct {
	ID   string `gorm:"column:ID;primary_key"`
	Name string `gorm:"column:Username"`
}
