package matchingsystem

import (
	"fmt"
	"sync"

	"github.com/MatchSystem/dto"
)

func NewSimpleSystem(matchQueueSize int) *SimpleSystem {
	return &SimpleSystem{
		MatchQueue: make(chan *dto.User, matchQueueSize),
	}
}

type SimpleSystem struct {
	Users      sync.Map
	MatchQueue chan *dto.User
}

func (s *SimpleSystem) Create(user *dto.User) {
	key := user.Name
	s.Users.Store(key, user)
}

func (s *SimpleSystem) Get(name string) (*dto.User, bool) {
	u, exists := s.Users.Load(name)
	if !exists {
		return nil, false
	}

	user, ok := u.(*dto.User)
	if !ok {
		return nil, false
	}
	return user, true
}

func (s *SimpleSystem) GetMatchUserList(user *dto.User) (userList []*dto.User) {
	s.Users.Range(func(_, value interface{}) bool {
		u := value.(*dto.User)
		if u != user && user.Match(u) {
			userList = append(userList, u)
		}
		return true
	})
	return userList
}

func (s *SimpleSystem) Match(*dto.User, *dto.User) bool {
	return false
}

func (s *SimpleSystem) Print() string {
	all := "all: "
	s.Users.Range(func(_, value interface{}) bool {
		all += value.(*dto.User).String()
		return true
	})
	return all
}

func (s *SimpleSystem) Remove(name string) {
	s.Users.Delete(name)
}

func (s *SimpleSystem) Run() {
	for {
		user1 := <-s.MatchQueue
		user2 := <-s.MatchQueue

		// Atomic operation to decrease the number of wanted dates
		user1.Dates--
		user2.Dates--

		if user1.Dates <= 0 {
			s.Users.Delete(user1.Name)
		}

		if user2.Dates <= 0 {
			s.Users.Delete(user2.Name)
		}

		fmt.Printf("Matched: %s and %s\n", user1.Name, user2.Name)
	}
}
