package converter

import (
	"place-container/internal/entity"
	"place-container/internal/model"
)

func YardToResponse(yard *entity.Yard) *model.YardResponse {
	return &model.YardResponse{
		ID:          yard.ID,
		Name:        yard.Name,
		Description: yard.Description,
	}
}
