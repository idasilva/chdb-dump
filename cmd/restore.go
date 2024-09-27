package cmd

import (
	"fmt"
	"os"

	"github.com/idasilva/chdb-dump/pkg/context"
	"github.com/idasilva/chdb-dump/pkg/dump"
	"github.com/idasilva/chdb-dump/pkg/logger"
	"github.com/spf13/cobra"
)

type restoreOpts struct {
	persistence string
}

type restoreCmd struct {
	cmd  *cobra.Command
	opts restoreOpts
}

func (r *restoreCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRestoreCmd() *restoreCmd {
	logger := logger.New()

	restore := &restoreCmd{}
	cmd := &cobra.Command{
		Use:           "restore",
		Aliases:       []string{"r"},
		Short:         "restore couchdb database",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			dump, err := dump.NewRestore(&context.Context{
				Persistence: restore.opts.persistence,
			})

			if err != nil {
				logger.Fatal(err.Error())
			}

			err = dump.Exec()
			if err != nil {
				logger.Fatal(err.Error())
			}

			return nil
		},
	}

	restore.cmd = cmd
	return restore
}
