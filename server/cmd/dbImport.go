package cmd

import (
	"os"
	"server/global"
	"strings"

	"github.com/spf13/cobra"
)

// DBImportCmd represents the DBImport command
var DBImportCmd = &cobra.Command{
	Use:   "dbImport",
	Short: "导入 MySQL 数据",
	Long:  `导入 MySQL 数据`,
	Run: func(cmd *cobra.Command, args []string) {
		sqlfile, _ := cmd.Flags().GetString("sqlfile")
		err := SQLImport(sqlfile)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	DBImportCmd.Flags().StringP("sqlfile", "f", "", "导入的数据文件路径")
	rootCmd.AddCommand(DBImportCmd)
}

func SQLImport(sqlPath string) (errs []error) {
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		return append(errs, err)
	}
	// 分割数据
	sqlList := strings.Split(string(byteData), ";")
	for _, sql := range sqlList {
		// 去除字符串开头和结尾的空白符
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		// 执行sql语句
		err = global.DB.Exec(sql).Error
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	return nil
}
