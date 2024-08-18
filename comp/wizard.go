package comp

// Wizard 表单向导 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/wizard
//
// @author  slowlyo
// @version 6.7.0
type Wizard Schema

// NewWizard 创建一个新的 Wizard 实例
func NewWizard() Wizard {
	return Wizard{}.set("type", "wizard")
}

func (w Wizard) set(key string, value interface{}) Wizard {
	w[key] = value
	return w
}

// ActionClassName 配置按钮 className (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wizard) ActionClassName(value string) Wizard {
	return w.set("actionClassName", value)
}

// ActionFinishLabel 完成按钮的文字描述
func (w Wizard) ActionFinishLabel(value string) Wizard {
	return w.set("actionFinishLabel", value)
}

// ActionNextLabel 下一步按钮的文字描述
func (w Wizard) ActionNextLabel(value string) Wizard {
	return w.set("actionNextLabel", value)
}

// ActionNextSaveLabel 下一步并且保存按钮的文字描述
func (w Wizard) ActionNextSaveLabel(value string) Wizard {
	return w.set("actionNextSaveLabel", value)
}

// ActionPrevLabel 上一步按钮的文字描述
func (w Wizard) ActionPrevLabel(value string) Wizard {
	return w.set("actionPrevLabel", value)
}

// AffixFooter 是否将底部按钮固定在底部。
func (w Wizard) AffixFooter(value string) Wizard {
	return w.set("affixFooter", value)
}

// Api Wizard 用来保存数据的 api。 [详情](https://baidu.github.io/amis/docs/api#wizard) (Wizard 用来保存数据的 api。 [详情](https://baidu.github.io/amis/docs/api#wizard))
func (w Wizard) Api(value string) Wizard {
	return w.set("api", value)
}

// BodyClassName 表单区域css类
func (w Wizard) BodyClassName(value string) Wizard {
	return w.set("bodyClassName", value)
}

// BulkSubmit 是否合并后再提交
func (w Wizard) BulkSubmit(value bool) Wizard {
	return w.set("bulkSubmit", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wizard) ClassName(value string) Wizard {
	return w.set("className", value)
}

// Disabled 是否禁用
func (w Wizard) Disabled(value bool) Wizard {
	return w.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wizard) DisabledOn(value string) Wizard {
	return w.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (w Wizard) EditorSetting(value string) Wizard {
	return w.set("editorSetting", value)
}

// FooterClassName 底部操作栏的css类
func (w Wizard) FooterClassName(value string) Wizard {
	return w.set("footerClassName", value)
}

// Hidden 是否隐藏
func (w Wizard) Hidden(value bool) Wizard {
	return w.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wizard) HiddenOn(value string) Wizard {
	return w.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (w Wizard) Id(value string) Wizard {
	return w.set("id", value)
}

// InitApi Wizard 用来获取初始数据的 api。 (Wizard 用来获取初始数据的 api。)
func (w Wizard) InitApi(value string) Wizard {
	return w.set("initApi", value)
}

// LoadingConfig
func (w Wizard) LoadingConfig(value string) Wizard {
	return w.set("loadingConfig", value)
}

// Mode 展示模式 可选值: vertical | horizontal
func (w Wizard) Mode(value string) Wizard {
	return w.set("mode", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (w Wizard) Name(value string) Wizard {
	return w.set("name", value)
}

// OnEvent 事件动作配置
func (w Wizard) OnEvent(value string) Wizard {
	return w.set("onEvent", value)
}

// ReadOnly 是否为只读模式。
func (w Wizard) ReadOnly(value bool) Wizard {
	return w.set("readOnly", value)
}

// Redirect 保存完后，可以指定跳转地址，支持相对路径和组内绝对路径，同时可以通过 $xxx 使用变量
func (w Wizard) Redirect(value string) Wizard {
	return w.set("redirect", value)
}

// Reload 配置刷新动作，这个动作通常在完成渲染器本省的固定动作后出发。一般用来配置目标组件的 name 属性。多个目标可以用逗号隔开。当目标是 windows 时表示刷新整个页面。刷新目标的同时还支持传递参数如： `foo?a=${a}&b=${b},boo?c=${c}`
func (w Wizard) Reload(value string) Wizard {
	return w.set("reload", value)
}

// StartStep
func (w Wizard) StartStep(value string) Wizard {
	return w.set("startStep", value)
}

// Static 是否静态展示
func (w Wizard) Static(value bool) Wizard {
	return w.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wizard) StaticClassName(value string) Wizard {
	return w.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wizard) StaticInputClassName(value string) Wizard {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wizard) StaticLabelClassName(value string) Wizard {
	return w.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wizard) StaticOn(value string) Wizard {
	return w.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (w Wizard) StaticPlaceholder(value string) Wizard {
	return w.set("staticPlaceholder", value)
}

// StaticSchema
func (w Wizard) StaticSchema(value string) Wizard {
	return w.set("staticSchema", value)
}

// StepClassName step + body区域css类
func (w Wizard) StepClassName(value string) Wizard {
	return w.set("stepClassName", value)
}

// Steps
func (w Wizard) Steps(value string) Wizard {
	return w.set("steps", value)
}

// StepsClassName 步骤条区域css类
func (w Wizard) StepsClassName(value string) Wizard {
	return w.set("stepsClassName", value)
}

// Style 组件样式
func (w Wizard) Style(value string) Wizard {
	return w.set("style", value)
}

// Target 默认表单提交自己会通过发送 api 保存数据，但是也可以设定另外一个 form 的 name 值，或者另外一个 `CRUD` 模型的 name 值。如果 target 目标是一个 `Form` ，则目标 `Form` 会重新触发 `initApi` 和 `schemaApi`，api 可以拿到当前 form 数据。如果目标是一个 `CRUD` 模型，则目标模型会重新触发搜索，参数为当前 Form 数据。
func (w Wizard) Target(value string) Wizard {
	return w.set("target", value)
}

// TestIdBuilder
func (w Wizard) TestIdBuilder(value string) Wizard {
	return w.set("testIdBuilder", value)
}

// Testid
func (w Wizard) Testid(value string) Wizard {
	return w.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (w Wizard) UseMobileUI(value bool) Wizard {
	return w.set("useMobileUI", value)
}

// Visible 是否显示
func (w Wizard) Visible(value bool) Wizard {
	return w.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wizard) VisibleOn(value string) Wizard {
	return w.set("visibleOn", value)
}

// WrapWithPanel 是否用panel包裹
func (w Wizard) WrapWithPanel(value bool) Wizard {
	return w.set("wrapWithPanel", value)
}
