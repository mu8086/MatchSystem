package matchingsystem

import (
	"fmt"
	"os"

	"github.com/MatchSystem/dto"
)

var Engine IMatchingSystem
var MatchQueue *chan *dto.User

type IMatchingSystem interface {
	Create(*dto.User)
	Get(string) (*dto.User, bool)
	GetMatchUserList(*dto.User, int) []*dto.User
	Remove(string)
	Run()
}

func InitEngine() {
	Engine = NewSimpleSystem(100)
	if simpleSystem, ok := Engine.(*SimpleSystem); ok {
		MatchQueue = &simpleSystem.MatchQueue
	} else {
		fmt.Printf("error: type-assertion SimpleSystem failed.")
		os.Exit(-1)
	}
	go Engine.Run()
}
