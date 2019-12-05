package models

import (
	"time"
)

type Book struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt uint64
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"type:varchar(20);not null;"`
	Count     string     `gorm:"type:varchar(10);not null;"`
	Author    string     `gorm:"type:varchar(20);not null;"`
	Type      string     `gorm:"type:varchar(20);not null;"`
}
