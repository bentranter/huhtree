# `huhtree`

Combining [`cobra`](https://github.com/spf13/cobra) and [`huh`](https://github.com/charmbracelet/huh) for the ultimate CLI experience.

## Usage

With [Go](https://go.dev/dl/) installed, run the following command to compile the CLI and setup autocompletion.

```sh
# For bash
$ go build -o huhtree main.go && ./huhtree completion bash > /tmp/completion && source /tmp/completion

# For zsh
$ go build -o huhtree main.go && ./huhtree completion zsh > /tmp/completion && source /tmp/completion
```

Once it's compiled, you can run the binary:

```sh
$ ./huhtree
```
