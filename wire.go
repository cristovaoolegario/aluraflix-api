package main

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/router"
	"github.com/google/wire"
)

func initApp() App {
	wire.Build(db.ProvideDatabaseService,
			   db.ProvideCategoryService,
		       db.ProvideVideoService,
		       router.ProvideCategoryRouter,
			   router.ProvideVideoRouter,
			   router.ProvideRouter, ProvideApp)
	return App{}
}
