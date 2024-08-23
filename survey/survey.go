package survey

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/gofiber/fiber/v2/log"
)

func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	err := survey.AskOne(prompt, &res)
	if err != nil {
		if err == terminal.InterruptErr {
			log.Fatal("Interrupted")
		}
	}

	return res
}

func SingleSelect(label string, opts []string) string {
	res := ""
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	err := survey.AskOne(prompt, &res)
	if err != nil {
		if err == terminal.InterruptErr {
			log.Fatal("Interrupted")
		}
	}

	return res
}

func StringPrompt(label string) string {
	res := ""
	prompt := &survey.Input{
		Message: label,
	}
	err := survey.AskOne(prompt, &res)
	if err != nil {
		if err == terminal.InterruptErr {
			log.Fatal("Interrupted")
		}
	}

	return res
}
