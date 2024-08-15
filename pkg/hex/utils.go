package hex

import (
	"fmt"
)

// formatHexBytes converts a byte slice into a formatted hex string, with optional highlighting.
func formatHexBytes(chunk, highlight []byte, withColor bool) []byte {
	hexBytes := make([]byte, 0, 3*len(chunk))

	for j := 0; j < len(chunk); j++ {
		if j > 0 {
			hexBytes = append(hexBytes, ' ')
		}

		if matchBytes(chunk[j:], highlight) {
			if withColor {
				hexBytes = append(hexBytes, RedBold...)
			}
			
			for k := 0; k < len(highlight); k++ {
				hexBytes = append(hexBytes, fmt.Sprintf("%02X", chunk[j+k])...)
				if k < len(highlight)-1 {
					hexBytes = append(hexBytes, ' ')
				}
			}
			if withColor {
				hexBytes = append(hexBytes, Reset...)
			}
			j += len(highlight) - 1
		} else {
			hexBytes = append(hexBytes, fmt.Sprintf("%02X", chunk[j])...)
		}
	}

	return hexBytes
}

// formatAscii converts a byte slice into a formatted ASCII string.
func formatAscii(chunk []byte) []byte {
	ascii := make([]byte, len(chunk))
	for j, b := range chunk {
		if b >= 32 && b <= 126 {
			ascii[j] = b
		} else {
			ascii[j] = '.'
		}
	}
	return ascii
}

// matchBytes checks if the start of a slice matches a given pattern.
func matchBytes(data, pattern []byte) bool {
	if len(data) < len(pattern) {
		return false
	}
	for i := range pattern {
		if data[i] != pattern[i] {
			return false
		}
	}
	return true
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
