package models

import (
	"database/sql"
	"gorm.io/gorm"
	"strings"
	"time"
)

// Модель текущего значения левелов у персонажей
type CharLevels struct {
	ID               uint
	AccountId        uint
	Level            uint
	LastSessionStart sql.NullTime
	LevelDuration    sql.NullString
}

// Модель данных о времени получения уровня
type LevelsSummary struct {
	ID               uint
	Level            int
	CharId           uint
	LevelDuration    string
	LevelGtmDuration string
	ReceivedAt       time.Time
}

// Модель ошибок коллектора левелов
type LevelError struct {
	ID        uint
	Type      string
	Message   string
	Level     string
	Event     string
	ShardGuid string
	CreatedAt time.Time
}

// Для использования в Scope при запросе в бд
func LevelsSummaryTable(shardName string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tableName := "levels_summary_" + strings.ReplaceAll(shardName, " ", "_")
		return tx.Table(tableName)
	}
}

//Для использования в Scope при запросе к БД
func CharLevelsTable(shardName string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tableName := "char_levels_" + strings.ReplaceAll(shardName, " ", "_")
		return tx.Table(tableName)
	}
}
