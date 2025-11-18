package model

type PlaceRequest struct {
	BlockID   string
	SlotStart int
	SlotEnd   int
	RowStart  int
	RowEnd    int
}
