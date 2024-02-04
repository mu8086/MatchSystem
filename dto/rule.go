package dto

import (
	"github.com/MatchSystem/constant"
)

type Rule interface {
	Match(...any) bool
	Name() string
}

type GenderRule struct{}

func (g GenderRule) Match(params ...any) bool {
	if len(params) != 2 {
		return false
	}

	gender1, ok1 := params[0].(constant.Gender)
	gender2, ok2 := params[1].(constant.Gender)

	return ok1 && ok2 && gender1 == gender2
}

func (g GenderRule) Name() string {
	return "GenderRule"
}

type HeightRule struct{}

func (h HeightRule) Match(params ...any) bool {
	if len(params) != 2 {
		return false
	}

	height1, ok1 := params[0].(int)
	height2, ok2 := params[1].(int)
	return ok1 && ok2 && height1 < height2
}

func (h HeightRule) Name() string {
	return "HeightRule"
}

var (
	genderRule Rule = GenderRule{}
	heightRule Rule = HeightRule{}
)
