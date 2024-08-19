package survey

import "github.com/AlecAivazis/survey/v2"

func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func SingleSelect(label string, opts []string) string {
	res := ""
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func StringPrompt(label string) string {
	res := ""
	prompt := &survey.Input{
		Message: label,
	}
	survey.AskOne(prompt, &res)

	return res
}
