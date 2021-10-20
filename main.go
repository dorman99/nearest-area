package main

import (
	"net/http"

	"github.com/dorman99/area-scan/controller"
	"github.com/dorman99/area-scan/model"
	"github.com/dorman99/area-scan/provider"
	"github.com/dorman99/area-scan/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = provider.InitPsql()

func main() {

	addressModel := model.NewAddressModel(db)
	addressService := service.NewAddressService(addressModel)
	addressController := controller.NewAddressController(addressService)

	defer provider.ClosePsqlDBConnection(db)
	app := gin.Default()
	r := app.Group("/app/v1")
	r.Handle(http.MethodPost, "/addresses", addressController.Insert)
	app.Run()
}
