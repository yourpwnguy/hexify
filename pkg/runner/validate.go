package runner

import (
	"os"
	"strings"

	"github.com/yourpwnguy/hexify/pkg/logger"
)

// Function for validating the file given
func ValidateFile(f *string) []byte {

	// Checking if file string is empty
	if strings.TrimSpace(*f) == "" {
		logger.LogErr("Invalid filename")
		os.Exit(0)
	}

	// Read data from file
	data, err := os.ReadFile(*f)
	if err != nil {
		logger.LogErr("Error reading file")
		os.Exit(0)
	}

	return data
}

// Function to convert the given string to byte array
func asciiToHex(ascii string) []byte {

	// if ascii is 00, then we return null
	if ascii == "00" {
		return []byte{0x00}
	}

	// Otherwise we perform conversion
	var hexBytes []byte
	for i := 0; i < len(ascii); i++ {
		hexBytes = append(hexBytes, ascii[i])
	}
	return hexBytes
}
