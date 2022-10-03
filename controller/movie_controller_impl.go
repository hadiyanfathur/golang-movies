package controller

import (
	"net/http"
	"strconv"

	"github.com/hadiyanfathur/golang-movies/helper"
	"github.com/hadiyanfathur/golang-movies/model/web"
	"github.com/hadiyanfathur/golang-movies/service"
	"github.com/julienschmidt/httprouter"
)

type MovieControllerImpl struct {
	MovieService service.MovieService
}

func NewMovieController(movieService service.MovieService) *MovieControllerImpl {
	return &MovieControllerImpl{
		MovieService: movieService,
	}
}

func (controller *MovieControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieCreateRequest := web.MovieCreateRequest{}
	helper.GetRequestBody(request, &movieCreateRequest)

	movieResponse := controller.MovieService.Save(request.Context(), movieCreateRequest)
	response := web.BaseResponse{
		Code:    200,
		Message: "success",
		Data:    movieResponse,
	}

	helper.SetResponseBody(writer, response)
}

func (controller *MovieControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieUpdateRequest := web.MovieUpdateRequest{}
	helper.GetRequestBody(request, &movieUpdateRequest)

	movieId := params.ByName("movieId")
	id, err := strconv.ParseUint(movieId, 10, 0)
	helper.PanicIfError(err)

	movieUpdateRequest.Id = uint(id)

	movieResponse := controller.MovieService.Update(request.Context(), movieUpdateRequest)
	response := web.BaseResponse{
		Code:    200,
		Message: "success",
		Data:    movieResponse,
	}

	helper.SetResponseBody(writer, response)
}

func (controller *MovieControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieId := params.ByName("movieId")
	id, err := strconv.ParseUint(movieId, 10, 0)
	helper.PanicIfError(err)

	controller.MovieService.Delete(request.Context(), uint(id))
	response := web.BaseResponse{
		Code:    200,
		Message: "success",
		Data:    nil,
	}

	helper.SetResponseBody(writer, response)
}

func (controller *MovieControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieId := params.ByName("movieId")
	id, err := strconv.ParseUint(movieId, 10, 0)
	helper.PanicIfError(err)

	movieResponse := controller.MovieService.FindById(request.Context(), uint(id))
	response := web.BaseResponse{
		Code:    200,
		Message: "success",
		Data:    movieResponse,
	}

	helper.SetResponseBody(writer, response)
}

func (controller *MovieControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieResponses := controller.MovieService.FindAll(request.Context())
	response := web.BaseResponse{
		Code:    200,
		Message: "success",
		Data:    movieResponses,
	}

	helper.SetResponseBody(writer, response)
}
