package cmd

import (
	"fmt"
	"os"

	"github.com/idasilva/chdb-dump/pkg/dump"
	"github.com/spf13/cobra"
)

type backupOpts struct {
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

			bkp, err := dump.New(backup.opts)
			if err != nil {
				return err
			}

			bkp.Store()

			return nil
		},
	}

	backup.cmd = cmd
	return backup
}
