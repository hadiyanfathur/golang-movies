package web

type BaseResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}
