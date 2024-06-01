package cmd

import (
	"fmt"
	"os"

	"github.com/idasilva/chdb-dump/pkg/context"
	"github.com/idasilva/chdb-dump/pkg/dump"
	"github.com/spf13/cobra"
)

type backupOpts struct {
	persistence string
}

type backupCmd struct {
	cmd  *cobra.Command
	opts backupOpts
}

func (r *backupCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newBackupCmd() *backupCmd {
	backup := &backupCmd{}
	cmd := &cobra.Command{
		Use:           "backup",
		Aliases:       []string{"c"},
		Short:         "backup couchdb database",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			dump, err := dump.New(&context.Context{
				Persistence: backup.opts.persistence,
			})

			if err != nil {
				return err
			}

			err = dump.Exec()
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&backup.opts.persistence,
		"persistence", "p", "local", "where you need to pull your data")

	backup.cmd = cmd
	return backup
}
