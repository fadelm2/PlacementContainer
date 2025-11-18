package model

type BlockResponse struct {
	ID        string `json:"id"`
	YardID    string `json:"yard_id"`
	Name      string `json:"name"`
	TotalSlot int    `json:"total_slot"`
	TotalRow  int    `json:"total_row"`
	TotalTier int    `json:"total_tier"`
}

type CreateBlockRequest struct {
	ID        string `json:"id" validate:"required"`
	YardID    string `json:"yard_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	TotalSlot int    `json:"total_slot" validate:"required,min=1"`
	TotalRow  int    `json:"total_row" validate:"required,min=1"`
	TotalTier int    `json:"total_tier" validate:"required,min=1"`
}

type UpdateBlockRequest struct {
	ID        string `json:"id" validate:"required"`
	YardID    string `json:"yard_id,omitempty"`
	Name      string `json:"name,omitempty"`
	TotalSlot *int   `json:"total_slot,omitempty"`
	TotalRow  *int   `json:"total_row,omitempty"`
	TotalTier *int   `json:"total_tier,omitempty"`
}

type GetBlockRequest struct {
	ID string `json:"-" validate:"required"`
}
