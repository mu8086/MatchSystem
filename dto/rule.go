package dto

import (
	"github.com/MatchSystem/constant"
)

type Rule interface {
	Match(...any) bool
}

type Gender struct{}

func (g Gender) Match(target, desire constant.Gender) bool {
	return target == desire
}

type Height struct{}

func (h Height) Match(shortHeight, tallHeight int, less func(int, int) bool) bool {
	return less(shortHeight, tallHeight)
}
