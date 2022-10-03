package exception

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hadiyanfathur/golang-movies/helper"
	"github.com/hadiyanfathur/golang-movies/model/web"
	"gorm.io/gorm"
)

func ExceptionHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if validationErrors(writer, err) {
		return
	}

	if notFoundError(writer, err) {
		return
	}

	internalServerError(writer, err)
}

func validationErrors(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.BaseResponse{
			Code:         http.StatusBadRequest,
			Message:      "Bad Request",
			ErrorMessage: exception.Error(),
			Data:         nil,
		}

		helper.SetResponseBody(writer, webResponse)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, err interface{}) bool {
	if errors.Is(err.(error), gorm.ErrRecordNotFound) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.BaseResponse{
			Code:         http.StatusNotFound,
			Message:      "Not Found",
			ErrorMessage: gorm.ErrRecordNotFound.Error(),
			Data:         nil,
		}

		helper.SetResponseBody(writer, webResponse)
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.BaseResponse{
		Code:         http.StatusInternalServerError,
		Message:      "Internal Server Error",
		ErrorMessage: err.(error).Error(),
		Data:         nil,
	}

	helper.SetResponseBody(writer, webResponse)
}
