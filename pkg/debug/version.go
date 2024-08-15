package debug

import (
	"fmt"
	"os"

	"github.com/iaakanshff/gostyle"
)

// Function for checking the version info
func CheckVersion() {
	g := gostyle.New()
	fmt.Fprintln(os.Stderr, g.Inf(), "Current hexify version:", g.BrGreen("v1.0"))
}
