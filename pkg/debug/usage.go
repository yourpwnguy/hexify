package debug

import (
	"fmt"
	"os"
)

// Usage displays the command-line usage information for the hexify tool.
func Usage() {
	h := "\nUsage: hexify [options]\n\n"
	h += "Options: [flag] [argument] [Description]\n\n"
	h += "BASIC:\n"
	h += "  -f string\t\tSpecify the input file to read data from.\n"
	h += "  -o string\t\tSpecify the output file to write the hex dump to\n\n"
	h += "PROCESSING:\n"
	h += "  -l int\t\tSpecify the number of lines to print (default: all lines).\n"
	h += "  -w int\t\tSpecify the number of hex bytes to print per line (default: 16).\n"
	h += "  -ha string\t\tHighlight a specific ASCII string pattern (default: 'MZ').\n\n"
	h += "DEBUG:\n"
	h += "  -h, --help\t\tDisplay this help message.\n"
	h += "  -v, --version\t\tCheck current version"

	fmt.Fprintln(os.Stderr, h)
}
