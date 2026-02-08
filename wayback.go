package main

import (
	"os/exec"
)

func RunWayback(domain string) {
	exec.Command("gau", domain).Run()
}
