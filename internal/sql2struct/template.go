package sql2struct

import (
	"fmt"
	"github.com/cdtft/tour/internal/word"
	"html/template"
	"os"
)

const structTemplate = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}  {{ $length := len .Comment}} {{ if gt $length 0 }}//
{{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typelen := len.Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}
	{{.Type}}    {{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTemplate}
}

// 数据库字段转换为模版渲染前的字段信息
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:"+"%s"+"`", column.ColumnName),
			Comment: column.ColumnName,
		})
	}
	return tplColumns
}

//渲染模版数据
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToLowerCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}

