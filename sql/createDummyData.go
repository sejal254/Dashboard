package main

import (
	"PMD/model"
	"PMD/util"
	"fmt"
)

const (
	USERCOUNT    = 1000
	PROJECTCOUNT = 500
)

func main() {
	util.ConnectToDatabase()
	for i := 0; i < USERCOUNT; i++ {
		role := "User"
		if i%9 == 0 {
			role = "Admin"
		}
		u := model.User{
			Username: fmt.Sprintf("User%d", i),
			Email:    fmt.Sprintf("User%d@testmail.com", i),
			Role:     role,
			Status:   "Active",
		}
		model.AddUser(&u)
	}
	for i := 0; i < PROJECTCOUNT; i++ {
		u := model.Project{
			Name:   fmt.Sprintf("Project%d", i),
			Status: "Active",
		}
		model.AddProject(&u)
	}
}
