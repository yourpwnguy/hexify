package hex

import (
	"fmt"
	"os"
	"strings"

	"github.com/iaakanshff/gostyle"
	"github.com/yourpwnguy/hexify/pkg/logger"
	"github.com/yourpwnguy/hexify/pkg/runner"
)

// ANSI Codes for Coloring
const (
	Yellow  = "\033[0;33m"
	RedBold = "\033[1;31m" // Bold red
	Reset   = "\033[0m"    // Reset formatting
)

// The styling library
var g = gostyle.New()

// HexPrint displays the hex dump of the file, optionally outputting to a file.
func HexPrint(options *runner.Options) {

	// Determine the number of lines to print
	totalLines := (len(options.FileData) + options.Width - 1) / options.Width
	if options.Lines == 0 || options.Lines > totalLines {
		options.Lines = totalLines
	}

	// Variables for later use
	var (
		// Output related
		output *os.File
		err    error

		// For padding related
		padded string
		spaces = options.Width - 1
		bytes  = ((options.Width)*2 + spaces)
	)

	// Open the file if given otherwise set the stream to stderr
	if options.OutputFile != "" {
		output, err = os.Create(options.OutputFile)
		if err != nil {
			logger.LogErr(fmt.Sprintf("Error creating output file: %v\n", err))
			return
		}
		defer output.Close()
	} else {
		output = os.Stderr
	}

	// Loop over the file data and process chunks of `Width` bytes at a time
	for i := 0; i < len(options.FileData) && options.Lines > 0; i += options.Width {
		// Loop over the file data and process chunks of `Width` bytes at a time
		chunk := options.FileData[i:min(i+options.Width, len(options.FileData))]

		// 2nd Part: Format the hex bytes
		hexBytes := formatHexBytes(chunk, options.Highlight, options.OutputFile == "")

		// 3rd Part: Format the ASCII bytes
		ascii := formatAscii(chunk)

		// Calculating the padding needed for each chunk
		needed := bytes - (len(chunk)*2 + len(chunk) - 1)
		padded = strings.Repeat(" ", needed)

		
		// Either save to a file without any ANSI or Print it to the console with coloring
		if options.OutputFile != "" {
			// Write to the file without color formatting
			fmt.Fprintf(output, "0x%08x   %s    %s%s\n", i, string(hexBytes), padded, string(ascii))

		} else {
			fmt.Fprintf(output, "%s0x%08x%s   %s    %s%s\n",
				Yellow,
				i,
				Reset,
				string(hexBytes),
				padded,
				g.Bold(g.BrCyan(string(ascii))),
			)
		}
		options.Lines--
	}
}
