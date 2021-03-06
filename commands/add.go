package commands

import (
	"os"
	"os/exec"

	"code.google.com/archive/p/gopass"
)

func AddKeychainItem(name string) {
	keychainName := "mmfa_" + name
	output, err := addItem(keychainName, secretFromUser())
	if err != nil {
		os.Stderr.Write([]byte("Error adding item: " + err.Error() + "\n" + output))
		os.Exit(1)
	}
	println("Added", name)
}

func addItem(name string, secret string) (string, error) {
	command, args := addItemCommand(name, secret)
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func addItemCommand(name string, secret string) (string, []string) {
	args := []string{
		"add-generic-password",
		"-a",
		name,
		"-s",
		name,
		"-w",
		secret,
	}
	return "security", args
}

func secretFromUser() string {
	secret, err := gopass.GetPass("Enter secret:\n")
	if err != nil {
		os.Stderr.Write([]byte("Error: " + err.Error()))
		os.Exit(1)
	}

	if len(secret) == 0 {
		os.Stderr.Write([]byte("Secret can't be blank"))
		os.Exit(1)
	}

	return secret
}
