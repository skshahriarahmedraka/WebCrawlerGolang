package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/skshahriarahmedraka/WebCrawlerGolang/cmd/crawler"
	commander "github.com/skshahriarahmedraka/WebCrawlerGolang/pkg/commands"
	"github.com/spf13/cobra"
)

type Command struct {
	commands []commander.Commander
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Register() {
	c.commands = append(c.commands, crawler.NewCrawlerCommand())
}

func (c *Command) Commands() []*cobra.Command {
	var commands []*cobra.Command
	for _, command := range c.commands {

		commands = append(commands, c.prepareCommand( command))
	}
	return commands
}

func (c *Command) prepareCommand (cmd commander.Commander) *cobra.Command {
	cobraCmd := &cobra.Command{
		Use:   cmd.Signature(),
		Short: cmd.Short(),
		Long:  cmd.Long(),
		Run: func(c *cobra.Command, args []string) {
			ctx := context.Background()
			if err := cmd.Run(ctx, args); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}
	cmd.Setup(cobraCmd)
	return cobraCmd
}
