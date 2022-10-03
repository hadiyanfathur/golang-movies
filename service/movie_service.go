package service

import (
	"context"

	"github.com/hadiyanfathur/golang-movies/model/web"
)

type MovieService interface {
	Save(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse
	Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse
	Delete(ctx context.Context, movieId uint)
	FindById(ctx context.Context, movieId uint) web.MovieResponse
	FindAll(ctx context.Context) []web.MovieResponse
}
