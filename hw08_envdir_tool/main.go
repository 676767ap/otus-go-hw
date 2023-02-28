package main

import (
	"log"
	"os"
)

func main() {
	pathEnv := os.Args[1]
	cmd := os.Args[2:]

	env, err := ReadDir(pathEnv)
	if err != nil {
		log.Fatalf("Read dir error: %v", err)
		os.Exit(1)
	} else {
		os.Exit(RunCmd(cmd, env))
	}
}
