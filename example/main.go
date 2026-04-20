package main

import (
	"os"

	"github.com/asgardehs/ratatoskr/python"
)

func main() {
	ep, err := python.NewEmbeddedPython("example")
	if err != nil {
		panic(err)
	}

	cmd, err := ep.PythonCmd("-c", "print('hello')")
	if err != nil {
		panic(err)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
