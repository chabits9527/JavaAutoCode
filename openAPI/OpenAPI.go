package openAPI

import (
	"AutoCode/dbUtils/connect/mysqlConn"
	"AutoCode/entity"
	"AutoCode/repository/DatabaseConfigRepository"
)

type OpenAPI struct {
}

func NewOpenAPI() *OpenAPI {
	return &OpenAPI{}
}

func (*OpenAPI) SaveDatabaseConfig(config entity.DatabaseConfig) {
	DatabaseConfigRepository.Save(&config)
}

func (*OpenAPI) Page(pageNum, pageSize int) entity.PageInfo {
	return DatabaseConfigRepository.Page(pageNum, pageSize)
}

func (*OpenAPI) DeleteDatabaseConfig(id int64) {
	DatabaseConfigRepository.DeleteById(id)
}

func (*OpenAPI) TestConnect(config entity.DatabaseConfig) (bool, error) {
	return mysqlConn.MysqlTestConnect(config)
}
