package os

import (
	"os"

	"mvdan.cc/sh/v3/interp"
)

// ExecModule provides os-specific functionality to the interpreter.
func ExecModule() []func(interp.ExecModule) interp.ExecModule {
	return []func(interp.ExecModule) interp.ExecModule{
		interp.ExecBuiltin("os::pwd", func(ctx interp.ModuleCtx, args []string) error {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			_, err = ctx.Stdout.Write([]byte(dir))
			return err
		}),
	}
}
