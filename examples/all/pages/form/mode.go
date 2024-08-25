package form

import (
	"github.com/zrcoder/amisgo/comp"
)

var Modes = comp.Page().Title("表单各种展示模式汇总").Remark("展示各种模式的 Form").Body(
	comp.Grid().Columns(
		comp.Form().Title("常规模式").Mode("normal").Body(
			comp.InputEmail().Name("email").Required(true).Placeholder("请输入邮箱").Label("邮箱").Size("full"),
			comp.InputPassword().Name("password").Label("密码").Required(true).Placeholder("请输入密码").Size("full"),
			comp.Checkbox().Name("rememberMe").Label("记住登录"),
			comp.Submit().Label("登录"),
		),
		comp.Form().Title("常规模式 input md 尺寸").Mode("normal").Body(
			comp.InputEmail().Name("email").Required(true).Placeholder("请输入邮箱").Label("邮箱").
				Size("md").Remark("xxxx").Hint("bla bla bla"),
			comp.InputPassword().Name("password").Label("密码").Required(true).Placeholder("请输入密码").Size("md"),
			comp.Checkbox().Name("rememberMe").Label("记住登录"),
			comp.Submit().Label("登录"),
		),
	),
	comp.Grid().Columns(
		comp.Form().Title("水平模式，左右摆放，左右比率分配").Mode("horizontal").AutoFocus(true).Horizontal(
			comp.Horizontal().Left("col-sm-2").Right("col-sm-10").Offset("col-sm-offset-2"),
		).Body(
			comp.InputEmail().Name("email").Required(true).Placeholder("请输入邮箱").Label("邮箱").Desc("表单描述文字"),
			comp.InputPassword().Name("password").Label("密码").Required(true).Placeholder("请输入密码"),
			comp.Checkbox().Name("rememberMe").Label("记住登录"),
			comp.Control().Body(
				comp.Submit().Label("Sumbit"),
			),
		),
		comp.Form().Title("水平模式，左右摆放 左侧固定宽度 input md 尺寸").Mode("horizontal").AutoFocus(false).Horizontal(comp.Horizontal().LeftFixed("xs")).Body(
			comp.InputEmail().Name("email").Required(true).Placeholder("请输入邮箱").Label("邮箱").
				Size("md").Remark("xxxx").Hint("bla bla bla"),
			comp.InputPassword().Name("password").Label("密码").Required(true).Placeholder("请输入密码").Size("md"),
			comp.Checkbox().Name("rememberMe").Label("记住登录"),
			comp.Control().Body(
				comp.Submit().Label("Sumbit"),
			),
		),
	),
	comp.Form().ClassName("m-b").Body(
		comp.Property().Title("机器配置").Items(
			comp.PropertyItem().Label("cpu").Content(
				comp.Select().Name("cpu").Value("1").Options(
					comp.Option().Label("1 core").Value("1"),
					comp.Option().Label("4 core").Value("4"),
					comp.Option().Label("8 core").Value("8"),
				),
			),
			comp.PropertyItem().Label("memory").Content("4G"),
			comp.PropertyItem().Label("disk").Content("80G"),
			comp.PropertyItem().Label("network").Content("4M").Span(2),
			comp.PropertyItem().Label("IDC").Content("beijing"),
			comp.PropertyItem().Label("Note").Content(
				comp.Textarea().Required(true).Name("note").Placeholder("Enter..."),
			).Span(3),
		),
	).Actions(comp.Submit().Label("Submit")),
	comp.Form().Title("内联模式").Mode("inline").AutoFocus(false).Body(
		comp.InputEmail().Name("email").Placeholder("Enter Email").Label("Email"),
		comp.InputPassword().Name("password").Placeholder("Enter Password").Remark("Bla bla bla"),
		comp.Checkbox().Name("rememberMe").Label("Remember me"),
		comp.Submit().Label("Login"),
		comp.Button().Label("Export").Url("https://www.baidu.com").Level("success"),
	),
	comp.Form().Title("常规模式下用数组包起来还能控制一行显示多个").Mode("normal").AutoFocus(false).Body(
		comp.InputText().Name("name").Placeholder("请输入...").Label("名字").Size("full"),
		comp.Divider(),
		comp.Group().Body(
			comp.InputEmail().Name("email").Placeholder("输入邮箱").Label("邮箱").Size("full"),
			comp.InputPassword().Name("password").Label("密码").Placeholder("请输入密码").Size("full"),
		),
		comp.Divider(),
		comp.Group().Body(
			comp.InputEmail().Name("email2").Mode("inlie").Placeholder("输入邮箱").Label("邮箱").Size("full"),
			comp.InputPassword().Name("password2").Mode("inlie").Label("密码").Placeholder("请输入密码").Size("full"),
		),
		comp.Divider(),
		comp.Group().Body(
			comp.InputEmail().Name("email3").Mode("inlie").Placeholder("输入邮箱").Label("邮箱").Size("full").ColumnClassName("v-bottom"),
			comp.InputPassword().Name("password3").Mode("inlie").Label("密码").Placeholder("请输入密码").Size("full"),
		),
		comp.Divider(),
		comp.Checkbox().Name("rememberMe").Label("记住我"),
		comp.Submit().Label("提交"),
	),
	comp.Form().Title("水平模式用数组包起来也能控制一行显示多个").Mode("horizontal").AutoFocus(false).Body(
		comp.InputEmail().Name("email").Label("Email").Size("full"),
		comp.Divider(),
		comp.Group().Body(
			comp.InputEmail().Name("email3").Mode("inlie").Placeholder("输入邮箱").Label("邮箱").Size("full").ColumnClassName("v-bottom"),
			comp.InputPassword().Name("password3").Mode("inlie").Label("密码").Placeholder("请输入密码").Size("full"),
		),
	),
	comp.Form().Title("inline 模式用数组包起来也能控制一行显示多个").Mode("inline").AutoFocus(false).Body(
		comp.InputEmail().Name("email").Label("Email").Size("full"),
		comp.Checkbox().Name("rememberMe").Label("记住我"),
		comp.Group().Body(
			comp.InputEmail().Name("email3").Mode("inlie").Placeholder("输入邮箱").Label("邮箱").Size("full").ColumnClassName("v-bottom"),
			comp.InputPassword().Name("password3").Mode("inlie").Label("密码").Placeholder("请输入密码").Size("full"),
		),
	),
)
