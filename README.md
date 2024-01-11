# `huhtree`

Combining [`cobra`](https://github.com/spf13/cobra) and [`huh`](https://github.com/charmbracelet/huh) for the ultimate CLI experience.

## Demo

[![asciicast](https://asciinema.org/a/631028.svg)](https://asciinema.org/a/631028)

## Usage

### Building

With [Go](https://go.dev/dl/) installed, run the following command to compile the CLI and setup autocompletion.

```sh
# For bash
$ go build -o huhtree main.go && ./huhtree completion bash > /tmp/completion && source /tmp/completion

# For zsh
$ go build -o huhtree main.go && ./huhtree completion zsh > /tmp/completion && source /tmp/completion
```

### Running

Once it's compiled, you can run the binary:

```sh
$ ./huhtree
```

Doing so with no arguments will ask you which command you wish to run from a list of options.

If you want to see the normal help command, you can still run:

```
$ ./huhtree -h
```

### Command Autocomplete

The goal of this demo is to demonstrate how discoverable a CLI's different commands can be.

From the root command, there are three ways to discover the commands, with two being runnable:

1. Output the help info of the root command.
2. Run the root command with no arguments to choose from a list of subcommands.
3. Hit <Tab> after the CLI name to view and select from a list of subcommands.

When running a subcommand, there are again three ways to discover the options, with two bring runnable:

1. Output the help info for the subcommands.
2. Run the subcommand with no arguments to be prompted to fill out a "form" for each required argument.
3. Hit <Tab> after the subcommand's name in the CLI to view and select from a list of suggested arguments.
