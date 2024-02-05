package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MatchSystem/constant"
	"github.com/MatchSystem/dto"
	"github.com/MatchSystem/matchingsystem"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	matchingsystem.InitEngine()
}

var (
	shortFemale = dto.NewUser("Female-160", string(constant.Female), 160, 1)
	tallMale    = dto.NewUser("Male-170", string(constant.Male), 170, 3)
)

func TestAddSinglePersonAndMatch(t *testing.T) {
	testUser := tallMale

	userData := map[string]interface{}{
		"name":         testUser.Name,
		"gender":       testUser.Gender,
		"height":       testUser.Height,
		"wanted_dates": testUser.Dates,
	}

	jsonData, err := json.Marshal(userData)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	AddSinglePersonAndMatch(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRemoveSinglePerson(t *testing.T) {
	testUser := tallMale
	matchingsystem.Engine.Create(testUser)

	req, err := http.NewRequest("DELETE", "/users/"+testUser.Name, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "name", Value: testUser.Name})

	RemoveSinglePerson(c)

	assert.Equal(t, http.StatusOK, w.Code)

	_, exists := matchingsystem.Engine.Get(testUser.Name)
	assert.False(t, exists)
}

func TestQuerySinglePeople(t *testing.T) {
	testUser1 := tallMale
	testUser2 := shortFemale

	matchingsystem.Engine.Create(testUser1)
	matchingsystem.Engine.Create(testUser2)

	req, err := http.NewRequest("GET", "/users/"+testUser2.Name+"/matches?N=5", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "name", Value: testUser2.Name})

	QuerySinglePeople(c)

	assert.Equal(t, http.StatusOK, w.Code)

	if bs, err := json.Marshal([]dto.User{*testUser1}); err == nil {
		assert.Equal(t, w.Body.String(), string(bs))
	}
}
