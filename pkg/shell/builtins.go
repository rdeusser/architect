package shell

import (
	"mvdan.cc/sh/v3/interp"

	"github.com/rdeusser/please/pkg/shell/builtins/git"
	"github.com/rdeusser/please/pkg/shell/builtins/log"
	"github.com/rdeusser/please/pkg/shell/builtins/os"
	"github.com/rdeusser/please/pkg/shell/builtins/project"
)

var builtins = flattenModules(
	git.ExecModule(),
	log.ExecModule(),
	os.ExecModule(),
	project.ExecModule(),
)

func flattenModules(modules ...[]func(module interp.ExecModule) interp.ExecModule) (mods []func(interp.ExecModule) interp.ExecModule) {
	for _, m := range modules {
		mods = append(mods, m...)
	}
	return mods
}
