package comp

// group 表单集合渲染器，能让多个表单在一行显示
type group schema

// Group 创建一个新的 GroupControl 实例
func Group() group {
	return make(group).set("type", "group")
}

func (g group) set(key string, value any) group {
	g[key] = value
	return g
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (g group) AutoFill(value string) group {
	return g.set("autoFill", value)
}

// Body FormItem 集合
func (g group) Body(value ...any) group {
	return g.set("body", value)
}

// ClassName 容器 css 类名
func (g group) ClassName(value string) group {
	return g.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (g group) ClearValueOnHidden(value bool) group {
	return g.set("clearValueOnHidden", value)
}

// Desc
func (g group) Desc(value string) group {
	return g.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (g group) Description(value string) group {
	return g.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (g group) DescriptionClassName(value string) group {
	return g.set("descriptionClassName", value)
}

// Direction 配置时垂直摆放还是左右摆放。 可选值: horizontal | vertical
func (g group) Direction(value string) group {
	return g.set("direction", value)
}

// Disabled 是否禁用
func (g group) Disabled(value bool) group {
	return g.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (g group) DisabledOn(value string) group {
	return g.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (g group) EditorSetting(value string) group {
	return g.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (g group) ExtraName(value string) group {
	return g.set("extraName", value)
}

// Gap 间隔 可选值: xs | sm | normal
func (g group) Gap(value string) group {
	return g.set("gap", value)
}

// Hidden 是否隐藏
func (g group) Hidden(value bool) group {
	return g.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (g group) HiddenOn(value string) group {
	return g.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (g group) Hint(value string) group {
	return g.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。
func (g group) Horizontal(value string) group {
	return g.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (g group) ID(value string) group {
	return g.set("id", value)
}

// InitAutoFill
func (g group) InitAutoFill(value string) group {
	return g.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (g group) Inline(value bool) group {
	return g.set("inline", value)
}

// InputClassName 配置 input className
func (g group) InputClassName(value string) group {
	return g.set("inputClassName", value)
}

// Label 描述标题
func (g group) Label(value string) group {
	return g.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (g group) LabelAlign(value string) group {
	return g.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (g group) LabelClassName(value string) group {
	return g.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (g group) LabelRemark(value string) group {
	return g.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (g group) LabelWidth(value string) group {
	return g.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (g group) Mode(value string) group {
	return g.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (g group) Name(value string) group {
	return g.set("name", value)
}

// OnEvent 事件动作配置
func (g group) OnEvent(value any) group {
	return g.set("onEvent", value)
}

// Placeholder 占位符
func (g group) Placeholder(value string) group {
	return g.set("placeholder", value)
}

// ReadOnly 是否只读
func (g group) ReadOnly(value bool) group {
	return g.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (g group) ReadOnlyOn(value string) group {
	return g.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (g group) Remark(value string) group {
	return g.set("remark", value)
}

// Required 是否为必填
func (g group) Required(value bool) group {
	return g.set("required", value)
}

// Row
func (g group) Row(value string) group {
	return g.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (g group) SaveImmediately(value bool) group {
	return g.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (g group) Size(value string) group {
	return g.set("size", value)
}

// Static 是否静态展示
func (g group) Static(value bool) group {
	return g.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (g group) StaticClassName(value string) group {
	return g.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (g group) StaticInputClassName(value string) group {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (g group) StaticLabelClassName(value string) group {
	return g.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (g group) StaticOn(value string) group {
	return g.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (g group) StaticPlaceholder(value string) group {
	return g.set("staticPlaceholder", value)
}

// StaticSchema
func (g group) StaticSchema(value string) group {
	return g.set("staticSchema", value)
}

// Style 组件样式
func (g group) Style(value any) group {
	return g.set("style", value)
}

// SubFormHorizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (g group) SubFormHorizontal(value string) group {
	return g.set("subFormHorizontal", value)
}

// SubFormMode 配置子表单项默认的展示方式。 可选值: normal | inline | horizontal
func (g group) SubFormMode(value string) group {
	return g.set("subFormMode", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (g group) SubmitOnChange(value bool) group {
	return g.set("submitOnChange", value)
}

// TestIdBuilder
func (g group) TestIdBuilder(value string) group {
	return g.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (g group) UseMobileUI(value bool) group {
	return g.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (g group) ValidateApi(value string) group {
	return g.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (g group) ValidateOnChange(value bool) group {
	return g.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (g group) ValidationErrors(value string) group {
	return g.set("validationErrors", value)
}

// Validations
func (g group) Validations(value string) group {
	return g.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (g group) Value(value string) group {
	return g.set("value", value)
}

// Visible 是否显示
func (g group) Visible(value bool) group {
	return g.set("visible", value)
}

// VisibleOn 是否显示表达式
func (g group) VisibleOn(value string) group {
	return g.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (g group) Width(value string) group {
	return g.set("width", value)
}
