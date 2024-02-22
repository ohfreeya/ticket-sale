package controller

import (
	"net/http"
	"ticketsale/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var registerForm struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Phone    string `json:"phone"`
		Type     int8   `json:"type"`
		Sex      int8   `json:"sex"`
	}
	var userM model.User
	// Bind the form to the struct
	if err := ctx.ShouldBindJSON(&registerForm); err != nil {
		ctx.JSON(400, gin.H{"msg": "Invalid form", "error": err.Error()})
		return
	}
	// Check if the email is already registered
	userM.Email = registerForm.Email
	if userM.CheckEmailExist() {
		ctx.JSON(400, gin.H{"msg": "Email already registered"})
		return
	}

	// assign the form to the model
	userM.Account = registerForm.Account
	userM.Name = registerForm.Name
	userM.Email = registerForm.Email
	userM.Phone = registerForm.Phone
	userM.Type = registerForm.Type
	userM.Sex = registerForm.Sex

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to hash the password", 
			"error": err.Error(),
		})
	}
	userM.Password = string(hash)

	// Create the user
	result := userM.Create()
	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed to create the user",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User created",
	})

}
