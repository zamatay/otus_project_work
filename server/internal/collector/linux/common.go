package sysInfo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"project_work/internal/log"
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

func ExecuteCommand(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	stdOut, err := cmd.Output()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			log.Logger.Log.Error(fmt.Sprintf("Ошибка при получении статистики. Необходимо установить утилиту %s", command))
		} else {
			log.Logger.Log.Error(fmt.Sprintf("Ошибка при выполнении %s ", command), err)
		}
		return nil, err
	}
	return stdOut, nil
}

func AsyncExecuteCommand(command string, args ...string) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	cmd := exec.Command(command, args...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Logger.Fatal(err)
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
