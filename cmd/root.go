package cmd

import (
	"fmt"
	"os"
	"health_check/cmd/check"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "health_check",
	Short: "健康检查工具",
	Long:  `一个用于系统健康检查的命令行工具，可以检查端口占用等信息`,
}

func init() {
	rootCmd.AddCommand(check.CheckCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
} 