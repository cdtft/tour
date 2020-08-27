package cmd

import (
	"github.com/cdtft/tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "please input database username")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "please input database password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "", "please input database host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "", "please input database charset")
	sql2structCmd.Flags().StringVarP(&dbType, "dbType", "", "mysql", "please input database dbType")
	sql2structCmd.Flags().StringVarP(&dbName, "dbName", "", "", "please input database dbName")
	sql2structCmd.Flags().StringVarP(&tableName, "tableName", "", "", "please input database tableName")
}

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换处理",
	Long:  "sql转换处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}
		dbMode := sql2struct.NewDBModel(dbInfo)
		err := dbMode.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err : %v", err)
		}
		columns, err := dbMode.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err : %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err : %v", err)
		}
	},
}
