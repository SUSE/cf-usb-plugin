package trace

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/gofileutils/fileutils"
)

type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type nullLogger struct{}

func (*nullLogger) Print(v ...interface{})                 {}
func (*nullLogger) Printf(format string, v ...interface{}) {}
func (*nullLogger) Println(v ...interface{})               {}

var stdOut io.Writer = os.Stdout
var Logger Printer

func init() {
	Logger = NewLogger("")
}

func EnableTrace() {
	Logger = newStdoutLogger()
}

func DisableTrace() {
	Logger = new(nullLogger)
}

func SetStdout(s io.Writer) {
	stdOut = s
}

func NewLogger(cf_trace string) Printer {
	switch cf_trace {
	case "", "false":
		Logger = new(nullLogger)
	case "true":
		Logger = newStdoutLogger()
	default:
		Logger = newFileLogger(cf_trace)
	}

	return Logger
}

func newStdoutLogger() Printer {
	return log.New(stdOut, "", 0)
}

func newFileLogger(path string) Printer {
	file, err := fileutils.Open(path)
	if err != nil {
		logger := newStdoutLogger()
		logger.Printf(T("CF_TRACE ERROR CREATING LOG FILE {{.Path}}:\n{{.Err}}",
			map[string]interface{}{"Path": path, "Err": err}))
		return logger
	}

	return log.New(file, "", 0)
}

func Sanitize(input string) (sanitized string) {
	var sanitizeJson = func(propertyName string, json string) string {
		regex := regexp.MustCompile(fmt.Sprintf(`"%s":\s*"[^\,]*"`, propertyName))
		return regex.ReplaceAllString(json, fmt.Sprintf(`"%s":"%s"`, propertyName, PRIVATE_DATA_PLACEHOLDER()))
	}

	re := regexp.MustCompile(`(?m)^Authorization: .*`)
	sanitized = re.ReplaceAllString(input, "Authorization: "+PRIVATE_DATA_PLACEHOLDER())
	re = regexp.MustCompile(`password=[^&]*&`)
	sanitized = re.ReplaceAllString(sanitized, "password="+PRIVATE_DATA_PLACEHOLDER()+"&")

	sanitized = sanitizeJson("access_token", sanitized)
	sanitized = sanitizeJson("refresh_token", sanitized)
	sanitized = sanitizeJson("token", sanitized)
	sanitized = sanitizeJson("password", sanitized)
	sanitized = sanitizeJson("oldPassword", sanitized)

	return
}

func PRIVATE_DATA_PLACEHOLDER() string {
	return T("[PRIVATE DATA HIDDEN]")
}
