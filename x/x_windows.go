//go:build windows

package x

import "os/exec"

func OpenUrl(url string) error {
	return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
}
