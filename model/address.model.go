package model

import (
	"github.com/dorman99/area-scan/dto"
	"github.com/dorman99/area-scan/entity"
	"gorm.io/gorm"
)

type AddressModel interface {
	Insert(insertDto dto.InsertAddressDto) (entity.AddressEntity, error)
	Find()
	FindOthersNearest()
}

type addressModel struct {
	connection *gorm.DB
}

func NewAddressModel(db *gorm.DB) AddressModel {
	return &addressModel{
		connection: db,
	}
}

func (m *addressModel) Insert(insertDto dto.InsertAddressDto) (entity.AddressEntity, error) {
	var inserted entity.AddressEntity
	ress := m.connection.Save(insertDto).Take(&inserted)
	if ress.Error != nil {
		return inserted, ress.Error
	}
	return inserted, nil
}

func (m *addressModel) Find()              {}
func (m *addressModel) FindOthersNearest() {}
