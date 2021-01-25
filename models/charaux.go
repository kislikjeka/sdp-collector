package models

import "time"

type LevelUp struct {
	CharId int `gorm:"column:charid"`
	Level  int
	Time   time.Time
}

func (LevelUp) TableName() string {
	return "levelup"
}
