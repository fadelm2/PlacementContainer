package dto

type BlockResponse struct {
	ID        string `json:"id"`
	YardID    string `json:"yard_id"`
	Code      string `json:"code"`
	TotalSlot int    `json:"total_slot"`
	TotalRow  int    `json:"total_row"`
	TotalTier int    `json:"total_tier"`
}

type CreateBlockRequest struct {
	ID        string `json:"id" validate:"required"`
	YardID    string `json:"yard_id" validate:"required"`
	Code      string `json:"code" validate:"required"`
	TotalSlot int    `json:"total_slot" validate:"required,min=1"`
	TotalRow  int    `json:"total_row" validate:"required,min=1"`
	TotalTier int    `json:"total_tier" validate:"required,min=1"`
}

type UpdateBlockRequest struct {
	ID        string `json:"id" validate:"required"`
	YardID    string `json:"yard_id,omitempty"`
	Code      string `json:"name,omitempty"`
	TotalSlot *int   `json:"total_slot,omitempty"`
	TotalRow  *int   `json:"total_row,omitempty"`
	TotalTier *int   `json:"total_tier,omitempty"`
}

type GetBlockRequest struct {
	ID string `json:"-" validate:"required"`
}
