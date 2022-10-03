package helper

import (
	"encoding/json"
	"net/http"
)

const ContentTypeJSON = "application/json"

func GetRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func SetResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", ContentTypeJSON)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
