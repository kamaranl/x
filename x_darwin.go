//go:build darwin

package x

import (
	"fmt"
	"os/exec"
)

func (a *Alert) alert() {
	var level string
	switch a.Level {
	case Informational:
		level = "note"
	case Warning:
		level = "caution"
	case Critical:
		level = "stop"
	}

	format := "display dialog %q with title %q with icon %s"
	script := fmt.Sprintf(format, a.Message, a.Title, level)
	cmd := exec.Command("osascript", "-e", script)
	_ = cmd.Start()
	a.OK = cmd.ProcessState.Success()
}

func OpenUrl(url string) error { return exec.Command("open", url).Start() }
