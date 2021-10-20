package dto

type InsertAddressDto struct {
	Username    string  `json:"username" gorm:"username" binding:"required"`
	Long        float32 `json:"long" gorm:"long" binding:"required"`
	Lat         float32 `json:"lat" gorm:"lat" binding:"required"`
	City        string  `json:"city" gorm:"city" binding:"required"`
	Provience   string  `json:"provience" gorm:"provience" binding:"required"`
	FullAddress string  `json:"fullAddress" gorm:"full_address" binding:"required"`
}
