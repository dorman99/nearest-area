package controller

import (
	"net/http"

	commoon "github.com/dorman99/area-scan/common"
	"github.com/dorman99/area-scan/dto"
	"github.com/dorman99/area-scan/service"
	"github.com/gin-gonic/gin"
)

type AddressController interface {
	Find(ctx *gin.Context)
	Insert(ctx *gin.Context)
	FindNearestArea(ctx *gin.Context)
}

type addressController struct {
	addressService service.AddressService
}

func NewAddressController(addressService service.AddressService) AddressController {
	return &addressController{
		addressService: addressService,
	}
}

func (c *addressController) Find(ctx *gin.Context) {}
func (c *addressController) Insert(ctx *gin.Context) {
	var insertDto dto.InsertAddressDto
	errDto := ctx.ShouldBind(&insertDto)
	if errDto != nil {
		response := commoon.FailedResponse("bad request", errDto.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result, err := c.addressService.Insert(insertDto)
	if err != nil {
		response := commoon.FailedResponse("bad request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := commoon.SuccessResponse(true, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) FindNearestArea(ctx *gin.Context) {
	var findNearestDto dto.FindNearestDto
	errDto := ctx.ShouldBind(&findNearestDto)
	if errDto != nil {
		response := commoon.FailedResponse("bad request", errDto.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	results, err := c.addressService.FindNearestArea(findNearestDto.Lat, findNearestDto.Long)
	if err != nil {
		response := commoon.FailedResponse("bad request", errDto.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := commoon.SuccessResponse(true, "success", results)
	ctx.JSON(http.StatusOK, response)
}
