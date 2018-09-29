package heapsize

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime/debug"
	"strconv"
	"strings"
)

// Size returns the size of the heap dump
func Size() (uint64, error) {
	// Setup the command tobe run
	cmd := exec.Command("wc", "-c")

	// Get stdin reader
	stdinReader, stdinWriter, err := stdinPipe(cmd)
	if err != nil {
		return 0, err
	}

	// Get stdout reader
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}

	// Start the process
	if err := cmd.Start(); err != nil {
		return 0, err
	}
	// We no longer need this!
	stdinReader.Close()

	// Write the data
	debug.WriteHeapDump(stdinWriter.Fd())
	stdinWriter.Close()

	// Get the result
	result, err := ioutil.ReadAll(stdout)
	if err != nil {
		return 0, err
	}

	// Print the result
	strSize := strings.Trim(string(result), " \r\n")
	size, err := strconv.ParseUint(strSize, 10, 64)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func stdinPipe(cmd *exec.Cmd) (*os.File, *os.File, error) {
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	cmd.Stdin = stdinReader

	return stdinReader, stdinWriter, nil
}
