package templates

import "html/template"

// 全局变量，保存模板实例
var Template *template.Template

// 加载模板文件的函数
func LoadTemplate(templateFile string) {
	// 使用 Must 函数解析所有匹配的模板文件
	Template = template.Must(template.ParseGlob(templateFile))
}
