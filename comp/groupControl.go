package comp

// GroupControl 表单集合渲染器，能让多个表单在一行显示
type GroupControl BaseRenderer

// NewGroupControl 创建一个新的 GroupControl 实例
func NewGroupControl() GroupControl {
	return GroupControl(make(BaseRenderer)).set("type", "group")
}

// set 设置属性并返回当前对象
func (g GroupControl) set(key string, value interface{}) GroupControl {
	g[key] = value
	return g
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (g GroupControl) AutoFill(value string) GroupControl {
	return g.set("autoFill", value)
}

// Body FormItem 集合
func (g GroupControl) Body(value ...interface{}) GroupControl {
	return g.set("body", value)
}

// ClassName 容器 css 类名
func (g GroupControl) ClassName(value string) GroupControl {
	return g.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (g GroupControl) ClearValueOnHidden(value bool) GroupControl {
	return g.set("clearValueOnHidden", value)
}

// Desc
func (g GroupControl) Desc(value string) GroupControl {
	return g.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (g GroupControl) Description(value string) GroupControl {
	return g.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (g GroupControl) DescriptionClassName(value string) GroupControl {
	return g.set("descriptionClassName", value)
}

// Direction 配置时垂直摆放还是左右摆放。 可选值: horizontal | vertical
func (g GroupControl) Direction(value string) GroupControl {
	return g.set("direction", value)
}

// Disabled 是否禁用
func (g GroupControl) Disabled(value bool) GroupControl {
	return g.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (g GroupControl) DisabledOn(value string) GroupControl {
	return g.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (g GroupControl) EditorSetting(value string) GroupControl {
	return g.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (g GroupControl) ExtraName(value string) GroupControl {
	return g.set("extraName", value)
}

// Gap 间隔 可选值: xs | sm | normal
func (g GroupControl) Gap(value string) GroupControl {
	return g.set("gap", value)
}

// Hidden 是否隐藏
func (g GroupControl) Hidden(value bool) GroupControl {
	return g.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (g GroupControl) HiddenOn(value string) GroupControl {
	return g.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (g GroupControl) Hint(value string) GroupControl {
	return g.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。
func (g GroupControl) Horizontal(value string) GroupControl {
	return g.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (g GroupControl) Id(value string) GroupControl {
	return g.set("id", value)
}

// InitAutoFill
func (g GroupControl) InitAutoFill(value string) GroupControl {
	return g.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (g GroupControl) Inline(value bool) GroupControl {
	return g.set("inline", value)
}

// InputClassName 配置 input className
func (g GroupControl) InputClassName(value string) GroupControl {
	return g.set("inputClassName", value)
}

// Label 描述标题
func (g GroupControl) Label(value string) GroupControl {
	return g.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (g GroupControl) LabelAlign(value string) GroupControl {
	return g.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (g GroupControl) LabelClassName(value string) GroupControl {
	return g.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (g GroupControl) LabelRemark(value string) GroupControl {
	return g.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (g GroupControl) LabelWidth(value string) GroupControl {
	return g.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (g GroupControl) Mode(value string) GroupControl {
	return g.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (g GroupControl) Name(value string) GroupControl {
	return g.set("name", value)
}

// OnEvent 事件动作配置
func (g GroupControl) OnEvent(value string) GroupControl {
	return g.set("onEvent", value)
}

// Placeholder 占位符
func (g GroupControl) Placeholder(value string) GroupControl {
	return g.set("placeholder", value)
}

// ReadOnly 是否只读
func (g GroupControl) ReadOnly(value bool) GroupControl {
	return g.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (g GroupControl) ReadOnlyOn(value string) GroupControl {
	return g.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (g GroupControl) Remark(value string) GroupControl {
	return g.set("remark", value)
}

// Required 是否为必填
func (g GroupControl) Required(value bool) GroupControl {
	return g.set("required", value)
}

// Row
func (g GroupControl) Row(value string) GroupControl {
	return g.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (g GroupControl) SaveImmediately(value bool) GroupControl {
	return g.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (g GroupControl) Size(value string) GroupControl {
	return g.set("size", value)
}

// Static 是否静态展示
func (g GroupControl) Static(value bool) GroupControl {
	return g.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (g GroupControl) StaticClassName(value string) GroupControl {
	return g.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (g GroupControl) StaticInputClassName(value string) GroupControl {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (g GroupControl) StaticLabelClassName(value string) GroupControl {
	return g.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (g GroupControl) StaticOn(value string) GroupControl {
	return g.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (g GroupControl) StaticPlaceholder(value string) GroupControl {
	return g.set("staticPlaceholder", value)
}

// StaticSchema
func (g GroupControl) StaticSchema(value string) GroupControl {
	return g.set("staticSchema", value)
}

// Style 组件样式
func (g GroupControl) Style(value string) GroupControl {
	return g.set("style", value)
}

// SubFormHorizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (g GroupControl) SubFormHorizontal(value string) GroupControl {
	return g.set("subFormHorizontal", value)
}

// SubFormMode 配置子表单项默认的展示方式。 可选值: normal | inline | horizontal
func (g GroupControl) SubFormMode(value string) GroupControl {
	return g.set("subFormMode", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (g GroupControl) SubmitOnChange(value bool) GroupControl {
	return g.set("submitOnChange", value)
}

// TestIdBuilder
func (g GroupControl) TestIdBuilder(value string) GroupControl {
	return g.set("testIdBuilder", value)
}

// Type 表单项类型
func (g GroupControl) Type(value string) GroupControl {
	return g.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (g GroupControl) UseMobileUI(value bool) GroupControl {
	return g.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (g GroupControl) ValidateApi(value string) GroupControl {
	return g.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (g GroupControl) ValidateOnChange(value bool) GroupControl {
	return g.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (g GroupControl) ValidationErrors(value string) GroupControl {
	return g.set("validationErrors", value)
}

// Validations
func (g GroupControl) Validations(value string) GroupControl {
	return g.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (g GroupControl) Value(value string) GroupControl {
	return g.set("value", value)
}

// Visible 是否显示
func (g GroupControl) Visible(value bool) GroupControl {
	return g.set("visible", value)
}

// VisibleOn 是否显示表达式
func (g GroupControl) VisibleOn(value string) GroupControl {
	return g.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (g GroupControl) Width(value string) GroupControl {
	return g.set("width", value)
}
