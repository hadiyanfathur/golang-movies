package repository

import (
	"context"

	"github.com/hadiyanfathur/golang-movies/helper"
	"github.com/hadiyanfathur/golang-movies/model/entity"
	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {
}

func NewMovieRepositoryImpl() *MovieRepositoryImpl {
	return &MovieRepositoryImpl{}
}

func (repository *MovieRepositoryImpl) Save(ctx context.Context, db *gorm.DB, movie entity.Movie) entity.Movie {
	result := db.Create(&movie)
	helper.PanicIfError(result.Error)
	return movie
}

func (repository *MovieRepositoryImpl) Update(ctx context.Context, db *gorm.DB, movie entity.Movie) entity.Movie {
	result := db.Save(&movie)
	helper.PanicIfError(result.Error)
	return movie
}

func (repository *MovieRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, movie entity.Movie) {
	result := db.Delete(&movie)
	helper.PanicIfError(result.Error)
}

func (repository *MovieRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, movieId uint) (entity.Movie, error) {
	var movie entity.Movie
	result := db.First(&movie, movieId)
	helper.PanicIfError(result.Error)

	if result.RowsAffected > 0 {
		return movie, nil
	}

	return movie, result.Error
}

func (repository *MovieRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) []entity.Movie {
	movies := []entity.Movie{}
	result := db.Find(&movies)
	helper.PanicIfError(result.Error)
	return movies
}
