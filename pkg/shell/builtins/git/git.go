package git

import (
	"bytes"
	"context"
	"os/exec"

	"mvdan.cc/sh/v3/interp"
)

// Funcs provides git builtins to the interpreter.
var Funcs = map[string]interp.ExecHandlerFunc{
	"git::remote": func(ctx context.Context, args []string) error {
		hc := interp.HandlerCtx(ctx)

		out, err := exec.Command("git", "config", "--local", "remote.origin.url").Output()
		if err != nil {
			return err
		}

		_, err = hc.Stdout.Write(bytes.TrimSpace(out))
		return err
	},
}
