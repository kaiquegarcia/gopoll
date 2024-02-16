package ask

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func ChooseOption(
	labels []string,
	format string,
	args ...interface{},
) (int, error) {
	Clear()
	var response string
	prompt := &survey.Select{
		Message: fmt.Sprintf(format, args...),
		Options: labels,
	}
	err := survey.AskOne(prompt, &response)
	if err != nil {
		return 0, err
	}

	selectedIndex := 0
	for index, label := range labels {
		if label == response {
			selectedIndex = index
			break
		}
	}

	return selectedIndex, nil
}
