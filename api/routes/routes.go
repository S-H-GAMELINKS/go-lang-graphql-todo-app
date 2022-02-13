package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/handler"
	"github.com/S-H-GAMELINKS/go-lang-graphql-todo-app/api/graph"
	"github.com/S-H-GAMELINKS/go-lang-graphql-todo-app/api/graph/generated"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlPlaygroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func SetupRoutes(router *gin.Engine) {
	router.POST("/query", graphqlHandler())
	router.GET("/playground", graphqlPlaygroundHandler())
}