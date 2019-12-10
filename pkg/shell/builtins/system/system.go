package system

import (
	"context"
	"os"

	"mvdan.cc/sh/v3/interp"
)

// Funcs provides system-like functionality to the interpreter.
var Funcs = map[string]interp.ExecHandlerFunc{
	"system::pwd": func(ctx context.Context, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write([]byte(dir))
		return err
	},
}
