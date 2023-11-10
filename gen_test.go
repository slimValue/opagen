package opagen

import "testing"

var (
	test_imports = []string{"future.keywords.in", "future.keywords.if"}
)

func Test_Gen(t *testing.T) {

	r := Rego{
		PkgName:   "test",
		LogicType: LOGIC_TYPE_AND,
		Imports:   test_imports,
		RuleItems: []RuleItem{
			{
				FnName: "test",
				Effect: EFFECT_TYPE_ALLOW,
				Subs: []Rule{
					{
						FnName: "user_is_admin",
						FnBody: `u.user = "admin"`,
						Desc:   "用户是管理员",
					},
				},
				Rules: []Rule{
					{
						FnName: "menu_match_setting",
						FnBody: `input.env.menu_uri == "setting/list"`,
						Desc:   "菜单匹配",
					},
				},
			},
		},
	}

	t.Run("Gen", func(t *testing.T) {
		t.Log(Gen(r))
	})
}
