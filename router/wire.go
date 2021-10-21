package router

import (
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func initVideoRouter() VideoRouter{
	wire.Build(db.ProvideVideoService, ProvideVideoRouter)
	return VideoRouter{}
}

func initRouter() *mux.Router{
	wire.Build(db.ProvideCategoryService,
		       db.ProvideVideoService,
		       ProvideVideoRouter,
		       ProvideRouter)

	return &mux.Router{}
}