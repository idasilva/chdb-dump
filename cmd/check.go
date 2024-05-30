package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type checkCmd struct {
	cmd *cobra.Command
}

func (r *checkCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newCheckCmd() *checkCmd {
	check := &checkCmd{}
	cmd := &cobra.Command{
		Use:           "check",
		Aliases:       []string{"c"},
		Short:         "Checks the environment",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("check")

			return nil
		},
	}

	check.cmd = cmd
	return check
}
