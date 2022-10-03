//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/hadiyanfathur/golang-movies/app"
	"github.com/hadiyanfathur/golang-movies/controller"
	"github.com/hadiyanfathur/golang-movies/repository"
	"github.com/hadiyanfathur/golang-movies/service"
	"github.com/julienschmidt/httprouter"
)

var movieSet = wire.NewSet(
	repository.NewMovieRepositoryImpl,
	wire.Bind(new(repository.MovieRepository), new(*repository.MovieRepositoryImpl)),
	service.NewMovieService,
	wire.Bind(new(service.MovieService), new(*service.MovieServiceImpl)),
	controller.NewMovieController,
	wire.Bind(new(controller.MovieController), new(*controller.MovieControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.InitDB,
		validator.New,
		movieSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		NewServer,
	)

	return nil
}
