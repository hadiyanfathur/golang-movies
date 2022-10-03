package app

import (
	"github.com/hadiyanfathur/golang-movies/controller"
	"github.com/hadiyanfathur/golang-movies/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(movieController controller.MovieController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/movies", movieController.FindAll)
	router.GET("/api/movies/:movieId", movieController.FindById)
	router.POST("/api/movies", movieController.Create)
	router.PUT("/api/movies/:movieId", movieController.Update)
	router.DELETE("/api/movies/:movieId", movieController.Delete)

	router.NotFound = exception.NotFoundHandler()
	router.PanicHandler = exception.ExceptionHandler

	return router
}
