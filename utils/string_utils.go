package utils

// InvertString reverses the order of characters in a string
func InvertString(input string) string {
	// Convert string to rune slice to properly handle Unicode characters
	runes := []rune(input)
	
	// Get the length of the rune slice
	length := len(runes)
	
	// Perform in-place swapping
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	// Convert back to string and return
	return string(runes)
} 