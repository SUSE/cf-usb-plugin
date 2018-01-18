package commands

// This file contains the logging / output-related functions.
import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/SUSE/termui"
	"github.com/SUSE/termui/sigint"
	"github.com/SUSE/termui/termpassword"
)

// writerImpl is the type the implements the Writer interface
type writerImpl struct {
	*termui.UI
	sigHandler *sigint.Handler
}

var writer *writerImpl

func init() {
	var actualWriter io.Writer
	actualWriter = os.Stdout

	traceEnv := os.Getenv("CF_TRACE")
	switch strings.ToLower(traceEnv) {
	case "-", "1", "y", "yes", "t", "true":
		// Extra logging to stdout, which we don't have
	case "", "0", "n", "no", "f", "false":
		// Disable the extra logging that we don't have
	default:
		// Log to file
		log, err := os.OpenFile(traceEnv, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening log file: %s\n", err)
		} else {
			actualWriter = io.MultiWriter(os.Stdout, log)
		}
	}

	writer = &writerImpl{
		UI:         termui.New(os.Stdin, actualWriter, termpassword.NewReader()),
		sigHandler: sigint.NewHandler(),
	}
}

// ShowFailed lets the user know some action failed, and exits the application
func ShowFailed(message ...interface{}) {
	writer.Println(message...)
	writer.sigHandler.Exit(1)
}

// ShowOK lets the user know some action succeeded
func ShowOK(message string) {
	writer.Println(message)
}

// Confirm the user wishes to undergo some action using the given message
func Confirm(message string) bool {
	result := strings.ToLower(writer.Prompt(message))
	return result == "y" || result == "yes"
}
