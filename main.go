package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	flags := &struct {
		next string
	}{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Which command do you want to run?").
				Options(
					huh.NewOption("Input example", "input"),
					huh.NewOption("Confirm example", "confirm"),
				).
				Value(&flags.next),
		),
	).WithTheme(huh.ThemeBase())

	run := func(cmd *cobra.Command) error {
		if err := form.Run(); err != nil {
			return err
		}

		for _, c := range cmd.Commands() {
			if name := c.Name(); name == flags.next {
				if err := c.RunE(c, []string{}); err != nil {
					return err
				}
				return nil
			}
		}

		return errors.New("command " + flags.next + " not found")
	}

	return &cobra.Command{
		Use:   "huhtree", // Must compile as 'huhtree' for autocompletion to work.
		Short: "Trying out cobra and huh for the ultimate CLI experience",
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Println("TODO output default help")
			return run(cmd)
		},
	}
}

func inputCmd() *cobra.Command {
	flags := &struct {
		name string
	}{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("What's your name?").Value(&flags.name),
		),
	).WithTheme(huh.ThemeBase())

	run := func() error {
		fmt.Println("Your name is", flags.name)
		return nil
	}

	return &cobra.Command{
		Use:   "input <name>",
		Short: "Input your name and write it to stdout",
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				if err := form.Run(); err != nil {
					return err
				}
			} else {
				flags.name = args[0]
			}
			return run()
		},
		ValidArgsFunction: func(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			options := []string{"Ben", "Bob", "Tato", "Example"}
			suggestions := make([]string, 0, len(options))

			if toComplete != "" {
				for _, o := range options {
					if strings.HasPrefix(o, toComplete) {
						suggestions = append(suggestions, o)
					}
				}
			} else {
				suggestions = options
			}

			return suggestions, cobra.ShellCompDirectiveNoFileComp
		},
	}
}

func confirmCmd() *cobra.Command {
	flags := &struct {
		choice bool
	}{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Confirm yes or no").Value(&flags.choice),
		),
	).WithTheme(huh.ThemeBase())

	run := func() error {
		fmt.Println("You chose", flags.choice)
		return nil
	}

	return &cobra.Command{
		Use:   "confirm",
		Short: "Confirm something yes or no",
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				if err := form.Run(); err != nil {
					return err
				}
			} else {
				v, err := strconv.ParseBool(args[0])
				if err != nil {
					return err
				}
				flags.choice = v
			}
			return run()
		},
		ValidArgsFunction: func(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			options := []string{"true", "false"}
			suggestions := make([]string, 0, len(options))

			if toComplete != "" {
				for _, o := range options {
					if strings.HasPrefix(o, toComplete) {
						suggestions = append(suggestions, o)
					}
				}
			} else {
				suggestions = options
			}

			return suggestions, cobra.ShellCompDirectiveNoFileComp
		},
	}
}

func main() {
	root := rootCmd()
	root.AddCommand(inputCmd())
	root.AddCommand(confirmCmd())
	cobra.CheckErr(root.Execute())
}
