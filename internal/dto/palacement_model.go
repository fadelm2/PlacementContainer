package dto

type CreatePlacementRequest struct {
	BlockID     string `json:"block_id" binding:"required"`
	ContainerID string `json:"container_id" binding:"required"`
	RowNum      int    `json:"row_num" binding:"required"`
	Tier        int    `json:"tier" binding:"required"`
	Slot        string `json:"slot" binding:"required"`
}

//	type PlaceRequest struct {
//		BlockID string
//		Container dto.Container
//		SlotStart int
//		SlotEnd int
//		RowStart int
//		RowEnd int
//	}
type PlacementRequest struct {
	Yard            string  `json:"yard" validate:"required"`
	ContainerNumber string  `json:"container_number" validate:"required"`
	Block           string  `json:"block" validate:"required"`
	Slot            int     `json:"slot" validate:"required"`
	Row             int     `json:"row" validate:"required"`
	Tier            int     `json:"tier" validate:"required"`
	Size            int     `json:"size" validate:"required,oneof=20 40"`
	Type            string  `json:"type" validate:"required"`
	Height          float64 `json:"height" validate:"required"`
}

type SuggestionRequest struct {
	Yard            string  `json:"yard" validate:"required"`
	ContainerNumber string  `json:"container_number" validate:"required"`
	ContainerSize   int     `json:"container_size" validate:"required,oneof=20 40"`
	ContainerHeight float64 `json:"container_height" validate:"required"`
	ContainerType   string  `json:"container_type" validate:"required"`
}

type SuggestionResponse struct {
	Block string `json:"block"`
	Slot  int    `json:"slot"`
	Row   int    `json:"row"`
	Tier  int    `json:"tier"`
	Width int    `json:"width"`
}

type PickupRequest struct {
	Yard            string `json:"yard" validate:"required"`
	ContainerNumber string `json:"container_number" validate:"required"`
}
