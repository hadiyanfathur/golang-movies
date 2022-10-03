package mapper

import (
	"github.com/hadiyanfathur/golang-movies/model/entity"
	"github.com/hadiyanfathur/golang-movies/model/web"
)

func ToMovieResponse(movie entity.Movie) web.MovieResponse {
	return web.MovieResponse{
		Id:          movie.ID,
		Name:        movie.Name,
		Description: movie.Description,
	}
}

func ToMovieResponses(movies []entity.Movie) []web.MovieResponse {
	movieResponses := []web.MovieResponse{}
	for _, movie := range movies {
		movieResponses = append(movieResponses, ToMovieResponse(movie))
	}
	return movieResponses
}
