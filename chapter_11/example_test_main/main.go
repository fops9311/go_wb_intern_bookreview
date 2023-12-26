package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

// функция - имитатор мэйна, которую уже можно протестировать
func run(args []string, stdout io.Writer) error {
	if len(args) < 2 {
		return errors.New("no input")
	}
	// продолжаем реализацию run(),
	// как будто это main()
	return nil
}
