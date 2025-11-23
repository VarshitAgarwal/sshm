package sshparser

import (
	"os"
	"path/filepath"

	"github.com/kevinburke/ssh_config"
)

func GetSSHHosts() ([]string, error) {
	home, _ := os.UserHomeDir()
	configPath := filepath.Join(home, ".ssh", "config")

	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg, err := ssh_config.Decode(f)
	if err != nil {
		return nil, err
	}

	var hosts []string

	for _, host := range cfg.Hosts {
		for _, pattern := range host.Patterns {
			hosts = append(hosts, pattern.String())
		}
	}

	return hosts, nil
}
