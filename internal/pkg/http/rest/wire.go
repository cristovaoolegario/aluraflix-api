package rest

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/rest/resources"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/services"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func initVideoRouter() resources.VideoRouter {
	wire.Build(services.ProvideVideoService, resources.ProvideVideoRouter)
	return resources.VideoRouter{}
}
func initCategoryRouter() resources.CategoryRouter {
	wire.Build(services.ProvideCategoryService, resources.ProvideCategoryRouter)
	return resources.CategoryRouter{}
}

func initRouter() *mux.Router {
	wire.Build(services.ProvideCategoryService,
		services.ProvideVideoService,
		resources.ProvideCategoryRouter,
		resources.ProvideVideoRouter,
		ProvideRouter)

	return &mux.Router{}
}
