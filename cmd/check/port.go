package check

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port [port_number]",
	Short: "检查指定端口的占用情况",
	Long:  `检查系统中指定端口被哪些进程占用`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		port := args[0]
		
		// 使用 lsof 命令查看端口占用
		output, err := exec.Command("lsof", "-i", ":"+port).Output()
		if err != nil {
			if strings.Contains(err.Error(), "exit status 1") {
				fmt.Printf("端口 %s 当前没有被占用\n", port)
				return nil
			}
			return fmt.Errorf("检查端口时发生错误: %v", err)
		}

		fmt.Printf("端口 %s 的占用情况：\n", port)
		fmt.Println(string(output))
		return nil
	},
} 