//go:build darwin

package x

import "os/exec"

func OpenUrl(url string) error { return exec.Command("open", url).Start() }
