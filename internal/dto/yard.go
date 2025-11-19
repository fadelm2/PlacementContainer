package dto

type YardResponse struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"name,omitempty"`
}

type CreateYardRequest struct {
	ID          string `json:"id" validate:"required,max=100"`
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=100"`
}

type UpdateYardRequest struct {
	ID          string `json:"-" validate:"required"`
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=100"`
}
type GetYardRequest struct {
	ID string `json:"-" validate:"required"`
}
