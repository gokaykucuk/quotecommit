package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "commit", "-m", "quote_placeholder")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Run()
}
