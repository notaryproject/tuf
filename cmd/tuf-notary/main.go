package main

import (
	"fmt"
	"log"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `
Usage:
  tuf-notary <command> [<args>....]
  tuf-notary <command> [<args>....] [--repo=<repository>]

Commands:
  help          Show usage for a specific command
  init			Initialize a TUF repository
	`

	args, _ := docopt.ParseDoc(usage)
	cmd := args["<command>"].(string)
	cmdArgs := args["<args>"].([]string)

	if cmd == "help" {
		if len(cmdArgs) == 0 { // 'tuf-notary help'
			fmt.Println(usage)
			return
		} else { // `tuf-notary help <command>`
			cmd = cmdArgs[0]
			cmdArgs = []string{"--help"}
		}
	}

	if err := runCommand(cmd, cmdArgs, args); err != nil {
		log.Fatalln("ERROR:", err)
	}
}

type cmdFunc func([]string, docopt.Opts) error

type command struct {
	usage string
	f     cmdFunc
}

var commands = make(map[string]*command)

func register(name string, f cmdFunc, usage string) {
	commands[name] = &command{usage: usage, f: f}
}

func runCommand(name string, args []string, opts docopt.Opts) error {
	argv := make([]string, 1, 1+len(args))
	argv[0] = name
	argv = append(argv, args...)

	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("%s is not a tuf-notary command. See 'tuf-notary help'", name)
	}

	_, err := docopt.ParseDoc(cmd.usage)
	if err != nil {
		return err
	}

	return cmd.f(args, opts)
}
