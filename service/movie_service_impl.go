package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/hadiyanfathur/golang-movies/helper"
	"github.com/hadiyanfathur/golang-movies/helper/mapper"
	"github.com/hadiyanfathur/golang-movies/model/entity"
	"github.com/hadiyanfathur/golang-movies/model/web"
	"github.com/hadiyanfathur/golang-movies/repository"
	"gorm.io/gorm"
)

type MovieServiceImpl struct {
	MovieRepository repository.MovieRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewMovieService(movieRepository repository.MovieRepository, DB *gorm.DB, validate *validator.Validate) *MovieServiceImpl {
	return &MovieServiceImpl{
		MovieRepository: movieRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *MovieServiceImpl) Save(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	movie := entity.Movie{
		Name:        request.Name,
		Description: request.Description,
	}

	movie = service.MovieRepository.Save(ctx, service.DB, movie)
	return mapper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	movie, err := service.MovieRepository.FindById(ctx, service.DB, request.Id)
	helper.PanicIfError(err)
	movie.Name = request.Name
	movie.Description = request.Description

	movie = service.MovieRepository.Update(ctx, service.DB, movie)
	return mapper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Delete(ctx context.Context, movieId uint) {
	movie, err := service.MovieRepository.FindById(ctx, service.DB, movieId)
	helper.PanicIfError(err)

	service.MovieRepository.Delete(ctx, service.DB, movie)
}

func (service *MovieServiceImpl) FindById(ctx context.Context, movieId uint) web.MovieResponse {
	movie, err := service.MovieRepository.FindById(ctx, service.DB, movieId)
	helper.PanicIfError(err)

	return mapper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) FindAll(ctx context.Context) []web.MovieResponse {
	movies := service.MovieRepository.FindAll(ctx, service.DB)

	return mapper.ToMovieResponses(movies)
}
