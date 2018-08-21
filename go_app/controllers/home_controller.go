package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// you can import models
	//m "../models"
)

func HomeHandler(c *gin.Context) {
	// you can use model functions to do CRUD
	//
	// user, _ := m.FindUser(1)
	// u, err := json.Marshal(user)
	// if err != nil {
	// 	log.Printf("JSON encoding error: %v\n", err)
	// 	u = []byte("Get data error!")
	// }

	type Envs struct {
		GoOnRailsVer string
		GolangVer    string
	}

	gorVer := "0.3.1"
	golangVer := "go version go1.10.2 darwin/amd64"

	envs := Envs{GoOnRailsVer: gorVer, GolangVer: golangVer}
	c.HTML(http.StatusOK, "index.tmpl", envs)
}
