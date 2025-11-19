package dto

type PlaceRequest struct {
	BlockID   string
	SlotStart int
	SlotEnd   int
	RowStart  int
	RowEnd    int
}
