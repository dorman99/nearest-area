package model

import (
	"fmt"

	"github.com/dorman99/area-scan/dto"
	"github.com/dorman99/area-scan/entity"
	"gorm.io/gorm"
)

type AddressModel interface {
	Insert(insertDto dto.InsertAddressDto) (*entity.AddressEntity, error)
	Find()
	FindOthersNearest(lat float32, long float32) ([]entity.AddressEntity, error)
}

type addressModel struct {
	connection *gorm.DB
}

func NewAddressModel(db *gorm.DB) AddressModel {
	return &addressModel{
		connection: db,
	}
}

func (m *addressModel) Insert(insertDto dto.InsertAddressDto) (*entity.AddressEntity, error) {
	inserted := &entity.AddressEntity{
		Username:     insertDto.Username,
		Lat:          insertDto.Lat,
		Long:         insertDto.Long,
		Full_Address: insertDto.FullAddress,
		Provience:    insertDto.Provience,
		City:         insertDto.City,
		Point:        fmt.Sprintf("(%f, %f)", insertDto.Lat, insertDto.Long),
	}
	ress := m.connection.Create(inserted)
	if ress.Error != nil {
		return inserted, ress.Error
	}
	return inserted, nil
}

func (m *addressModel) Find() {}
func (m *addressModel) FindOthersNearest(lat float32, long float32) ([]entity.AddressEntity, error) {
	var addresses []entity.AddressEntity
	wq := fmt.Sprintf("distance(%f, %f, lat, long) * 1000 < 10000", lat, long)
	ress := m.connection.Find(&addresses, wq)
	if ress.Error != nil {
		return nil, ress.Error
	}
	return addresses, nil
}
