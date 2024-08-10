package util

import (
	"os"
	"os/exec"
)

func Run(entrypoint string, stdin bool) error {
	cmd := exec.Command("go", "run", entrypoint)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if stdin {
		cmd.Stdin = os.Stdin
	}

	return RunAndClear(cmd)
}

func RunArgv(args []string) error {
	cmd := exec.Command("go", append([]string{"run"}, args...)...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func TestArgv(args []string) error {
	cmd := exec.Command("go", append([]string{"test"}, args...)...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return RunAndClear(cmd)
}

func BuildArgv(args []string) error {
	cmd := exec.Command("go", append([]string{"build"}, args...)...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return RunAndClear(cmd)
}
