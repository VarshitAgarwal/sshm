package sshclient

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/kevinburke/ssh_config"
)

func LoadSSHConfig(alias string) (string, string, string, error) {
	usr, _ := user.Current()
	cfg, err := ssh_config.Open(usr.HomeDir + "/.ssh/config")
	if err != nil {
		return "", "", "", err
	}

	host := cfg.Get(alias, "HostName")
	userStr := cfg.Get(alias, "User")
	keyPath := cfg.Get(alias, "IdentityFile")

	if host == "" {
		return "", "", "", fmt.Errorf("host %s not found in SSH config", alias)
	}

	return fmt.Sprintf("%s:22", host), userStr, keyPath, nil
}
