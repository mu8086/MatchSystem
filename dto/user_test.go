package dto

import (
	"testing"

	"github.com/MatchSystem/constant"
)

func TestUser_IsFemale(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "male",
			user: User{
				Gender: "male",
			},
			want: false,
		},
		{
			name: "female",
			user: User{
				Gender: "female",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.IsFemale(); got != tt.want {
				t.Errorf("User.IsFemale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_IsMale(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "male",
			user: User{
				Gender: "male",
			},
			want: true,
		},
		{
			name: "female",
			user: User{
				Gender: "female",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.IsMale(); got != tt.want {
				t.Errorf("User.IsMale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Match(t *testing.T) {
	type args struct {
		another *User
	}
	tall := 170
	short := 160
	tests := []struct {
		name string
		user User
		args args
		want bool
	}{
		{
			name: "GenderRule_Female_Female",
			user: User{Gender: constant.Female, Rules: []*Rule{&genderRule}},
			args: args{another: &User{Gender: constant.Female}},
			want: false,
		},
		{
			name: "GenderRule_Female_Male",
			user: User{Gender: constant.Female, Rules: []*Rule{&genderRule}},
			args: args{another: &User{Gender: constant.Male}},
			want: true,
		},
		{
			name: "GenderRule_Male_Female",
			user: User{Gender: constant.Male, Rules: []*Rule{&genderRule}},
			args: args{another: &User{Gender: constant.Female}},
			want: true,
		},
		{
			name: "GenderRule_Male_Male",
			user: User{Gender: constant.Male, Rules: []*Rule{&genderRule}},
			args: args{another: &User{Gender: constant.Male}},
			want: false,
		},
		{
			name: "HeightRule_Male_Taller_Other",
			user: User{Gender: constant.Male, Height: tall, Rules: []*Rule{&heightRule}},
			args: args{another: &User{Height: short}},
			want: true,
		},
		{
			name: "HeightRule_Male_Same_Other",
			user: User{Gender: constant.Male, Height: tall, Rules: []*Rule{&heightRule}},
			args: args{another: &User{Height: tall}},
			want: false,
		},
		{
			name: "HeightRule_Male_Shorter_Other",
			user: User{Gender: constant.Male, Height: short, Rules: []*Rule{&heightRule}},
			args: args{another: &User{Height: tall}},
			want: false,
		},
		{
			name: "HeightRule_Female_Taller_Other",
			user: User{Gender: constant.Female, Height: tall, Rules: []*Rule{&heightRule}},
			args: args{another: &User{Height: short}},
			want: false,
		},
		{
			name: "HeightRule_Female_Same_Other",
			user: User{Gender: constant.Female, Height: tall, Rules: []*Rule{&heightRule}},
			args: args{another: &User{Height: tall}},
			want: false,
		},
		{
			name: "HeightRule_Female_Shorter_Other",
			user: User{Gender: constant.Female, Height: short, Rules: []*Rule{&heightRule}},
			args: args{another: &User{Height: tall}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.Match(tt.args.another); got != tt.want {
				t.Errorf("User.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateName(t *testing.T) {
	tests := []struct {
		name    string
		argName string
		want    bool
	}{
		{name: "empty", argName: "", want: false},
		{name: "size1", argName: "1", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateName(tt.argName); got != tt.want {
				t.Errorf("validateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateGender(t *testing.T) {
	tests := []struct {
		name      string
		argGender string
		want      bool
	}{
		{name: "lowercase_female", argGender: "female", want: true},
		{name: "uppercase_female", argGender: "FEMALE", want: true},
		{name: "prefix_female", argGender: "femaleABC", want: false},
		{name: "suffix_female", argGender: "ABCfemale", want: false},
		{name: "lowecase_male", argGender: "male", want: true},
		{name: "uppercase_male", argGender: "MALE", want: true},
		{name: "prefix_male", argGender: "maleABC", want: false},
		{name: "suffix_male", argGender: "ABCmale", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateGender(tt.argGender); got != tt.want {
				t.Errorf("validateGender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHeight(t *testing.T) {
	tests := []struct {
		name      string
		argHeight int
		want      bool
	}{
		{name: "negative", argHeight: -1, want: false},
		{name: "zero", argHeight: 0, want: false},
		{name: "postive", argHeight: 1, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHeight(tt.argHeight); got != tt.want {
				t.Errorf("validateHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateDates(t *testing.T) {
	tests := []struct {
		name     string
		argDates int
		want     bool
	}{
		{name: "negative", argDates: -1, want: false},
		{name: "zero", argDates: 0, want: false},
		{name: "postive", argDates: 1, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateDates(tt.argDates); got != tt.want {
				t.Errorf("validateDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
