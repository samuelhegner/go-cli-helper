package goHelper

import (
	"log"
	"os"
	"path/filepath"

	"github.com/samuelhegner/go-cli-helper/commandRunner"
	"github.com/samuelhegner/go-cli-helper/constants"
)

const main string = `package main

func main() {

}`

func InitGoMod(dir string, name string) {
	log.Println("Initialising Go Mod file...")
	commandRunner.RunInDirectory(dir, "go", "mod", "init", constants.GoPackageRoot+name)

	mainDir := filepath.Join(dir, "main.go")

	os.WriteFile(mainDir, []byte(main), 0777)

	log.Println("Initialised Go Mod file")
}
