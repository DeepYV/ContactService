package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	model "test.com/test/Model"
	repositories "test.com/test/repository"
	"test.com/test/request"
)

type IService interface {
	CreateContact(ctx *gin.Context) error
	UpdateContact(ctx *gin.Context) error
	DeleteContact(ctx *gin.Context) error
	GetContact(ctx *gin.Context) error
	GetContactByName(ctx *gin.Context) error
}

type Service struct {
	service repositories.IRepository
}

func NewService(servicedata repositories.IRepository) *Service {
	return &Service{service: servicedata}

}
func (u *Service) CreateContact(ctx *gin.Context) {
	var contact model.Contact
	if err := json.NewDecoder(ctx.Request.Body).Decode(&contact); err != nil {
		log.Fatalln("error while decoding request body", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Wrong body "})
		return
	}
	if contact.Validate() != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "retry again with correct format"})
		return
	}

	err := u.service.CreateContact(ctx, contact)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"contact added": contact.Name})
	return
}

func (u *Service) GetContact(ctx *gin.Context) {

	GetContact, GetNumber, err := u.service.GetContact(ctx)
	if err != nil {

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return
	}

	ctx.IndentedJSON(http.StatusAccepted,
		gin.H{
			"name":   GetContact,
			"number": GetNumber,
		})

	return

}

func (u *Service) DeleteContact(ctx *gin.Context) {

	name := ctx.Params.ByName("name")

	err := u.service.DeleteContact(ctx, name)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again deleting "})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"contact deleted": name})
	return

}
func (u *Service) UpdateContact(ctx *gin.Context) {

	var updatedcontact request.UpdateContact

	if err := json.NewDecoder(ctx.Request.Body).Decode(&updatedcontact); err != nil {
		log.Fatalln("error while decoding request body", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Wrong body "})
		return
	}
	err := u.service.UpdatedContact(ctx, updatedcontact)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again updating number"})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{" Number changed To ": *updatedcontact.Number})
	return

}

func (u *Service) GetContactByName(ctx *gin.Context) {

	name := ctx.Request.URL.Query().Get("name")

	number, err := u.service.GetContactByName(ctx, name)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again fetching "})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"contact number": *number})
	return

}
