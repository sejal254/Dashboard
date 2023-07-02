package controllers

import (
	"PMD/model"
	"PMD/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserPayload struct {
	Name     string `json:"userName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Password string `json:"password"`
}

var auth struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !model.CheckIfUserExists(auth.Email){
		c.JSON(http.StatusBadRequest,gin.H{"message":"User Not exist"})
		return
    }
	var u model.User
    u,err:=model.GetUserByCred(auth.Email)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"message":"Invalid email"})
		return
	}
	if !util.VerifyPassword(u.Password,auth.Password){
      c.JSON(http.StatusBadRequest,gin.H{
		"message":"Wrong credentials",
	  })
	  return;
	}

	
	if u.User_id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Welcome user!",
			"userId":   u.User_id,
			"userName": u.Username,
			"role":     u.Role,
		})

	}

}

func GetAllUsers(c *gin.Context) {
	const functionName = "controllers.GetAllUsers"
	u, err := model.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, u)
}

func DeleteUser(c *gin.Context) {
	const functionName = "controllers.DeleteUser"
	idp := c.Param("user_id")
	id, err := strconv.Atoi(idp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid ID"})
		return
	}
	err = model.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Query Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Deleted Successfully"})

}

func Adduser(c *gin.Context) {
	const funnctionName = "controllers.Adduser"
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	var dummy UserPayload
	err = json.Unmarshal(body, &dummy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to Unmarshal")
		return
	}

	if model.CheckIfUserExists(dummy.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exist"})
		return
	}
	hashPassword, err := util.HashPassword(dummy.Password)
	if err != nil {
		panic(err)

	}
	fmt.Println(hashPassword)
	var nuser model.User
	if dummy.Status == "" {
		dummy.Status = "Active"
	}
	nuser = model.User{
		Username: dummy.Name,
		Email:    dummy.Email,
		Role:     dummy.Role,
		Status:   dummy.Status,
		Password: hashPassword,
	}
	id, err := model.AddUser(&nuser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to Add")
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateUserRole(c *gin.Context) {

	const functionName = "controllers.UpdateRole"
	idp := c.Param("user_id")
	id, err := strconv.Atoi(idp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid ID"})
		return
	}
	err = model.UpdateRole(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Query Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Updated Successfully"})

}

func GetUserById(c *gin.Context) {
	const functionName = "controllers.GetUserById"
	idp := c.Param("user_id")
	id, err := strconv.Atoi(idp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid ID"})
		return
	}
	u, err := model.GetDetailsOfUserID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Query Error")
		return
	}
	c.JSON(http.StatusOK, u)
}
