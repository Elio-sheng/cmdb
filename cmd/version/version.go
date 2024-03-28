package version

import (
	"fmt"
	"gin-cmdb/internal/global"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "version",
		Short:   "GetUserInfo version info",
		Example: "cmdb version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(global.Version)
		},
	}
)
