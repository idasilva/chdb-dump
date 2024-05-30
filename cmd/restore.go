package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type releaseOpts struct {
}

type restoreCmd struct {
	cmd *cobra.Command
}

func (r *restoreCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRestoreCmd() *restoreCmd {
	restore := &restoreCmd{}
	cmd := &cobra.Command{
		Use:           "restore",
		Aliases:       []string{"r"},
		Short:         "restore couchdb database",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("retore")

			return nil
		},
	}

	restore.cmd = cmd
	return restore
}
