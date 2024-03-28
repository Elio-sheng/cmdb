package demo

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "demo",
		Short:   "This is Demo Command",
		Example: "cmdb command demo",
		Run: func(cmd *cobra.Command, args []string) {
			demo()
		},
	}
	test string
)

func init() {
	Cmd.Flags().StringVarP(&test, "test", "t", "test", "用于测试的自定义参数 ")
}

func demo() {
	fmt.Println("hello console!", test)
}
