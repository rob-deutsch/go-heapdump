package main

import (
	"os"
	"os/exec"
)

func main() {
	//cmd := exec.Command("wc", "-b")
	cmd := exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return
}
