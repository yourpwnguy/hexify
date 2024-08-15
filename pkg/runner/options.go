package runner

import (
	"flag"
	"os"

	"github.com/yourpwnguy/hexify/pkg/debug"
)

// Struct for holding the args
type Options struct {
	FileData   []byte // Representing the file data
	Lines      int    // Representing the no of lines to print
	Width      int    // Representing total hex bytes to print in one line
	Highlight  []byte // Representing which pattern to highlight
	OutputFile string // Representing the output file
}

// Parsing the command line flags
func ParseOptions() *Options {

	options := &Options{}

	// Flags for command line arguments
	filename := flag.String("f", "", "")
	hexStr := flag.String("ha", "MZ", "")
	vers1 := flag.Bool("v", false, "")
	vers2 := flag.Bool("version", false, "")
	flag.IntVar(&options.Lines, "l", 0, "")
	flag.IntVar(&options.Width, "w", 16, "")
	flag.StringVar(&options.OutputFile, "o", "", "")

	// Setting the Usage function to a custom one
	flag.Usage = debug.Usage
	flag.Parse()

	// Checking if no flags are given
	if flag.NArg() == 0 && flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// for Checking version info
	if *vers1 || *vers2 {
		debug.CheckVersion()
		os.Exit(0)
	}

	// Checking if the file given exists or not
	options.FileData = ValidateFile(filename)

	// Convert the given pattern to equivalent byte array
	options.Highlight = asciiToHex(*hexStr)

	return options
}
