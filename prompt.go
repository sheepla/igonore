package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Prompt(label string) error {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("an error occured on prompt: %w", err)
	}

	return nil
}
