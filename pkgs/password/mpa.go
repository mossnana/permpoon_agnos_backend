package password

import (
	"unicode"
)

const (
	MPA_MIN_LENGTH = 6
	MPA_MAX_LENGTH = 20
)

// Minimum Password Action
type MPA struct{}

func NewMPA() *MPA {
	return &MPA{}
}

func (mpa *MPA) ByLength(password string) (action uint) {
	passwordLength := len(password)
	if passwordLength < MPA_MIN_LENGTH {
		action += uint(MPA_MIN_LENGTH - passwordLength)
	} else if passwordLength > MPA_MAX_LENGTH {
		action += uint(passwordLength - MPA_MAX_LENGTH)
	}
	return
}

func (mpa *MPA) ByCharacter(password string) (action uint) {
	var isUpper, isLower, isDigit bool
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			isLower = true
		case unicode.IsUpper(char):
			isUpper = true
		case unicode.IsDigit(char):
			isDigit = true
		case char == '.' || char == '!':
		default:
			action += 1
		}
	}

	haveActions := []bool{isUpper, isLower, isDigit}
	for _, hasAction := range haveActions {
		if !hasAction {
			action += 1
		}
	}

	return
}

func (mpa *MPA) ByRepeatingCharacter(password string) (action uint) {
	count := map[rune]int{}
	for _, char := range password {
		count[char]++
		if count[char] > 2 {
			action += 1
		}
	}
	return
}

func (mpa *MPA) GetMinimumActionToValid(password string) (action uint) {
	action = mpa.ByLength(password)
	if action > 0 {
		return
	}
	return mpa.ByCharacter(password) + mpa.ByRepeatingCharacter(password)
}
