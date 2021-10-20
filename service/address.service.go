package service

import (
	"github.com/dorman99/area-scan/dto"
	"github.com/dorman99/area-scan/entity"
	"github.com/dorman99/area-scan/model"
)

type AddressService interface {
	Insert(insertDto dto.InsertAddressDto) (*entity.AddressEntity, error)
	Find()
	FindNearestArea()
}

type addressService struct {
	addressModel model.AddressModel
}

func NewAddressService(addressModel model.AddressModel) AddressService {
	return &addressService{
		addressModel: addressModel,
	}
}

func (s *addressService) Insert(insertDto dto.InsertAddressDto) (*entity.AddressEntity, error) {
	return s.addressModel.Insert(insertDto)
}

func (s *addressService) Find()            {}
func (s *addressService) FindNearestArea() {}
