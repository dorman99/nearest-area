package entity

type AddressEntity struct {
	ID           int64  `gorm:"primary_key:auto_increment" json:"id"`
	Username     string `gorm:"type:varchar(200)" json:"username"`
	Long         string `gorm:"type:varchar(200)" json:"long"`
	Lat          string `gorm:"type:varchar(200)" json:"lat"`
	City         string `gorm:"type:varchar(200)" json:"city"`
	Provience    string `gorm:"type:varchar(200)" json:"provience"`
	Full_Address string `gorm:"type:text" json:"fullAddress"`
}
