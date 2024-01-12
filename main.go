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

	form := func(cmd *cobra.Command) *huh.Form {
		// Iterate over the cmd, getting each subcommand name and "short"
		// description. By default, on the parent command, there is the built-in
		// "help" and "completion" commands. The "help" command has no arguments,
		// so it can run fine. The "completion" requires arguments and doesn't
		// really make sense to include, so we filter it out.
		opts := make([]huh.Option[string], 0)
		for _, c := range cmd.Commands() {
			if c.Name() != "completion" {
				opts = append(opts, huh.NewOption[string](c.Name() + " -- " + c.Short, c.Name()))
			}
		}

		return huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Which subcommand do you want to run?").
					Options(opts...).
					Value(&flags.next),
			),
		).WithTheme(huh.ThemeBase())
	}

	run := func(cmd *cobra.Command) error {
		if err := form(cmd).Run(); err != nil {
			return err
		}

		for _, c := range cmd.Commands() {
			if name := c.Name(); name == flags.next {
				if c.Run != nil {
					c.Run(c, []string{})
					return nil
				}
				if c.RunE != nil {
					if err := c.RunE(c, []string{}); err != nil {
						return err
					}
					return nil
				}
			}
		}

		return errors.New("command " + flags.next + " not found")
	}

	return &cobra.Command{
		Use:   "huhtree", // Must compile as 'huhtree' for autocompletion to work.
		Short: "Trying out cobra and huh for the ultimate CLI experience",
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Is it better to also show the default help message here? The issue is
			// that a regular printf writes to stdout and therefore persists beyond the
			// invocation of the huh form.
			return run(cmd)
		},
	}
}

func inputCmd() *cobra.Command {
	flags := &struct {
		name string
	}{}

	options := []string{"Ben", "Bob", "Tato", "Example"}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("What's your name?").Value(&flags.name).Suggestions(options),
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
