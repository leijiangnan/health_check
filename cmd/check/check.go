package check

import (
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "执行各种检查命令",
	Long:  `执行系统各种健康状态检查，如端口检查等`,
}

func init() {
	CheckCmd.AddCommand(portCmd)
} 