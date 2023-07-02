package controllers

import (
	"PMD/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProjectsOfUserID(c *gin.Context) {
	const functionName = "controllers.GetProjectsOfUserID"
	idn := c.Param("user_id")
	id, err := strconv.Atoi(idn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Invalid User ID"})
		return
	}
	p, err := model.GetProjectsByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot Fetch Projects Of User"})
		return
	}
	c.JSON(http.StatusOK, p)

}

func AssignProjectByID(c *gin.Context) {
	const functionName = "controllers.AssignProjectByName"
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Invalid Body"})
		return
	}
	var dummy model.UserProject
	err = json.Unmarshal(body, &dummy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to Unmarshal"})
		return
	}
	if model.IsUserAssignedToProject(dummy) {
		c.JSON(http.StatusOK, gin.H{"Error": "User already working on project"})
		return
	}
	id, err := model.AssignProjectByID(dummy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to Assign"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": fmt.Sprintf("Assigned %d", id)})
}
