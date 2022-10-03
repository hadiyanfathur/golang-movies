package web

type MovieUpdateRequest struct {
	Id          uint   `json:"id"`
	Name        string `validate:"required,min=1,max=255" json:"name"`
	Description string `json:"description"`
}
