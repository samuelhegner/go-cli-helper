package gitHelper

import (
	"fmt"

	"github.com/samuelhegner/go-cli-helper/commandRunner"
)

func InitLocalRepository(dir string) {
	fmt.Println("Initialising local repository...")
	commandRunner.RunInDirectory(dir, "git", "init")
	fmt.Println("Initialised local repository")
}

var defaultRemoteCreateFlags = []string{
	"--add-readme",
	"--public",
	"-g=Go",
}

func CreateRemoteRepository(name string, dir string) {
	fmt.Println("Creating remote Repository:", name)
	args := append([]string{"repo", "create", name}, defaultRemoteCreateFlags...)
	commandRunner.RunInDirectory(dir, "gh", args...)
	fmt.Println("Created remote repository")
}

func LinkRemoteToLocal(remoteUrl string, dir string) {
	fmt.Println("Linking local and remote repository...")
	commandRunner.RunInDirectory(dir, "git", "remote", "add", "origin", remoteUrl)
	commandRunner.RunInDirectory(dir, "git", "pull")
	fmt.Println("Linked remote repository to local and pulled files")
}

func CreateInitialCommit(dir string) {
	fmt.Println("Creating initial commit...")
	commandRunner.RunInDirectory(dir, "git", "add", ".")
	commandRunner.RunInDirectory(dir, "git", "commit", "-m", "\"Initial Commit\"")
	fmt.Println("Created initial commit")
}

func PushLocalFiles(dir string) {
	fmt.Println("Pushing local changes to remote...")
	commandRunner.RunInDirectory(dir, "git", "push")
	fmt.Println("Pushed local changes to remote")
}
