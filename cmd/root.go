package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sshm",
	Short: "A brief description of your application",
	Long: `sshm is a modern SSH management CLI designed to simplify working with multiple servers.

It automatically reads your ~/.ssh/config file, provides intelligent tab-completion, 
and lets you connect or execute remote commands seamlessly.

Features:
  • Auto-detect and list SSH hosts from ~/.ssh/config
  • Fast connection using alias-based SSH profiles
  • Execute commands remotely without interactive login (sshm shell)
  • Friendly output and intuitive UX, inspired by kubectx + kubectl exec
  • Zsh/Fish/Bash completion support
	
Examples:
  # Connect to a host from ~/.ssh/config
  sshm connect prod-server

  # Open remote execution shell
  sshm shell --host dev-db

  # List available servers with auto-completion
  sshm connect <TAB>

sshm helps you work faster, safer, and more efficiently with SSH environments.`,
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.CompletionOptions.DisableDefaultCmd = false

}


