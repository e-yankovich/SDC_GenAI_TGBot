package utils

import (
	"fmt"
	"testing"
	"unicode"
)

func TestInvertString(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Simple word",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "Sentence with spaces",
			input:    "Hello World",
			expected: "dlroW olleH",
		},
		{
			name:     "With numbers",
			input:    "abc123",
			expected: "321cba",
		},
		{
			name:     "With special characters",
			input:    "hello!@#$%",
			expected: "%$#@!olleh",
		},
		{
			name:     "Unicode characters",
			input:    "привет мир",
			expected: "рим тевирп",
		},
		{
			name:     "Emojis",
			input:    "Hello 👋 World 🌍",
			expected: "🌍 dlroW 👋 olleH",
		},
		{
			name:     "Palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "Mixed case",
			input:    "AbCdEf",
			expected: "fEdCbA",
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := InvertString(tc.input)
			if result != tc.expected {
				t.Errorf("InvertString(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

// Test that our inversion function correctly handles all types of strings
func TestInvertStringProperties(t *testing.T) {
	t.Run("Length preservation", func(t *testing.T) {
		inputs := []string{
			"",
			"a",
			"hello",
			"Hello World!",
			"привет мир 👋",
		}
		
		for _, input := range inputs {
			inverted := InvertString(input)
			if len(inverted) != len(input) {
				t.Errorf("Length of InvertString(%q) = %d; want %d", input, len(inverted), len(input))
			}
		}
	})
	
	t.Run("Double inversion equals identity", func(t *testing.T) {
		inputs := []string{
			"",
			"a",
			"hello",
			"Hello World!",
			"привет мир 👋",
		}
		
		for _, input := range inputs {
			doubleInverted := InvertString(InvertString(input))
			if doubleInverted != input {
				t.Errorf("InvertString(InvertString(%q)) = %q; want %q", input, doubleInverted, input)
			}
		}
	})
	
	t.Run("Character preservation", func(t *testing.T) {
		inputs := []string{
			"Hello World!",
			"123456789",
			"!@#$%^&*()",
			"привет мир 👋",
		}
		
		for _, input := range inputs {
			inverted := InvertString(input)
			inputRunes := []rune(input)
			invertedRunes := []rune(inverted)
			
			// Check that the inverted string contains the same characters
			if len(inputRunes) != len(invertedRunes) {
				t.Errorf("Length mismatch between input and inverted")
				continue
			}
			
			// Count each character type
			inputLetters, inputDigits, inputSpaces, inputOther := countCharTypes(input)
			invertedLetters, invertedDigits, invertedSpaces, invertedOther := countCharTypes(inverted)
			
			if inputLetters != invertedLetters {
				t.Errorf("Letter count mismatch: original=%d, inverted=%d", inputLetters, invertedLetters)
			}
			if inputDigits != invertedDigits {
				t.Errorf("Digit count mismatch: original=%d, inverted=%d", inputDigits, invertedDigits)
			}
			if inputSpaces != invertedSpaces {
				t.Errorf("Space count mismatch: original=%d, inverted=%d", inputSpaces, invertedSpaces)
			}
			if inputOther != invertedOther {
				t.Errorf("Other character count mismatch: original=%d, inverted=%d", inputOther, invertedOther)
			}
		}
	})
}

// Helper function to count character types in a string
func countCharTypes(s string) (letters, digits, spaces, other int) {
	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsSpace(r):
			spaces++
		default:
			other++
		}
	}
	return
}

// Test for performance with larger inputs
func BenchmarkInvertString(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			// Create a string of specified length
			var input string
			chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
			for i := 0; i < size; i++ {
				input += string(chars[i%len(chars)])
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				InvertString(input)
			}
		})
	}
} 