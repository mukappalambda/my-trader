package common

import (
	"fmt"
	"os"
)

func PrintToStderrThenExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
