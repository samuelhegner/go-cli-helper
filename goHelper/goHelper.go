package goHelper

import (
	"fmt"

	"github.com/samuelhegner/go-cli-helper/commandRunner"
	"github.com/samuelhegner/go-cli-helper/constants"
)

func InitGoMod(dir string, name string) {
	fmt.Println("Initialising Go Mod file...")
	commandRunner.RunInDirectory(dir, "go", "mod", "init", constants.GoPackageRoot+name)
	fmt.Println("Initialised Go Mod file")
}
