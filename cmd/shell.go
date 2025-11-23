package cmd

import (
    "github.com/spf13/cobra"
	
)

var selectedHost string

func init() {
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringVarP(&selectedHost, "host", "H", "", "Host alias from ~/.ssh/config")
	shellCmd.MarkFlagRequired("host")
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open a remote execution shell for the selected server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startRemoteShell(selectedHost)
	},
}

func startRemoteShell(host string) error {
	// TODO: Implement remote shell logic
	return nil
}
