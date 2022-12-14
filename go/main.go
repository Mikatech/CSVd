package main

import (
	"github.com/Mikatech/CTF-AI/api"
	"github.com/Mikatech/CTF-AI/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	svc, err := database.NewSqliteDB("./database.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	env := api.Env{Svc: svc}

	router := gin.Default()
	router.GET("/user/:id", env.GetUser)
	router.GET("/users", env.GetUsers)
	router.POST("/user", env.CreateUser)
	router.DELETE("/user/:id", env.DeleteUser)
	log.Fatal(router.Run(":10000"))
}
