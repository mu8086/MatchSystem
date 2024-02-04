package common

import (
	"strings"

	"github.com/MatchSystem/constant"
)

func IsFemale(s string) bool {
	return string(constant.Female) == strings.ToLower(s)
}

func IsMale(s string) bool {
	return string(constant.Male) == strings.ToLower(s)
}