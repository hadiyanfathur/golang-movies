package repository

import (
	"context"

	"github.com/hadiyanfathur/golang-movies/model/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Save(ctx context.Context, db *gorm.DB, movie entity.Movie) entity.Movie
	Update(ctx context.Context, db *gorm.DB, movie entity.Movie) entity.Movie
	Delete(ctx context.Context, db *gorm.DB, movie entity.Movie)
	FindById(ctx context.Context, db *gorm.DB, movieId uint) (entity.Movie, error)
	FindAll(ctx context.Context, db *gorm.DB) []entity.Movie
}
