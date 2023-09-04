/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/samuelhegner/go-cli-helper/exec"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().String(nameString, "", "Name of the new project")
	createCmd.Flags().Bool(noGitString, false, "Do not initiate a git repository")
	createCmd.Flags().Bool(noRemoteString, false, "Do not create a remote repository on GitHub")
}

func run(cmd *cobra.Command, args []string) {

	n, _ := cmd.Flags().GetString(nameString)
	ng, _ := cmd.Flags().GetBool(noGitString)
	nr, _ := cmd.Flags().GetBool(noRemoteString)

	if n == "" {
		fmt.Println("Provide a project name using --name flag")
		os.Exit(1)
	}

	if ng {
		nr = true
		fmt.Println(ng, nr)
	}

	fmt.Println("Creating the Go project:", n)

	command := exec.Command("mkdir", n)

	stderr, _ := command.StderrPipe()
	if err := command.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("Created project directory...")
}
