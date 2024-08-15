package logger

import (
	"fmt"
	"os"

	"github.com/yourpwnguy/gostyle"
)

var g = gostyle.New()

// Function for displaying error messages with [ERR] Prefix
func LogErr(input string) {
	fmt.Fprintln(os.Stderr, g.Err(), input)
}
