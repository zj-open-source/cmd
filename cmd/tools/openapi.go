package main

import (
	"github.com/go-courier/httptransport/generators/openapi"
	_ "github.com/go-courier/httptransport/validator/strfmt"
	"github.com/go-courier/packagesx"
	"github.com/spf13/cobra"
	"github.com/zj-open-source/cmd/internal/generate"
)

var cmdSwagger = &cobra.Command{
	Use:     "openapi",
	Aliases: []string{"swagger"},
	Short:   "scan current project and generate openapi.json",
	Run: func(cmd *cobra.Command, args []string) {
		generate.RunGenerator(func(pkg *packagesx.Package) generate.Generator {
			g := openapi.NewOpenAPIGenerator(pkg)
			g.Scan(cmd.Context())
			return g
		})
	},
}

func init() {
	cmdRoot.AddCommand(cmdSwagger)
}
