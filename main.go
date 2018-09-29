package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// Setup the command tobe run
	cmd := exec.Command("wc", "-c")

	// Get stdin and stdout
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// Start the process
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Write the data
	if _, err := stdin.Write([]byte("abcracadabra")); err != nil {
		panic(err)
	}
	stdin.Close()

	// Get the result
	result, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}

	// Wait for the command to exit
	// if err = cmd.Wait(); err != nil {
	// 	panic(err)
	// }

	// Print the result
	fmt.Println(string(result))
	return
}
