package ssh

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"ioutil"
	"os"
	"os/user"
	"strings"
)

func startRemoteShell(alias string) error {
	host, userStr, keyPath, err := LoadSSHConfig(alias)
	if err != nil {
		return err
	}

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("cannot read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return fmt.Errorf("cannot parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: userStr,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}

	fmt.Printf("\n🔗 Connected to: %s\n", alias)
	fmt.Println("Type 'exit' to quit.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s ➜ ", alias)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "exit" {
			fmt.Println("Closing session...")
			return nil
		}

		sess, err := conn.NewSession()
		if err != nil {
			return fmt.Errorf("cannot create session: %v", err)
		}

		output, err := sess.CombinedOutput(command)
		sess.Close()

		fmt.Println(string(output))

		if err != nil {
			fmt.Printf("⚠️ Error: %v\n", err)
		}
	}
}
