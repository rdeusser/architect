package shell

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"

	"github.com/rdeusser/please/pkg/shell/builtins/git"
	"github.com/rdeusser/please/pkg/shell/builtins/log"
	"github.com/rdeusser/please/pkg/shell/builtins/project"
	"github.com/rdeusser/please/pkg/shell/builtins/system"
)

var allBuiltins = []map[string]interp.ExecHandlerFunc{
	git.Funcs,
	log.Funcs,
	project.Funcs,
	system.Funcs,
}

var builtins = make(map[string]interp.ExecHandlerFunc)

func init() {
	// Loop over all builtins and flatten them to a single map.
	for _, m := range allBuiltins {
		for k, v := range m {
			builtins[k] = v
		}
	}
}

// Commands builds a command list from shell scripts and turns them into cobra commands.
func Commands() ([]*cobra.Command, error) {
	var commands []*cobra.Command

	matches, err := filepath.Glob("scripts/*.sh")
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		in, err := ioutil.ReadFile(match)
		if err != nil {
			return nil, err
		}

		reader := strings.NewReader(string(in))

		f, err := syntax.NewParser().Parse(reader, commandName(match))
		if err != nil {
			return nil, err
		}

		usage, err := printNode(f, "usage")
		if err != nil {
			return nil, err
		}

		commands = append(commands, &cobra.Command{
			Use:   f.Name,
			Short: usage,
			RunE: func(cmd *cobra.Command, args []string) error {
				out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
				if err != nil {
					return err
				}
				repoRoot := string(bytes.TrimSpace(out))

				logrus.Debugf("Using repo root: %s", repoRoot)

				r, err := interp.New(
					interp.StdIO(nil, os.Stdout, os.Stdout),
					interp.Dir(repoRoot),
					interp.Env(expand.ListEnviron(os.Environ()...)),
					interp.ExecHandler(execHandler),
				)
				if err != nil {
					return err
				}

				err = r.Run(context.TODO(), f)
				if err != nil {
					return fmt.Errorf("could not run: %v", err)
				}

				// Delete the internal shell vars that nobody cares about.
				delete(r.Vars, "IFS")
				delete(r.Vars, "OPTIND")

				return nil
			},
		})
	}

	return commands, nil
}

func execHandler(ctx context.Context, args []string) error {
	if fn := builtins[args[0]]; fn != nil {
		return fn(ctx, args[1:])
	}
	return interp.DefaultExecHandler(2*time.Second)(ctx, args)
}

// commandName takes a path to a script in the scripts directory and returns what the command name should be without
// the file extension.
func commandName(filename string) string {
	return filepath.Base(strings.ReplaceAll(filename, filepath.Ext(filename), ""))
}

// printNode takes a variable name and returns the pretty-printed value of it.
//
// Currently, the only supported node syntax is Assign.
func printNode(f *syntax.File, name string) (string, error) {
	printer := syntax.NewPrinter()
	b := strings.Builder{}

	syntax.Walk(f, func(node syntax.Node) bool {
		switch x := node.(type) {
		case *syntax.Assign:
			if x.Name.Value == name {
				err := printer.Print(&b, x.Value)
				if err != nil {
					logrus.Error(err)
					return false
				}
			}
		}
		return true
	})

	if len(b.String()) == 0 {
		return "usage not provided", nil
	}

	v, err := strconv.Unquote(b.String())
	if err != nil {
		return "", err
	}

	return v, nil
}
