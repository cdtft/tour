package cmd

import "github.com/spf13/cobra"

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "SQL to struct",
	Short: "sql转换处理",
	Long:  "sql转换处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
