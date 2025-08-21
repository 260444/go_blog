package cmd

import (
	"server/global"
	"server/model/database"

	"github.com/spf13/cobra"
)

// DBCmd represents the DB command
var DBCmd = &cobra.Command{
	Use:   "db",
	Short: "生成数据库表结构",
	Long:  `生成数据库表结构`,
	Run: func(cmd *cobra.Command, args []string) {
		SQL()
	},
}

func init() {
	rootCmd.AddCommand(DBCmd)
}

func SQL() error {
	return global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&database.Advertisement{},
		&database.ArticleCategory{},
		&database.ArticleLike{},
		&database.ArticleTag{},
		&database.Comment{},
		&database.Feedback{},
		&database.FooterLink{},
		&database.FriendLink{},
		&database.Image{},
		&database.JwtBlacklist{},
		&database.Login{},
		&database.User{},
	)
}
