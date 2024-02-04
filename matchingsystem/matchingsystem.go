package matchingsystem

import (
	"github.com/MatchSystem/dto"
)

var Engine IMatchingSystem

type IMatchingSystem interface {
	Create(*dto.User)
	Get(*dto.User) []*dto.User
	Match(*dto.User, *dto.User) bool
	Remove(*dto.User)
	Run()
}

func InitEngine() {
	Engine = NewSimpleSystem()
	go Engine.Run()
}
