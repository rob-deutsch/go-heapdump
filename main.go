package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime/debug"
)

func main() {
	// Setup the command tobe run
	cmd := exec.Command("wc", "-c")

	// Get stdin reader
	stdinReader, stdinWriter, err := stdinPipe(cmd)
	if err != nil {
		panic(err)
	}

	// Get stdout reader
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// Start the process
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	// We no longer need this!
	stdinReader.Close()

	// Write the data
	debug.WriteHeapDump(stdinWriter.Fd())
	stdinWriter.Close()

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

func stdinPipe(cmd *exec.Cmd) (*os.File, *os.File, error) {
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	cmd.Stdin = stdinReader

	return stdinReader, stdinWriter, nil
}
