package ask

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func String(format string, args ...interface{}) (string, error) {
	return str(format, args)
}

func str(format string, args []interface{}) (string, error) {
	Clear()
	var response string
	prompt := &survey.Input{
		Message: fmt.Sprintf(format, args...),
	}
	err := survey.AskOne(prompt, &response)
	return strings.TrimSpace(response), err
}
