package sqliteConn

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("无法连接本地数据库")
	}
}

func GetDB() *gorm.DB {
	return db
}
