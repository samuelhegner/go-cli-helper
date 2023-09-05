/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/samuelhegner/go-cli-helper/commandRunner"
	"github.com/samuelhegner/go-cli-helper/constants"
	"github.com/samuelhegner/go-cli-helper/dirHelper"
	"github.com/samuelhegner/go-cli-helper/gitHelper"
	"github.com/samuelhegner/go-cli-helper/goHelper"
	"github.com/spf13/cobra"
)

var nameString = "name"
var noGitString = "no-git"
var noRemoteString = "no-remote"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go project",
	Long: `This command create a new Go project with the provided name.
	This includes creating a new directory, local and remote git repository, go.main and mod`,
	Run: run,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String(nameString, "", "Name of the new project")
	createCmd.Flags().Bool(noGitString, false, "Do not initiate a git repository")
	createCmd.Flags().Bool(noRemoteString, false, "Do not create a remote repository on GitHub")
}

func run(cmd *cobra.Command, args []string) {

	n, _ := cmd.Flags().GetString(nameString)
	ng, _ := cmd.Flags().GetBool(noGitString)
	nr, _ := cmd.Flags().GetBool(noRemoteString)
	wd, _ := os.Getwd()
	dir := filepath.Join(wd, n)

	earlyStop, err := dirHelper.Exists(dir)

	if earlyStop || err != nil {
		log.Fatal("Directory already exists or error occurred checking")
	}

	if n == "" {
		fmt.Println("Provide a project name using --name flag")
		os.Exit(1)
	}

	if ng {
		nr = true
		fmt.Println(ng, nr)
	}

	fmt.Println("Creating the Go project:", n, "...")

	commandRunner.Run("mkdir", n)

	goHelper.InitGoMod(dir, n)

	if !ng {
		gitHelper.InitLocalRepository(dir)
		gitHelper.CreateInitialCommit(dir)
	}

	if !nr {
		gitHelper.CreateRemoteRepository(n, dir)
		gitHelper.LinkRemoteToLocal(constants.GitHubUrl+n, dir)
		gitHelper.PushLocalFiles(dir)
	}

	fmt.Println("Created project directory")
}
