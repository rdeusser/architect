package project

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"

	"mvdan.cc/sh/v3/interp"
)

// ExecModule provides project information to the interpreter.
func ExecModule() []func(interp.ExecModule) interp.ExecModule {
	return []func(interp.ExecModule) interp.ExecModule{
		interp.ExecBuiltin("project::root", func(ctx interp.ModuleCtx, args []string) error {
			out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
			if err != nil {
				return err
			}
			_, err = ctx.Stdout.Write(bytes.TrimSpace(out))
			return err
		}),
		interp.ExecBuiltin("project::name", func(ctx interp.ModuleCtx, args []string) error {
			repo, err := gitRepo()
			if err != nil {
				return err
			}
			_, err = ctx.Stdout.Write([]byte(filepath.Base(repo)))
			return err
		}),
		interp.ExecBuiltin("project::repo", func(ctx interp.ModuleCtx, args []string) error {
			out, err := exec.Command("go", "list", "-m").Output()
			if err != nil {
				return err
			}
			// Are we in a Go project with modules setup?
			// The output is strange, but the command exits successfully
			// outputting this message so we need to check for it.
			if bytes.Contains(out, []byte("command-line-arguments")) {
				repo, err := gitRepo()
				if err != nil {
					return err
				}
				_, err = ctx.Stdout.Write([]byte(repo))
				if err != nil {
					return err
				}
				return nil
			}
			_, err = ctx.Stdout.Write(bytes.TrimSpace(out))
			return err
		}),
	}
}

func gitRepo() (string, error) {
	out, err := exec.Command("git", "config", "--local", "remote.origin.url").Output()
	if err != nil {
		return "", err
	}
	// We don't need or want the .git extension.
	out = bytes.ReplaceAll(out, []byte(".git"), []byte(""))
	// Return the basename of the git remote as the project name.
	return strings.TrimSpace(string(out)), nil
}
