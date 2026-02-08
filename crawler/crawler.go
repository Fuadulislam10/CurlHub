package crawler

import (
	"os/exec"
)

func Fetch(url string) string {
	cmd := exec.Command("curl", "-Ls", url)
	out, _ := cmd.Output()
	return string(out)
}
