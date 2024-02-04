package user

import (
	"github.com/MatchSystem/dto"
	"github.com/MatchSystem/matchingsystem"
	"github.com/gin-gonic/gin"
)

func AddSinglePersonAndMatch(c *gin.Context) {
	matchingsystem.Engine.Create(&dto.User{})
}

func RemoveSinglePerson(c *gin.Context) {
	matchingsystem.Engine.Remove(&dto.User{})
}

func QuerySinglePeople(c *gin.Context) {
	matchingsystem.Engine.Get(&dto.User{})
}

func Like(c *gin.Context) {
	
}
