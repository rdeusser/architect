package git

import (
	"bytes"
	"context"
	"os"
	"os/exec"

	"mvdan.cc/sh/v3/interp"
)

// Funcs provides git builtins to the interpreter.
var Funcs = map[string]interp.ExecHandlerFunc{
	"git::remote": func(ctx context.Context, args []string) error {
		out, err := exec.Command("git", "config", "--local", "remote.origin.url").Output()
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(bytes.TrimSpace(out))
		return err
	},
}
