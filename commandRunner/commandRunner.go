package commandRunner

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func Run(cmdStr string, arg ...string) {

	log.Println("Running command:", cmdStr, arg)
	cmd := exec.Command(cmdStr, arg...)
	execCmd(cmd)
}

func RunWithError(cmdStr string, arg ...string) error {
	cmd := exec.Command(cmdStr, arg...)
	return execCmdWithErrorReturn(cmd)
}

func RunInDirectory(dir string, cmdStr string, arg ...string) {
	log.Println("Running command:", cmdStr, arg, "In directory:", dir)

	cmd := exec.Command(cmdStr, arg...)

	cmd.Dir = dir
	execCmd(cmd)
}

func execCmd(cmd *exec.Cmd) {
	stderr, err := cmd.StderrPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func execCmdWithErrorReturn(cmd *exec.Cmd) error {
	_, err := cmd.StderrPipe()

	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
