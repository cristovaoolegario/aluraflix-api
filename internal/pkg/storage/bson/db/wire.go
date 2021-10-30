package db

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/services"
	"github.com/google/wire"
)

func initDatabaseService() services.DatabaseService {
	return services.ProvideDatabaseService()
}

func initCategoryService() services.CategoryService {
	wire.Build(services.ProvideDatabaseService, services.ProvideCategoryService)
	return services.CategoryService{}
}

func initVideoService() services.VideoService {
	wire.Build(services.ProvideDatabaseService, services.ProvideCategoryService, services.ProvideVideoService)
	return services.VideoService{}
}
