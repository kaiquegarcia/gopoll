package ask

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Password(format string, args ...interface{}) (string, error) {
	Clear()
	var response string
	prompt := &survey.Password{
		Message: fmt.Sprintf(format, args...),
	}
	err := survey.AskOne(prompt, &response)
	return response, err
}
