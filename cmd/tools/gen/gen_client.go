package gen

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"

	client "github.com/go-courier/httptransport/generators/client"
	"github.com/go-courier/packagesx"
	"github.com/spf13/cobra"
	"github.com/zj-open-source/cmd/internal/generate"
)

var (
	cmdGenClientFlagFile    string
	cmdGenClientFlagSpecURL string
)

func init() {
	CmdGen.AddCommand(cmdGenClient)

	cmdGenClient.Flags().
		StringVarP(&cmdGenClientFlagSpecURL, "spec-url", "", "", "client spec url")
	cmdGenClient.Flags().
		StringVarP(&cmdGenClientFlagFile, "file", "", "", "client spec file")

}

var cmdGenClient = &cobra.Command{
	Use:     "client",
	Example: "client demo",
	Short:   "generate client by open api",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			panic(fmt.Errorf("need service name"))
		}
		u := &url.URL{}

		if cmdGenClientFlagFile != "" {
			u.Scheme = "file"
			u.Path = cmdGenClientFlagFile

			if !filepath.IsAbs(u.Path) {
				cwd, _ := os.Getwd()
				u.Path = path.Join(cwd, u.Path)
			}
		}

		if cmdGenClientFlagSpecURL != "" {
			uri, err := url.Parse(cmdGenClientFlagSpecURL)
			if err != nil {
				panic(err)
			}
			u = uri
		}

		generate.RunGenerator(func(pkg *packagesx.Package) generate.Generator {
			g := client.NewClientGenerator(args[0], u, client.OptionVendorImportByGoMod())
			g.Load()
			return g
		})
	},
}
