package git

import (
	"bytes"
	"os/exec"

	"mvdan.cc/sh/v3/interp"
)

// ExecModule provides git information to the interpreter.
func ExecModule() []func(interp.ExecModule) interp.ExecModule {
	return []func(interp.ExecModule) interp.ExecModule{
		interp.ExecBuiltin("git::remote", func(ctx interp.ModuleCtx, args []string) error {
			out, err := exec.Command("git", "config", "--local", "remote.origin.url").Output()
			if err != nil {
				return err
			}
			_, err = ctx.Stdout.Write(bytes.TrimSpace(out))
			return err
		}),
	}
}
