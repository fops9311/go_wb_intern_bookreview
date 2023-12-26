package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	main()
}
func TestRun(t *testing.T) {
	{
		err := run([]string{}, os.Stdout)
		if err != nil {
			t.Error(err)
		}
	}
	{
		err := run([]string{"name"}, os.Stdout)
		if err != nil {
			t.Error(err)
		}
	}
	{
		err := run([]string{"name", "arg1"}, os.Stdout)
		if err != nil {
			t.Error(err)
		}
	}
}
