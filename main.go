package main

import (
	"fmt"
	"io"
	"os"

	"github.com/skshahriarahmedraka/WebCrawlerGolang/cmd"
	"github.com/spf13/cobra"
)

const (
	exitFail = 1
)


func main() {
	if err := run(os.Args,os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}
}


func run(args []string , stdout io.Writer) error {
	rootCmd := &cobra.Command{Use: "web crawler service"}

	command := cmd.NewCommand()
	command.Register()

	rootCmd.AddCommand(command.Commands()...)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of WebCrawlerGolang",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(stdout, "WebCrawlerGolang v0.1")
		},
	})



	if len(args) < 2 {
		return fmt.Errorf("usage: %s <command> [<args>]", args[0])
	}
	
    fmt.Println("🚀 ~ file: main.go ~ line 50 ~ funcrun ~ args[1:] : ", args[1:])
	rootCmd.SetArgs(args[1:])

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil 
}