package app

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/rest"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/rest/resources"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/services"
	"github.com/google/wire"
)

func initApp() App {
	wire.Build(services.ProvideDatabaseService,
		services.ProvideCategoryService,
		services.ProvideVideoService,
		resources.ProvideCategoryRouter,
		resources.ProvideVideoRouter,
		rest.ProvideRouter, ProvideApp)
	return App{}
}
