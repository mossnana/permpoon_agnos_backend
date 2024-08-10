package password_test

import (
	"agnos/backend/pkgs/password"
	"testing"
)

func TestMPA(t *testing.T) {
	t.Run("Password length >=6, < 20 characters", func(t *testing.T) {
		testCases := []struct {
			name     string
			password string
			expected uint
		}{
			{
				name:     "Password too short",
				password: "Pass",
				expected: 2,
			},
			{
				name:     "Password at minimum length",
				password: "Pass12",
				expected: 0,
			},
			{
				name:     "Password at maximum length",
				password: "Password123456789",
				expected: 0,
			},
			{
				name:     "Password too long",
				password: "PasswordTooLongToPass",
				expected: 1,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if result := password.NewMPA().ByLength(tc.password); result != tc.expected {
					t.Errorf("Expected %v, but got %v for password %v", tc.expected, result, tc.password)
				}
			})
		}
	})
	t.Run("Contains at least 1 lowercase letter, at least 1 uppercase letter, and at least 1 digit", func(t *testing.T) {
		testCases := []struct {
			name     string
			password string
			expected uint
		}{
			{
				name:     "Valid password with lowercase, uppercase, and digit",
				password: "Passw0rd",
				expected: 0,
			},
			{
				name:     "Valid password with multiple lowercase, uppercase, and digits",
				password: "Str0ngPassw0rd",
				expected: 0,
			},
			{
				name:     "Valid password with @",
				password: "Str0ngP@ssw0rd",
				expected: 1,
			},
			{
				name:     "Invalid password with no lowercase",
				password: "PASSWORD123",
				expected: 1,
			},
			{
				name:     "Invalid password with no uppercase",
				password: "password123",
				expected: 1,
			},
			{
				name:     "Invalid password with no digit",
				password: "Password",
				expected: 1,
			},
			{
				name:     "Empty password",
				password: "",
				expected: 3,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if result := password.NewMPA().ByCharacter(tc.password); result != tc.expected {
					t.Errorf("Expected %v, but got %v for password %v", tc.expected, result, tc.password)
				}
			})
		}
	})
	t.Run("Does not contain 3 repeating characters in a row", func(t *testing.T) {
		testCases := []struct {
			name     string
			password string
			expected uint
		}{
			{
				name:     "Empty string",
				password: "",
				expected: 0,
			},
			{
				name:     "No repeating characters",
				password: "abcde",
				expected: 0,
			},
			{
				name:     "Two repeating characters",
				password: "abccdde",
				expected: 0,
			},
			{
				name:     "Three repeating characters",
				password: "abcccdde",
				expected: 1,
			},
			{
				name:     "Four repeating characters",
				password: "abccccdde",
				expected: 2,
			},
			{
				name:     "Mixed case",
				password: "abCCCdde",
				expected: 1,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if result := password.NewMPA().ByRepeatingCharacter(tc.password); result != tc.expected {
					t.Errorf("Expected %v, but got %v for input %v", tc.expected, result, tc.password)
				}
			})
		}
	})
	t.Run("Get minimum to generate strong password", func(t *testing.T) {
		testCases := []struct {
			name     string
			password string
			expected uint
		}{
			{
				name:     "Required case 1",
				password: "aA1",
				expected: 3,
			},
			{
				name:     "Required case 2",
				password: "1445D1cd",
				expected: 0,
			},
			{
				name:     "Valid password",
				password: "Str0ngPassw0rd",
				expected: 0,
			},
			{
				name:     "Invalid password, too short",
				password: "Pass",
				expected: 2,
			},
			{
				name:     "Invalid password, no uppercase",
				password: "password123",
				expected: 1,
			},
			{
				name:     "Invalid password, repeating characters",
				password: "abcccdde",
				expected: 3,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if result := password.NewMPA().GetMinimumActionToValid(tc.password); result != tc.expected {
					t.Errorf("Expected %v, but got %v for input %v", tc.expected, result, tc.password)
				}
			})
		}
	})
}
