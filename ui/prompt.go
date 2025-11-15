package ui

import (
	"github.com/AlecAivazis/survey/v2"
)

func AskSelect(message string, options []string) (string, error) {
	var answer string
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	err := survey.AskOne(prompt, &answer)
	return answer, err
}

func AskConfirm(message string) (bool, error) {
	var ok bool
	prompt := &survey.Confirm{
		Message: message,
		Default: false,
	}
	err := survey.AskOne(prompt, &ok)
	return ok, err
}