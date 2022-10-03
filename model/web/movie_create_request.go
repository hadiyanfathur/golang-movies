package web

type MovieCreateRequest struct {
	Name        string `validate:"required,min=1,max=255" json:"name"`
	Description string `json:"description"`
}
