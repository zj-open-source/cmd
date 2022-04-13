package hook

import (
	"github.com/spf13/cobra"
	"github.com/zj-open-source/cmd/internal/githooks"
)

func init() {
	CmdHook.AddCommand(cmdHookInit)
}

var cmdHookInit = &cobra.Command{
	Use:   "init",
	Short: "git hook init",
	Run: func(cmd *cobra.Command, args []string) {
		githooks.Init()
	},
}
