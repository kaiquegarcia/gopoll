package ask

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func ChooseMultipleOptions(
	labels []string,
	format string,
	args ...interface{},
) ([]int, error) {
	Clear()
	var responses []string
	prompt := &survey.MultiSelect{
		Message: fmt.Sprintf(format, args...),
		Options: labels,
	}
	err := survey.AskOne(prompt, &responses)
	if err != nil {
		return nil, err
	}

	indexes := make([]int, len(responses))
	for i, response := range responses {
		for index, label := range labels {
			if label == response {
				indexes[i] = index
				break
			}
		}
	}

	return indexes, nil
}
