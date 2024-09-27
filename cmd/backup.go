package cmd

import (
	"fmt"
	"os"

	"github.com/idasilva/chdb-dump/pkg/context"
	"github.com/idasilva/chdb-dump/pkg/dump"
	"github.com/idasilva/chdb-dump/pkg/logger"
	"github.com/spf13/cobra"
)

type backupOpts struct {
	persistence string
	database    []string
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
	logger := logger.New()

	backup := &backupCmd{}
	cmd := &cobra.Command{
		Use:           "backup",
		Aliases:       []string{"b"},
		Short:         "backup couchdb database",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			dump, err := dump.NewBackup(&context.Context{
				Persistence: backup.opts.persistence,
				Database:    backup.opts.database,
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

	cmd.PersistentFlags().StringVarP(&backup.opts.persistence,
		"persistence", "p", "local", "where you need to pull your data")

	cmd.PersistentFlags().StringSliceVarP(&backup.opts.database,
		"database", "d", []string{}, "which database you need to backup")

	backup.cmd = cmd
	return backup
}
