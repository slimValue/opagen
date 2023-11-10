package opagen

import (
	"bytes"
	_ "embed"
	"os"
	"text/template"
)

const (
	LOGIC_TYPE_AND = "and" // 逻辑与
	LOGIC_TYPE_OR  = "or"  // 逻辑或

	EFFECT_TYPE_ALLOW = "allow" // 允许
	EFFECT_TYPE_DENY  = "deny"  // 拒绝
)

type Rego struct {
	PkgName   string     // 包名
	Imports   []string   // 导入包
	LogicType string     // 逻辑类型 and or
	RuleItems []RuleItem // 复合规则
}

// RuleItem 复合 授权规则(多个函数组合)
type RuleItem struct {
	FnName string
	Effect string // 授权作用 允许或拒绝
	Subs   []Rule // 主体条件
	Rules  []Rule // 限制规则
}

type Rule struct {
	FnName string // 函数名
	FnBody string // 函数方法体
	Desc   string // 描述
}

//go:embed rule.tpl
var tpl string

func Gen(data Rego) string {
	t := template.New("rule")
	t = template.Must(t.Parse(tpl))

	_ = os.Stdout
	var buf bytes.Buffer
	t.Execute(&buf, data)
	return buf.String()
}
