package dto

import (
	"github.com/MatchSystem/constant"
)

type User struct {
	Name   string
	Gender constant.Gender
	Height int
	Date   int
	Rules  []*Rule
}
