package controllers

import (
	"net/http"

	"github.com/biseshbhattarai/talwar/inputs"
	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func StoreEmail(context *gin.Context) {
	var input inputs.EmailInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := models.Email{
		Email:    input.Email,
		ScanType: input.ScanType,
	}
	savedEmail, err := email.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"email": savedEmail})

}

func SendEmails(context *gin.Context) {
	emails, err := models.ListEmails()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.SendEmails(emails, "test")
	context.JSON(http.StatusOK, gin.H{"message": "Emails sent"})
}
