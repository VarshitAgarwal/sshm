package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/VarshitAgarwal/sshm/internal/sshparser"
)

func init() {
	rootCmd.AddCommand(connectCmd)

	// register host completion
	connectCmd.ValidArgsFunction = hostCompletion
}

var connectCmd = &cobra.Command{
	Use:   "connect [hostname]",
	Short: "Connect to an SSH host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		fmt.Printf("Connecting to %s ...\n", host)
	},
}

// Dynamic host autocompletion
func hostCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	hosts, err := sshparser.GetSSHHosts()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	return hosts, cobra.ShellCompDirectiveNoFileComp
}
