package router

import (
	"github.com/MatchSystem/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	apiGroup := router.Group("/users")
	{
		apiGroup.POST("", user.AddSinglePersonAndMatch)
		apiGroup.DELETE("/:name", user.RemoveSinglePerson)
		apiGroup.GET("/:name/matches", user.QuerySinglePeople)
		apiGroup.POST("/:name/like", user.Like)
	}

	router.Run(":8080")
}
