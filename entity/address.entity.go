package entity

type AddressEntity struct {
	ID           int64   `gorm:"primary_key:auto_increment" json:"id"`
	Username     string  `gorm:"type:varchar(200)" json:"username"`
	Long         float32 `gorm:"type:float" json:"long"`
	Lat          float32 `gorm:"type:float" json:"lat"`
	City         string  `gorm:"type:varchar(200)" json:"city"`
	Provience    string  `gorm:"type:varchar(200)" json:"provience"`
	Point        string  `gorm:"type:point" json:"-"`
	Full_Address string  `gorm:"type:text" json:"fullAddress"`
}
