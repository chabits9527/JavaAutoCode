package DatabaseConfigRepository

import (
	"AutoCode/dbUtils/connect/sqliteConn"
	"AutoCode/entity"
)

func Save(config *entity.DatabaseConfig) {
	db := sqliteConn.GetDB()
	db.Save(config)
}

func FindById(id int) entity.DatabaseConfig {
	db := sqliteConn.GetDB()
	e := entity.DatabaseConfig{}
	db.Find(&e, id)
	return e
}

func Page(pageNum, pageSize int) entity.PageInfo {
	var list []entity.DatabaseConfig
	var total int64
	db := sqliteConn.GetDB()
	db.Model(&entity.DatabaseConfig{}).Count(&total)
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list)
	return entity.PageInfo{
		List:     list,
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
}
