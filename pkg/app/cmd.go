// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Command is a sub command structure of a cli application.
// It is recommended that a command be created with the app.NewCommand()
// function.
type Command struct {
	usage    string
	desc     string
	options  CliOptions
	commands []*Command
	runFunc  RunCommandFunc
}

// CommandOption defines optional parameters for initializing the command
// structure.
type CommandOption func(*Command)

// WithCommandOptions to open the application's function to read from the
// command line.
func WithCommandOptions(opt CliOptions) CommandOption {
	return func(c *Command) {
		c.options = opt
	}
}

// RunCommandFunc defines the application's command startup callback function.
type RunCommandFunc func(args []string) error

// WithCommandRunFunc is used to set the application's command startup callback
// function option.
func WithCommandRunFunc(run RunCommandFunc) CommandOption {
	return func(c *Command) {
		c.runFunc = run
	}
}

// NewCommand creates a new sub command instance based on the given command name
// and other options.
func NewCommand(usage string, desc string, opts ...CommandOption) *Command {
	c := &Command{
		usage: usage,
		desc:  desc,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// AddCommand adds sub command to the current command.
func (c *Command) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// AddCommands adds multiple sub commands to the current command.
func (c *Command) AddCommands(cmds ...*Command) {
	c.commands = append(c.commands, cmds...)
}

func (c *Command) cobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   c.usage,
		Short: c.desc,
	}
	cmd.SetOutput(os.Stdout)
	cmd.Flags().SortFlags = false
	if len(c.commands) > 0 {
		for _, command := range c.commands {
			cmd.AddCommand(command.cobraCommand())
		}
	}
	if c.runFunc != nil {
		cmd.Run = c.runCommand
	}
	if c.options != nil {
		for _, f := range c.options.Flags().FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
		// c.options.AddFlags(cmd.Flags())
	}
	addHelpCommandFlag(c.usage, cmd.Flags())

	return cmd
}

func (c *Command) runCommand(cmd *cobra.Command, args []string) {
	if c.runFunc != nil {
		if err := c.runFunc(args); err != nil {
			fmt.Printf("%v %v\n", color.RedString("Error:"), err)
			os.Exit(1)
		}
	}
}

// AddCommand adds sub command to the application.
func (a *App) AddCommand(cmd *Command) {
	a.commands = append(a.commands, cmd)
}

// AddCommands adds multiple sub commands to the application.
func (a *App) AddCommands(cmds ...*Command) {
	a.commands = append(a.commands, cmds...)
}

// FormatBaseName is formatted as an executable file name under different
// operating systems according to the given name.
func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}

/*
This Go code is a part of a CLI (Command Line Interface) application. It introduces a new concept: commands, which are sub-structures of a CLI application. Each command has a name (usage), description (desc), options, potentially sub-commands, and a run function.

Commands in CLI applications can be considered as specific tasks or functions that the application can perform. They are a way to structure a program and its functionalities.

Key functions and types in this code are:

1. `Command`: A structure that represents a command in a CLI application.

2. `CommandOption`: A function type used to apply optional settings to a command.

3. `WithCommandOptions` and `WithCommandRunFunc`: Functions to set options and the run function for a command.

4. `NewCommand`: A function to create a new command.

5. `AddCommand` and `AddCommands`: Methods of the `Command` and `App` structure that allow adding one or multiple sub-commands to the current command or application.

6. `cobraCommand`: A method of the `Command` structure that builds a cobra command (from the Cobra library, a widely used library for creating CLI applications in Go) from the command.

7. `runCommand`: A method of the `Command` structure that is used as the run function for the built cobra command.

8. `FormatBaseName`: A function that formats the basename (the name of the binary file of the application) according to the operating system.

9. `RunCommandFunc`: A type that defines a function which will be called when the command is executed.

This part of the code adds more structure and flexibility to the CLI application, allowing it to have multiple commands with their own options and run functions. This is a common design in many complex CLI applications, where different commands represent different functionalities of the application.
*/
