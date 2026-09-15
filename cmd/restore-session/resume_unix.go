//go:build !windows

package main

import "syscall"

// execReplace swaps this process for the one named by argv[0]. On Unix the
// agent's interactive session takes over the controlling terminal directly,
// which is what makes "resume into the same shell" work.
func execReplace(argv0 string, argv []string, env []string) error {
	return syscall.Exec(argv0, argv, env)
}
