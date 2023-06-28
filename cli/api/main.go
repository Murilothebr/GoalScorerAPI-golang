package main

import (
	"murilothebr/goalscorer/configs"
	"murilothebr/goalscorer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.PlayersRoute(router)

	configs.ConnectDB()

	router.Run("localhost:6000")
}
