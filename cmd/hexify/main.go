package main

import (
	"github.com/yourpwnguy/hexify/pkg/hex"
	"github.com/yourpwnguy/hexify/pkg/runner"
)

func main() {

	// Get all the args after parsing
	options := runner.ParseOptions()

	// Call the printing function
	hex.HexPrint(options)
}
