package entity

import "AutoCode/dbUtils/connect/sqliteConn"

type DatabaseConfig struct {
	Id       int64  `gorm:"type:int;primary_key;auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(50);not null;" json:"name"`
	DbType   string `gorm:"type:varchar(50);not null;" json:"dbType"`
	Host     string `gorm:"type:varchar(50);not null;" json:"host"`
	Port     string `gorm:"type:varchar(50);not null;" json:"port"`
	Username string `gorm:"type:varchar(50);not null;" json:"username"`
	Password string `gorm:"type:varchar(50);not null;" json:"password"`
	Database string `gorm:"type:varchar(200);not null;" json:"database"`
}

func init() {
	db := sqliteConn.GetDB()
	db.AutoMigrate(&DatabaseConfig{})
}
