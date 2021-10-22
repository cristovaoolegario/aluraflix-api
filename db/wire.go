package db

import (
	"github.com/google/wire"
)

func initDatabaseService() DatabaseService {
	return ProvideDatabaseService()
}

func initCategoryService() CategoryService {
	wire.Build(ProvideDatabaseService, ProvideCategoryService)
	return CategoryService{}
}

func initVideoService() VideoService {
	wire.Build(ProvideDatabaseService, ProvideCategoryService, ProvideVideoService)
	return VideoService{}
}
