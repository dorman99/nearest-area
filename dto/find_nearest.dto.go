package dto

type FindNearestDto struct {
	Lat  float32 `json:"lat" binding:"required"`
	Long float32 `json:"long" binding:"required"`
}
