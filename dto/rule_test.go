package dto

import (
	"testing"

	"github.com/MatchSystem/constant"
)

func TestGenderRule_Match(t *testing.T) {
	type args struct {
		params []any
	}
	tests := []struct {
		name string
		args []any
		want bool
	}{
		{name: "arg size 1", args: []any{constant.Female}, want: false},
		{name: "arg size 3", args: []any{constant.Female, constant.Female, constant.Female}, want: false},
		{name: "same female gender", args: []any{constant.Female, constant.Female}, want: true},
		{name: "same male gender", args: []any{constant.Male, constant.Male}, want: true},
		{name: "female male", args: []any{constant.Female, constant.Male}, want: false},
		{name: "male female", args: []any{constant.Male, constant.Female}, want: false},
		{name: "not gender", args: []any{"male", "male"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GenderRule{}
			if got := g.Match(tt.args...); got != tt.want {
				t.Errorf("GenderRule.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeightRule_Match(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want bool
	}{
		{name: "arg size 1", args: []any{160}, want: false},
		{name: "arg size 3", args: []any{160, 170, 180}, want: false},
		{name: "short tall", args: []any{160, 170}, want: true},
		{name: "same", args: []any{160, 160}, want: false},
		{name: "tall short", args: []any{170, 160}, want: false},
		{name: "not height", args: []any{"", ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HeightRule{}
			if got := h.Match(tt.args...); got != tt.want {
				t.Errorf("HeightRule.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
