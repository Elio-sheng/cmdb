package cmd

import (
	"fmt"
	"gin-cmdb/cmd/command"
	"gin-cmdb/cmd/cron"
	"gin-cmdb/cmd/server"
	"gin-cmdb/cmd/version"
	"gin-cmdb/config"
	"gin-cmdb/internal/global"
	"gin-cmdb/internal/pkg/logger"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:          "cmdb",
		Short:        "cmdb",
		SilenceUsage: true,
		Long:         `Gin Cmdb`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 1、初始化配置
			config.InitConfig(configPath)
			// 2、初始化zap日志
			logger.InitLogger()
		},
		Run: func(cmd *cobra.Command, args []string) {
			if printVersion {
				fmt.Println(global.Version)
				return
			}

			fmt.Printf("%s\n", "Welcome to cmdb. Use -h to see more commands")
		},
	}
	configPath   string
	printVersion bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "The absolute path of the configuration file")
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "GetUserInfo version info")
	// 查看版本 cmdb version
	rootCmd.AddCommand(version.Cmd)
	// 启动服务 cmdb server
	rootCmd.AddCommand(server.Cmd)
	// 启动单词运行脚本 cmdb command demo
	rootCmd.AddCommand(command.Cmd)
	// 启动计划任务
	rootCmd.AddCommand(cron.Cmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
