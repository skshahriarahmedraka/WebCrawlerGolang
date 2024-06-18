package commands

import (
	"context"

	"github.com/spf13/cobra"
)


type Commander interface {
	Signature() string
	Short() string
	Long() string
	Run(context.Context, []string) error
	Setup(cmd *cobra.Command)
}
