package util

import (
	"bytes"
	"fmt"
	"github.com/ahmetalpbalkan/go-cursor"
	"io"
	"os"
	"os/exec"
)

func RunAndClear(cmd *exec.Cmd) error {
	buffer := new(bytes.Buffer)

	// pipe stdout and stderr to buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, buffer)
	cmd.Stderr = io.MultiWriter(os.Stderr, buffer)

	// run the command
	if err := cmd.Run(); err != nil {
		return err
	}

	// count lines in buffer
	lines := bytes.Count(buffer.Bytes(), []byte("\n"))

	if lines == 0 {
		return nil
	}

	// move cursor up by the number of lines
	fmt.Print(cursor.MoveUp(lines))
	fmt.Print(cursor.ClearEntireLine())
	fmt.Print(cursor.ClearScreenDown())

	return nil
}
