package dto

import (
	"encoding/json"

	"github.com/MatchSystem/constant"
)

type Rule interface {
	MarshalJSON() ([]byte, error)
	Match(...any) bool
	String() string
}

type GenderRule struct{}

func (g GenderRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func (g GenderRule) Match(params ...any) bool {
	if len(params) != 2 {
		return false
	}

	gender1, ok1 := params[0].(constant.Gender)
	gender2, ok2 := params[1].(constant.Gender)

	return ok1 && ok2 && gender1 == gender2
}

func (g GenderRule) String() string {
	return "GenderRule"
}

type HeightRule struct{}

func (h HeightRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h HeightRule) Match(params ...any) bool {
	if len(params) != 2 {
		return false
	}

	height1, ok1 := params[0].(int)
	height2, ok2 := params[1].(int)
	return ok1 && ok2 && height1 < height2
}

func (h HeightRule) String() string {
	return "HeightRule"
}

var (
	genderRule Rule = GenderRule{}
	heightRule Rule = HeightRule{}
)
