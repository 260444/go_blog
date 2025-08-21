package main

import (
	"os"
	"server/cmd"
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectEs()

	defer global.Redis.Close()

	// 检查是否有命令行参数
	// 修改: os.Args[0] 是程序名，参数从 os.Args[1] 开始
	if len(os.Args) > 1 {
		// 如果有参数，则执行命令逻辑
		cmd.Execute()
	} else {
		// 如果没有参数，则直接执行主逻辑
		initialize.InitCron()
		core.RunServer()
	}
}
