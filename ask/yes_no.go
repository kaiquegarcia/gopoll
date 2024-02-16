package ask

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func YesNo(format string, args ...interface{}) (bool, error) {
	Clear()
	var response bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf(format, args...),
	}
	err := survey.AskOne(prompt, &response)
	return response, err
}
