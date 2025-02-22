package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 扫描常用端口
	portsToCheck := []int{80, 443, 3306, 6379, 8080, 8443, 9000, 9090}

	fmt.Println("正在检查端口占用情况...")

	for _, port := range portsToCheck {
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)

		if err != nil {
			if strings.Contains(err.Error(), "address already in use") {
				fmt.Printf("警告: 端口 %d 已被占用\n", port)
			} else {
				fmt.Printf("检查端口 %d 时发生错误: %v\n", port, err)
			}
		} else {
			listener.Close()
			fmt.Printf("端口 %d 可用\n", port)
		}
	}
}
