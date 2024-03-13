package controller

import (
	"net/http"
	"strings"
	"ticketsale/auth"
	"ticketsale/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var loginForm struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	var userM model.User
	// Bind the form to the struct
	if err := ctx.ShouldBindJSON(&loginForm); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	// Check account is email or account
	isEmail := strings.Contains(loginForm.Account, "@")
	// Check if the account exists. if not, return 404. if yes, check the password.
	checkCond := map[string]interface{}{}
	if isEmail {
		checkCond["email"] = loginForm.Account
	} else {
		checkCond["account"] = loginForm.Account
	}
	user, err := userM.Find(checkCond)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "User not found",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "Wrong password",
		})
		return
	}
	// Generate the JWT token
	tokenStr, err := auth.GenerateToken(user.Account, int(user.ID))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  500,
			"msg":   "Failed to generate the token",
			"error": err.Error(),
		})
		return
	}

	// store the token in cookie
	ctx.SetCookie("Auth-"+user.Account, tokenStr, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "Login success",
		"token": "Bearer " + tokenStr,
		"uid":   user.ID,
	})

}

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
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "Invalid form",
		})
		return
	}
	// Check if the email is already registered
	userM.Email = registerForm.Email
	if userM.CheckEmailExist() {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "Email already registered",
		})
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
		ctx.JSON(http.StatusOK, gin.H{
			"code":  500,
			"msg":   "Failed to hash the password",
			"error": err.Error(),
		})
	}
	userM.Password = string(hash)

	// Create the user
	result := userM.Create()
	if result != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  500,
			"msg":   "Failed to create the user",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "User created",
	})

}
