package routes

import (
	"murilothebr/goalscorer/controllers"

	"github.com/gin-gonic/gin"
)

func PlayersRoute(router *gin.Engine) {
	router.POST("/player", controllers.CreatePlayer())
	router.GET("/player", controllers.GetPlayers())
}
