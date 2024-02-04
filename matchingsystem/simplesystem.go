package matchingsystem

import "github.com/MatchSystem/dto"

func NewSimpleSystem() *SimpleSystem {
	return &SimpleSystem{}
}

type SimpleSystem struct{}

func (s *SimpleSystem) Create(*dto.User) {

}

func (s *SimpleSystem) Get(*dto.User) []*dto.User {
	return nil
}

func (s *SimpleSystem) Match(*dto.User, *dto.User) bool {
	return false
}

func (s *SimpleSystem) Remove(*dto.User) {

}

func (s *SimpleSystem) Run() {
	
}