package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select [server]",
	Short: "Select a remote server as context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		server := args[0]

		home, _ := os.UserHomeDir()
		path := filepath.Join(home, ".sshm", "context")

		os.MkdirAll(filepath.Dir(path), 0700)
		os.WriteFile(path, []byte(server), 0600)

		fmt.Printf("[sshm] context set to: %s\n", server)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
