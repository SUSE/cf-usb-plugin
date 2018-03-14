// +build !windows

package commands

import "os"

// Writer that will be used to display output.  This is just the normal one on
// *nix platforms.
var Writer = os.Stdout
