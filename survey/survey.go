package survey

import (
	"fmt"
	"os"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/term"
)

func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	err := survey.AskOne(prompt, &res, survey.WithIcons(func(icons *survey.IconSet) { icons.Question.Text = "<>" }))
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
	err := survey.AskOne(prompt, &res, survey.WithIcons(func(icons *survey.IconSet) { icons.Question.Text = "<>" }))
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
	err := survey.AskOne(prompt, &res, survey.WithIcons(func(icons *survey.IconSet) { icons.Question.Text = "<>" }))
	if err != nil {
		if err == terminal.InterruptErr {
			log.Fatal("Interrupted")
		}
	}

	return res
}

func PasswordPrompt(label string) string {
	s := ""
	for {
		fmt.Fprint(os.Stderr, "<> "+label+" ")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(b)
		if s != "" {
			break
		}
	}
	fmt.Println()
	return s
}
