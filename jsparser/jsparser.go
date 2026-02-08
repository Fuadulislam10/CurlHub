package jsparser

import (
	"os"
	"os/exec"
	"regexp"
)

func Parse(jsURL string) {
	cmd := exec.Command("curl", "-s", jsURL)
	out, err := cmd.Output()
	if err != nil {
		return
	}

	re := regexp.MustCompile(`["'](/api/[^"']+)["']`)
	matches := re.FindAllStringSubmatch(string(out), -1)

	f, _ := os.OpenFile("output/js_endpoints.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	for _, m := range matches {
		f.WriteString(m[1] + "\n")
	}
}
