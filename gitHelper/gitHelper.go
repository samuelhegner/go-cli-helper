package gitHelper

import (
	"log"
	"os"
	"path/filepath"

	"github.com/samuelhegner/go-cli-helper/commandRunner"
)

func InitLocalRepository(dir string) {
	log.Println("Initialising local repository...")
	commandRunner.RunInDirectory(dir, "git", "init")
	log.Println("Initialised local repository")
}

const (
	ignore = `*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
go.work`
)

func CreateIgnoreFile(dir string) {
	ig := filepath.Join(dir, ".gitignore")

	err := os.WriteFile(ig, []byte(ignore), 0777)

	if err != nil {
		log.Fatal(err)
	}
}

var defaultRemoteCreateFlags = []string{
	"--public",
	"--push",
	"-s=.",
}

func CreateRemoteRepository(name string, dir string) {
	log.Println("Creating remote Repository:", name)
	args := append([]string{"repo", "create", name}, defaultRemoteCreateFlags...)
	commandRunner.RunInDirectory(dir, "gh", args...)
	log.Println("Created remote repository")
}

func LinkRemoteToLocal(remoteUrl string, dir string) {
	log.Println("Linking local and remote repository...")
	commandRunner.RunInDirectory(dir, "git", "remote", "add", "origin", remoteUrl)
	commandRunner.RunInDirectory(dir, "git", "branch", "--set-upstream-to=origin/main", "main")
	commandRunner.RunInDirectory(dir, "git", "pull")
	log.Println("Linked remote repository to local and pulled files")
}

func CreateInitialCommit(dir string) {
	log.Println("Creating initial commit...")
	commandRunner.RunInDirectory(dir, "git", "add", ".")
	commandRunner.RunInDirectory(dir, "git", "commit", "-m", "\"Initial Commit\"")
	log.Println("Created initial commit")
}

func PushLocalFiles(dir string) {
	log.Println("Pushing local changes to remote...")
	commandRunner.RunInDirectory(dir, "git", "push")
	log.Println("Pushed local changes to remote")
}

func RemoteExists(name string) (bool, error) {
	err := commandRunner.RunWithError("gh", "repo", "view", name)
	if err != nil {
		return false, nil
	}

	return true, nil
}
