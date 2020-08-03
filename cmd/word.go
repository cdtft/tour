package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"tour/internal/word"
)

var str string
var mode int

const (
	MODE_UPPER = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var desc = strings.Join([]string{
	"改子命令支持各种单词格式转换，模式如下: ",
	"1: 转换为大写",
	"2: 转换为小写",
	"3: 下划线转换为大写驼峰",
	"4: 下划线转换为小写驼峰",
	"5: 驼峰转换为下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
			break
		case MODE_LOWER:
			content = word.ToLower(str)
			break
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
			break
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
			break
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
			break
		default:
			log.Fatalf("暂不支持该转换模式")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "输出入单词内容")
	wordCmd.Flags().IntVarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}