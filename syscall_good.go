// +build !plan9,!windows

package sftp

import "syscall"

const S_IFMT = syscall.S_IFMT
