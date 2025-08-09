package cmd

import "os"

func Exit() error {
	os.Exit(0)
	return nil
}
