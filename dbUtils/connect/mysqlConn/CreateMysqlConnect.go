package mysqlConn

import (
	"AutoCode/entity"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CreateConnect(config entity.DatabaseConfig) (*sql.DB, error) {
	datasourceStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", datasourceStr)
	if err != nil {
		return nil, fmt.Errorf("无法连接数据库: %w", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(fmt.Errorf("无法关闭数据库连接: %w", err))
		}
	}(db)
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("无法连接数据库: %w", err)
	}
	return db, nil
}

func MysqlTestConnect(config entity.DatabaseConfig) (bool, error) {
	connect, err := CreateConnect(config)
	if err != nil {
		return false, err
	}
	err = connect.Close()
	if err != nil {
		return false, fmt.Errorf("无法关闭数据库连接: %w", err)
	}
	return true, nil
}
