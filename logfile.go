//go:build !windows
// +build !windows

package well

import (
	"io"
	"syscall"

	"github.com/gswan-io/log"
)

func openLogFile(filename string) (io.Writer, error) {
	return log.NewFileReopener(filename, syscall.SIGUSR1)
}
