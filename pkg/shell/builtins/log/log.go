package log

import (
	"context"

	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/interp"
)

// Funcs provides log builtins to the interpreter.
var Funcs = map[string]interp.ExecHandlerFunc{
	"log::panic": func(ctx context.Context, args []string) error {
		for _, arg := range args {
			logrus.Panicln(arg)
		}
		return nil
	},
	"log::fatal": func(ctx context.Context, args []string) error {
		for _, arg := range args {
			logrus.Fatalln(arg)
		}
		return nil
	},
	"log::error": func(ctx context.Context, args []string) error {
		for _, arg := range args {
			logrus.Errorln(arg)
		}
		return nil
	},
	"log::warn": func(ctx context.Context, args []string) error {
		for _, arg := range args {
			logrus.Warnln(arg)
		}
		return nil
	},
	"log::info": func(ctx context.Context, args []string) error {
		for _, arg := range args {
			logrus.Infoln(arg)
		}
		return nil
	},
	"log::debug": func(ctx context.Context, args []string) error {
		for _, arg := range args {
			logrus.Debugln(arg)
		}
		return nil
	},
}
