package version

import (
	"os/exec"
	"strings"
)

var (
	Version   = "v0.1.0"
	GitCommit = gitCommit()
)

func gitCommit() string {
	out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		return "HEAD"
	}
	return strings.TrimSpace(string(out))
}
