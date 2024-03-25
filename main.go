package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Function to execute git add and git commit
func commitChanges() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter commit message: ")
	message, _ := reader.ReadString('\n')

	cmdAdd := exec.Command("git", "add", ".")
	if output, err := cmdAdd.CombinedOutput(); err != nil {

		return fmt.Errorf("error adding files to commit: %s", string(output))
	}

	cmdCommit := exec.Command("git", "commit", "-m", message)
	if output, err := cmdCommit.CombinedOutput(); err != nil {
		return fmt.Errorf("error committing changes: %s", string(output))
	}

	return nil
}

// Function to check if a remote URL is configured and prompt for it if not
func getRemoteURL() error {
	cmdRemote := exec.Command("git", "remote", "get-url", "origin")
	output, err := cmdRemote.CombinedOutput()
	if err != nil {
		// Print the error message
		fmt.Print(string(output))
		// If remote doesn't exist, prompt for the URL and add it
		if strings.Contains(string(output), "No such remote") {
			return addRemoteURL()
		}
		return fmt.Errorf("error getting remote URL: %s", string(output))
	}
	return nil
}

// Function to add remote URL
func addRemoteURL() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter remote URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	cmdRemote := exec.Command("git", "remote", "add", "origin", url)
	output, err := cmdRemote.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error adding remote URL: %v, %s", err, output)
	}
	return nil
}

// Function to execute git push
func pushChanges() error {
	if err := getRemoteURL(); err != nil {
		return err
	}
	cmdPush := exec.Command("git", "push", "-u", "origin", "main")
	output, err := cmdPush.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error pushing changes: %v, %s", err, output)
	}

	fmt.Println("Push successful:", string(output))
	return nil
}

// Main function to handle arguments and execute appropriate commands
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gitswift [push | commit]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "push":
		if err := commitChanges(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := pushChanges(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "commit":
		if err := commitChanges(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: gitswift [push | commit]")
		os.Exit(1)
	}
}
