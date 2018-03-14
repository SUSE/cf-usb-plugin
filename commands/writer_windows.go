// +build windows

package commands

import "github.com/fatih/color"

// Writer that will be used to display output.  A wrapper is needed on Windows
// for colour support.
var Writer = color.Output
