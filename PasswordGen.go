package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

func main() {
	// Define command-line flags
	length := flag.Int("length", 12, "Password length")
	complexity := flag.String("complexity", "medium", "Complexity level: low, medium, high")
	file := flag.String("file", "", "File to save the password (optional)")
	flag.Parse()

	// Validate length
	if *length < 1 {
		fmt.Println("Error: length must be greater than 0")
		os.Exit(1)
	}

	// Define character sets based on complexity
	var chars string
	switch *complexity {
	case "low":
		chars = "abcdefghijklmnopqrstuvwxyz"
	case "medium":
		chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case "high":
		chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?"
	default:
		fmt.Println("Error: complexity must be 'low', 'medium', or 'high'")
		os.Exit(1)
	}

	// Generate the password
	password, err := generatePassword(*length, chars, *complexity)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	// Output the password to the console
	fmt.Println("Generated password:", password)

	// Save to file if specified
	if *file != "" {
		err = os.WriteFile(*file, []byte(password), 0644)
		if err != nil {
			fmt.Printf("Error saving password to file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Password saved to %s\n", *file)
	}
}

// generatePassword generates a random password based on length, characters, and complexity
func generatePassword(length int, chars, complexity string) (string, error) {
	// For "high" complexity, ensure at least one of each required type
	if complexity == "high" && length >= 4 { // Minimum length to fit all required types
		return generateHighComplexityPassword(length)
	}

	// Standard generation for "low" and "medium"
	password := make([]byte, length)
	max := big.NewInt(int64(len(chars)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		password[i] = chars[n.Int64()]
	}

	return string(password), nil
}

// generateHighComplexityPassword ensures at least one lowercase, uppercase, digit, and symbol
func generateHighComplexityPassword(length int) (string, error) {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	symbols := "!@#$%^&*()_+-=[]{}|;:,.<>?"
	allChars := lowercase + uppercase + digits + symbols

	// Initialize password slice
	password := make([]byte, length)
	max := big.NewInt(int64(len(allChars)))

	// Ensure one of each required type in the first 4 positions
	choices := []string{lowercase, uppercase, digits, symbols}
	for i := 0; i < 4; i++ {
		maxChoice := big.NewInt(int64(len(choices[i])))
		n, err := rand.Int(rand.Reader, maxChoice)
		if err != nil {
			return "", err
		}
		password[i] = choices[i][n.Int64()]
	}

	// Fill the rest with random characters from allChars
	for i := 4; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		password[i] = allChars[n.Int64()]
	}

	// Shuffle the password to avoid predictable positions
	shuffle(password)

	return string(password), nil
}

// shuffle randomly reorders the bytes in the password
func shuffle(password []byte) {
	for i := len(password) - 1; i > 0; i-- {
		max := big.NewInt(int64(i + 1))
		j, _ := rand.Int(rand.Reader, max)
		password[i], password[j.Int64()] = password[j.Int64()], password[i]
	}
}
