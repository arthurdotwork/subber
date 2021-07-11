package service

import "github.com/manifoldco/promptui"

type ValidatorFunc func(value string) error

func NewPrompt(label string, validatorFunc ValidatorFunc) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: promptui.ValidateFunc(validatorFunc),
	}

	value, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return value, nil
}
