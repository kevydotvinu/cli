package main

import (
	"fmt"
	"os"

	"github.com/shipwright-io/cli/pkg/shp/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var dir string
var root = &cobra.Command{
	Use:   "gendoc",
	Short: "Generate shp's help docs",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		genericOpts := &genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
		cmd := cmd.NewCmdSHP(genericOpts)
		cmd.DisableAutoGenTag = true
		return doc.GenMarkdownTree(cmd, dir)
	},
}

func init() {
	os.Setenv("HOME", "~")
	root.Flags().StringVarP(&dir, "dir", "d", ".", "Path to directory in which to generate docs")
}

func main() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
