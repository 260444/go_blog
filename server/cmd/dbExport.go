package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"server/global"
	"time"

	"github.com/spf13/cobra"
)

// DBExportCmd represents the DBExport command
var DBExportCmd = &cobra.Command{
	Use:   "dbExport",
	Short: "导出数据库",
	Long:  `导出数据库`,
	Run: func(cmd *cobra.Command, args []string) {
		err := SQLExport()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(DBExportCmd)
}

func SQLExport() error {
	mysql := global.Config.Mysql
	timer := time.Now().Format("20060102")
	sqlPath := fmt.Sprintf("mysql_%s.sql", timer)
	cmd := exec.Command("docker", "exec", "mysql", "mysqldump", "-u"+mysql.Username, "-p"+mysql.Password, mysql.DBName)

	outFile, err := os.Create(sqlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	return cmd.Run()
}
