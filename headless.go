//go:build headless

package main

import (
	"fmt"
	"os"
)

func startGUI(app *App) {
	fmt.Fprintln(os.Stderr, "Error: This binary was compiled in headless mode and does not support the GUI.")
	fmt.Fprintln(os.Stderr, "Please provide arguments to run in CLI mode. Run with --help for usage.")
	os.Exit(1)
}
