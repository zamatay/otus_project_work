package sysInfo

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func GetByPath[T any](path string, f func(file io.Reader) (*T, error)) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return f(file)
}

func ExecuteCommand(command string, args ...string) []byte {
	cmd := exec.Command(command, args...)
	stdOut, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	return stdOut
}
