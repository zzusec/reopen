//go:build windows

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// execReplace cannot replace the process on Windows (no syscall.Exec), so it
// runs the agent as a child that inherits the console and forwards the exit
// code. The terminal is shared, so the resumed session reads and writes here.
func execReplace(argv0 string, argv []string, env []string) error {
	cmd := exec.Command(argv0, argv...)
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		// Preserve the exit code; a non-zero exit is reported as an error so
		// main() exits the browser the same way.
		if ee, ok := err.(*exec.ExitError); ok {
			os.Exit(ee.ExitCode())
		}
		return fmt.Errorf("%s: %w", argv0, err)
	}
	os.Exit(0)
	return nil
}
