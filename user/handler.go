package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MatchSystem/dto"
	"github.com/MatchSystem/matchingsystem"
	"github.com/gin-gonic/gin"
)

func AddSinglePersonAndMatch(c *gin.Context) {
	var userParams struct {
		Name        string `json:"name"`
		Gender      string `json:"gender"`
		Height      int    `json:"height"`
		WantedDates int    `json:"dates"`
	}

	if err := c.BindJSON(&userParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := dto.NewUser(userParams.Name, userParams.Gender, userParams.Height, userParams.WantedDates)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Bad Request: %v", userParams)})
		return
	}

	if _, exists := matchingsystem.Engine.Get(userParams.Name); exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Bad Request: (%v) User exists.", userParams)})
		return
	}

	// Step 1: call engine to create user
	matchingsystem.Engine.Create(user)

	// Step 2: get user list that match the new user's rule
	userList := matchingsystem.Engine.GetMatchUserList(user, 100)

	c.JSON(http.StatusOK, userList)
}

func RemoveSinglePerson(c *gin.Context) {
	name := c.Param("name")

	_, exists := matchingsystem.Engine.Get(name)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found.", name)})
		return
	}

	matchingsystem.Engine.Remove(name)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s removed successfully", name)})
}

func QuerySinglePeople(c *gin.Context) {
	name := c.Param("name")

	N, ok := c.GetQuery("N")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'N' not provided"})
		return
	}

	// Parse 'N' as an integer
	maxSize, err := strconv.Atoi(N)
	if err != nil || maxSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'N'"})
		return
	}

	user, exists := matchingsystem.Engine.Get(name)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found.", name)})
		return
	}

	userList := matchingsystem.Engine.GetMatchUserList(user, maxSize)

	c.JSON(http.StatusOK, userList)
}

func Like(c *gin.Context) {
	name := c.Param("name")

	var userParams struct {
		LikedName string `json:"likedName"`
	}

	if err := c.BindJSON(&userParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := matchingsystem.Engine.Get(name)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	likedUser, exists := matchingsystem.Engine.Get(userParams.LikedName)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Liked user not found"})
		return
	}

	_, liked := user.LikedUsers[likedUser]
	if liked {
		// If two users match, add them to the match system queue
		*matchingsystem.MatchQueue <- user
		*matchingsystem.MatchQueue <- likedUser

		c.JSON(http.StatusOK, gin.H{"message": "Matched!"})
	} else {
		likedUser.LikedUsers[user] = struct{}{}
		c.JSON(http.StatusOK, gin.H{"message": "Liked!"})
	}
}
