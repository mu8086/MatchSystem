package dto

import (
	"fmt"
	"strings"

	"github.com/MatchSystem/common"
	"github.com/MatchSystem/constant"
)

type User struct {
	Name       string
	Gender     constant.Gender
	Height     int
	Dates      int
	Rules      []*Rule
	LikedUsers map[*User]struct{}
}

func NewUser(name, gender string, height, dates int) *User {
	if !validate(name, gender, height, dates) {
		return nil
	}

	return &User{
		Name:       name,
		Gender:     constant.Gender(strings.ToLower(gender)),
		Height:     height,
		Dates:      dates,
		Rules:      []*Rule{&genderRule, &heightRule},
		LikedUsers: make(map[*User]struct{}),
	}
}

func (u *User) IsFemale() bool {
	return u.Gender == constant.Female
}

func (u *User) IsMale() bool {
	return u.Gender == constant.Male
}

func (u *User) String() string {
	rules := ""
	for _, rule := range u.Rules {
		switch rule := (*rule).(type) {
		case GenderRule, HeightRule:
			rules += rule.Name() + " "
		}
	}
	return fmt.Sprintf("[%v][%v][%v][%v][%v]", u.Name, u.Gender, u.Height, u.Dates, rules)
}

func (u *User) Match(another *User) bool {
	for _, rule := range u.Rules {
		switch rule := (*rule).(type) {
		case GenderRule:
			if u.IsFemale() {
				rule.Match(another.Gender, constant.Male)
			} else if u.IsMale() {
				rule.Match(another.Gender, constant.Female)
			} else {
				return false
			}

		case HeightRule:
			if u.IsFemale() {
				rule.Match(u.Height, another.Height)
			} else if u.IsMale() {
				rule.Match(another.Height, u.Height)
			} else {
				return false
			}

		default:
			fmt.Printf("NoRule, user: %s, another: %s\n", u.Name, another.Name)
			return false
		}
	}
	return true
}

func validateName(name string) bool {
	return len(name) > 0
}

func validateGender(gender string) bool {
	return common.IsFemale(gender) || common.IsMale(gender)
}

func validateHeight(height int) bool {
	return height > 0
}

func validateDates(dates int) bool {
	return dates > 0
}

func validate(name, gender string, height, dates int) bool {
	return validateName(name) && validateGender(gender) &&
		validateHeight(height) && validateDates(dates)
}
