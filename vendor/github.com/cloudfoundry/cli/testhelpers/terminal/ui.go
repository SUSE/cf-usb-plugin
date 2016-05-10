package terminal

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"os"

	"github.com/cloudfoundry/cli/cf"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	term "github.com/cloudfoundry/cli/cf/terminal"
)

const QuietPanic = "I should not print anything"

type FakeUI struct {
	Outputs                    []string
	UncapturedOutput           []string
	WarnOutputs                []string
	Prompts                    []string
	PasswordPrompts            []string
	Inputs                     []string
	FailedWithUsage            bool
	FailedWithUsageCommandName string
	PanickedQuietly            bool
	ShowConfigurationCalled    bool

	sayMutex sync.Mutex
}

func (ui *FakeUI) PrintPaginator(rows []string, err error) {
	if err != nil {
		ui.Failed(err.Error())
		return
	}

	for _, row := range rows {
		ui.Say(row)
	}
}

func (ui *FakeUI) Writer() io.Writer {
	return os.Stdout
}

func (ui *FakeUI) PrintCapturingNoOutput(message string, args ...interface{}) {
	ui.sayMutex.Lock()
	defer ui.sayMutex.Unlock()

	message = fmt.Sprintf(message, args...)
	ui.UncapturedOutput = append(ui.UncapturedOutput, strings.Split(message, "\n")...)
	return
}

func (ui *FakeUI) Say(message string, args ...interface{}) {
	ui.sayMutex.Lock()
	defer ui.sayMutex.Unlock()

	message = fmt.Sprintf(message, args...)
	ui.Outputs = append(ui.Outputs, strings.Split(message, "\n")...)
	return
}

func (ui *FakeUI) Warn(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	ui.WarnOutputs = append(ui.WarnOutputs, strings.Split(message, "\n")...)
	ui.Say(message, args...)
	return
}

func (ui *FakeUI) Ask(prompt string) string {
	ui.Prompts = append(ui.Prompts, prompt)

	if len(ui.Inputs) == 0 {
		panic("No input provided to Fake UI for prompt: " + prompt)
	}

	answer := ui.Inputs[0]
	ui.Inputs = ui.Inputs[1:]
	return answer
}

func (ui *FakeUI) ConfirmDelete(modelType, modelName string) bool {
	return ui.Confirm(fmt.Sprintf(
		"Really delete the %s %s?%s",
		modelType,
		term.EntityNameColor(modelName),
		term.PromptColor(">")))
}

func (ui *FakeUI) ConfirmDeleteWithAssociations(modelType, modelName string) bool {
	return ui.ConfirmDelete(modelType, modelName)
}

func (ui *FakeUI) Confirm(prompt string) bool {
	response := ui.Ask(prompt)
	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	}
	return false
}

func (ui *FakeUI) AskForPassword(prompt string) string {
	ui.PasswordPrompts = append(ui.PasswordPrompts, prompt)

	if len(ui.Inputs) == 0 {
		panic("No input provided to Fake UI for prompt: " + prompt)
	}

	answer := ui.Inputs[0]
	ui.Inputs = ui.Inputs[1:]
	return answer
}

func (ui *FakeUI) Ok() {
	ui.Say("OK")
}

func (ui *FakeUI) Failed(message string, args ...interface{}) {
	ui.Say("FAILED")
	ui.Say(message, args...)
	panic(QuietPanic)
}

func (ui *FakeUI) PanicQuietly() {
	ui.PanickedQuietly = true
}

func (ui *FakeUI) DumpWarnOutputs() string {
	return "****************************\n" + strings.Join(ui.WarnOutputs, "\n")
}

func (ui *FakeUI) DumpOutputs() string {
	return "****************************\n" + strings.Join(ui.Outputs, "\n")
}

func (ui *FakeUI) DumpPrompts() string {
	return "****************************\n" + strings.Join(ui.Prompts, "\n")
}

func (ui *FakeUI) ClearOutputs() {
	ui.Outputs = []string{}
}

func (ui *FakeUI) ShowConfiguration(config coreconfig.Reader) {
	ui.ShowConfigurationCalled = true
}

func (ui *FakeUI) LoadingIndication() {
}

func (ui *FakeUI) Wait(duration time.Duration) {
	time.Sleep(duration)
}

func (ui *FakeUI) Table(headers []string) *term.UITable {
	return &term.UITable{
		UI:    ui,
		Table: term.NewTable(headers),
	}
}

func (ui *FakeUI) NotifyUpdateIfNeeded(config coreconfig.Reader) {
	if !config.IsMinCLIVersion(cf.Version) {
		ui.Say("Cloud Foundry API version {{.APIVer}} requires CLI version " + config.MinCLIVersion() + "  You are currently on version {{.CLIVer}}. To upgrade your CLI, please visit: https://github.com/cloudfoundry/cli#downloads")
	}
}
