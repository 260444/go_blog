package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd 是应用程序的根命令，当不带任何子命令调用时会执行此命令
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "A blog server application",
	Long: `A blog server application built with Go.
This application provides various commands for managing the blog server.`,
	// 添加 Run 函数来定义根命令执行时的行为，避免默认显示帮助信息
	Run: func(cmd *cobra.Command, args []string) {
		// 检查是否是 help 命令，如果是则只显示帮助信息
		// 修改: 更准确地判断 help 参数
		if len(os.Args) > 1 && (os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h") {
			// Cobra 会自动处理 help 命令，这里不需要额外操作
			return
		}
		// 如果希望在非 help 情况下执行完逻辑后退出，可以调用 os.Exit(0)
		// os.Exit(0)
	},
}

// Execute 执行根命令，将所有子命令添加到根命令中并适当设置标志
// 这个函数由 main.main() 调用，对于 rootCmd 只需要执行一次
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init 函数在包初始化时自动执行，用于设置根命令的标志
func init() {
	// Flags().BoolP 定义一个布尔类型的标志
	// "toggle" 是标志的名称，"t" 是其短格式，false 是默认值
	// "Help message for toggle" 是帮助信息
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
