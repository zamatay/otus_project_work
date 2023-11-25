package sysInfo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
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

func AsyncExecuteCommand(command string, args ...string) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	cmd := exec.Command(command, args...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		wg.Done()
	}()
	cmd.Output()
	wg.Wait()
}
