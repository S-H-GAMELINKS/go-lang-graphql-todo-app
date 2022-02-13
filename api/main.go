package main

import (
	"github.com/gin-gonic/gin"
	"github.com/S-H-GAMELINKS/go-lang-graphql-todo-app/api/routes"
)

func main() {
	r := gin.New()

	routes.SetupRoutes(r)

	r.Run()
}