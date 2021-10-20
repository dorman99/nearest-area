package dto

type InsertAddressDto struct {
	Username  string `json:"username" gorm:"username" binding:"required"`
	Long      string `json:"long" gorm:"long" binding:"required"`
	Lat       string `json:"lat" gorm:"lat" binding:"required"`
	City      string `json:"city" gorm:"citry" binding:"required"`
	Provience string `json:"provience" gorm:"provience" binding:"required"`
}
