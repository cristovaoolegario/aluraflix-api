// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/rest"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/http/rest/resources"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/services"
)

// Injectors from wire.go:

func InitApp() App {
	databaseService := services.ProvideDatabaseService()
	categoryService := services.ProvideCategoryService(databaseService)
	videoService := services.ProvideVideoService(categoryService, databaseService)
	videoRouter := resources.ProvideVideoRouter(videoService)
	categoryRouter := resources.ProvideCategoryRouter(categoryService)
	router := rest.ProvideRouter(videoRouter, categoryRouter)
	app := ProvideApp(router, databaseService)
	return app
}
