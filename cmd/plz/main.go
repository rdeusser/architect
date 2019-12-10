package main

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/rdeusser/please/pkg/shell"
	"github.com/rdeusser/please/pkg/version"
)

var debug bool

func main() {
	cmd := &cobra.Command{
		Use:   "plz",
		Short: "A shell script-based build tool",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logrus.SetOutput(colorable.NewColorableStdout())

			if debug {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Version: fmt.Sprintf("%s (%s)", version.Version, version.GitCommit),
	}

	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug logs")

	commands, err := shell.Commands()
	if err != nil {
		logrus.Fatal(err)
	}

	cmd.AddCommand(commands...)

	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
