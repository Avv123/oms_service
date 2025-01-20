package handlers

import (
	//"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oms/models"
	"oms/repository"
)

func GetAllUsers(ctx *gin.Context) {
	users, err := repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(users)
	ctx.JSON(http.StatusOK, users)

}

func GetUserById(ctx *gin.Context) {

	id := ctx.Param("id")
	user, _ := repository.GetByID(id)
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": user})

}

func RegisterUser(ctx *gin.Context) {
	fmt.Println("issbs")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//if user.Username == "" || user.Password == "" || user.Email == "" {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": "fields required"})
	//	return
	//}
	passwordhash, _ := models.GeneratefromPassword(user.Password)
	user.Password = passwordhash
	//geturi := fmt.Sprintf("")
	err := repository.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"User created": user})
}
func Login(ctx *gin.Context) {
	var incominguser models.User
	if err := ctx.ShouldBindJSON(&incominguser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(incominguser)

	if incominguser.Password == "" || incominguser.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "fields required"})
		return
	}
	user, err := repository.GetUserByEmail(incominguser.Email)
	if err != nil || !user.CheckPassword(incominguser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})

}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := repository.DeleteByID(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}
