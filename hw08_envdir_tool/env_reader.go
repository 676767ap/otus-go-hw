package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		envLine, err := prepareLine(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		if len(envLine) == 0 {
			envs[file.Name()] = EnvValue{"", true}
		} else {
			envs[file.Name()] = EnvValue{envLine, false}
		}
	}
	return envs, nil
}

func prepareLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	line, _, err := scanner.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return "", nil
		}
		return "", err
	}

	line = bytes.ReplaceAll(line, []byte("\x00"), []byte("\n"))
	trimmedLine := strings.TrimRight(string(line), " ")
	return trimmedLine, nil
}
