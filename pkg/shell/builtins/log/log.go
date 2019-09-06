package log

import (
	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/interp"
)

// ExecModule provides log functionality to the interpreter.
func ExecModule() []func(interp.ExecModule) interp.ExecModule {
	return []func(interp.ExecModule) interp.ExecModule{
		interp.ExecBuiltin("log::panic", func(ctx interp.ModuleCtx, args []string) error {
			for _, arg := range args {
				logrus.Panicln(arg)
			}
			return nil
		}),
		interp.ExecBuiltin("log::fatal", func(ctx interp.ModuleCtx, args []string) error {
			for _, arg := range args {
				logrus.Fatalln(arg)
			}
			return nil
		}),
		interp.ExecBuiltin("log::error", func(ctx interp.ModuleCtx, args []string) error {
			for _, arg := range args {
				logrus.Errorln(arg)
			}
			return nil
		}),
		interp.ExecBuiltin("log::warn", func(ctx interp.ModuleCtx, args []string) error {
			for _, arg := range args {
				logrus.Warnln(arg)
			}
			return nil
		}),
		interp.ExecBuiltin("log::info", func(ctx interp.ModuleCtx, args []string) error {
			for _, arg := range args {
				logrus.Infoln(arg)
			}
			return nil
		}),
		interp.ExecBuiltin("log::debug", func(ctx interp.ModuleCtx, args []string) error {
			for _, arg := range args {
				logrus.Debugln(arg)
			}
			return nil
		}),
	}
}
