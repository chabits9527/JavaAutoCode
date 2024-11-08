package openAPI

import (
	"AutoCode/entity"
	"AutoCode/repository/DatabaseConfigRepository"
)

type OpenAPI struct {
}

func NewOpenAPI() *OpenAPI {
	return &OpenAPI{}
}

func (*OpenAPI) Save(config entity.DatabaseConfig) {
	DatabaseConfigRepository.Save(&config)
}

func (*OpenAPI) Page(pageNum, pageSize int) entity.PageInfo {
	return DatabaseConfigRepository.Page(pageNum, pageSize)
}
