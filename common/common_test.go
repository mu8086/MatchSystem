package common

import (
	"testing"
)

func TestIsFemale(t *testing.T) {
	tests := []struct {
		name   string
		argStr string
		want   bool
	}{
		{name: "lowercase", argStr: "female", want: true},
		{name: "uppercase", argStr: "FEMALE", want: true},
		{name: "prefix", argStr: "female123", want: false},
		{name: "suffix", argStr: "123female", want: false},
		{name: "empty", argStr: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFemale(tt.argStr); got != tt.want {
				t.Errorf("IsFemale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMale(t *testing.T) {
	tests := []struct {
		name   string
		argStr string
		want   bool
	}{
		{name: "lowercase", argStr: "male", want: true},
		{name: "uppercase", argStr: "MALE", want: true},
		{name: "prefix", argStr: "male123", want: false},
		{name: "suffix", argStr: "123male", want: false},
		{name: "empty", argStr: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMale(tt.argStr); got != tt.want {
				t.Errorf("IsMale() = %v, want %v", got, tt.want)
			}
		})
	}
}