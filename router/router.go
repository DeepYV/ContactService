package router

import (
	"github.com/gin-gonic/gin"
	repositories "test.com/test/repository"
	services "test.com/test/service"
)

func Router(router *gin.Engine) {

	userDBClient := repositories.NewRepository()
	usercontroller := services.NewService(userDBClient)

	router.POST(CREATECONTACT, usercontroller.CreateContact)
	router.GET(GETCONTACT, usercontroller.GetContact)
	router.GET(GETCONTACTBYNMAE, usercontroller.GetContactByName)
	router.DELETE(DELETECONTACT, usercontroller.DeleteContact)
	router.PATCH(UPDATECONTACT, usercontroller.UpdateContact)
	return
}
