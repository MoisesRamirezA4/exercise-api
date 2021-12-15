package api

import (
	"github.com/epa-datos/exercise/github.com/epa-datos/exercise-api/api/posts"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	server := gin.Default()
	posts.RegisterRoutes(server)
	server.Run(":8080")
}
