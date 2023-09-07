package goHelper

import (
	"log"

	"github.com/samuelhegner/go-cli-helper/commandRunner"
	"github.com/samuelhegner/go-cli-helper/constants"
)

func InitGoMod(dir string, name string) {
	log.Println("Initialising Go Mod file...")
	commandRunner.RunInDirectory(dir, "go", "mod", "init", constants.GoPackageRoot+name)
	log.Println("Initialised Go Mod file")
}
