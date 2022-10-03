package exception

import (
	"net/http"

	"gorm.io/gorm"
)

func NotFound(writer http.ResponseWriter, request *http.Request)  {
	ExceptionHandler(writer,request, gorm.ErrRecordNotFound)
}
func NotFoundHandler() http.Handler  {
	return http.HandlerFunc(NotFound)	
}