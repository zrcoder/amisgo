package comp

// control 表单项包裹。文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/control
//
// @version 6.7.0
type control schema

// Control 创建一个新的 Control 实例
func Control() control {
	return make(control).set("type", "control")
}

func (fc control) set(key string, value any) control {
	fc[key] = value
	return fc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (fc control) AutoFill(value string) control {
	return fc.set("autoFill", value)
}

// Body FormItem 内容
func (fc control) Body(value ...any) control {
	return fc.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (fc control) ClassName(value string) control {
	return fc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (fc control) ClearValueOnHidden(value bool) control {
	return fc.set("clearValueOnHidden", value)
}

// Desc
func (fc control) Desc(value string) control {
	return fc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (fc control) Description(value string) control {
	return fc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (fc control) DescriptionClassName(value string) control {
	return fc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (fc control) Disabled(value bool) control {
	return fc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (fc control) DisabledOn(value string) control {
	return fc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (fc control) EditorSetting(value string) control {
	return fc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (fc control) ExtraName(value string) control {
	return fc.set("extraName", value)
}

// Hidden 是否隐藏
func (fc control) Hidden(value bool) control {
	return fc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (fc control) HiddenOn(value string) control {
	return fc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (fc control) Hint(value string) control {
	return fc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (fc control) Horizontal(value string) control {
	return fc.set("horizontal", value)
}

// ID 组件唯一 id，主要用于日志采集
func (fc control) ID(value string) control {
	return fc.set("id", value)
}

// InitAutoFill
func (fc control) InitAutoFill(value string) control {
	return fc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (fc control) Inline(value bool) control {
	return fc.set("inline", value)
}

// InputClassName 配置 input className
func (fc control) InputClassName(value string) control {
	return fc.set("inputClassName", value)
}

// Label 描述标题
func (fc control) Label(value string) control {
	return fc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (fc control) LabelAlign(value string) control {
	return fc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (fc control) LabelClassName(value string) control {
	return fc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (fc control) LabelRemark(value string) control {
	return fc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (fc control) LabelWidth(value string) control {
	return fc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (fc control) Mode(value string) control {
	return fc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (fc control) Name(value string) control {
	return fc.set("name", value)
}

// OnEvent 事件动作配置
func (fc control) OnEvent(value any) control {
	return fc.set("onEvent", value)
}

// Placeholder 占位符
func (fc control) Placeholder(value string) control {
	return fc.set("placeholder", value)
}

// ReadOnly 是否只读
func (fc control) ReadOnly(value bool) control {
	return fc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (fc control) ReadOnlyOn(value string) control {
	return fc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (fc control) Remark(value string) control {
	return fc.set("remark", value)
}

// Required 是否为必填
func (fc control) Required(value bool) control {
	return fc.set("required", value)
}

// Row
func (fc control) Row(value string) control {
	return fc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (fc control) SaveImmediately(value bool) control {
	return fc.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (fc control) Size(value string) control {
	return fc.set("size", value)
}

// Static 是否静态展示
func (fc control) Static(value bool) control {
	return fc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (fc control) StaticClassName(value string) control {
	return fc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (fc control) StaticInputClassName(value string) control {
	return fc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (fc control) StaticLabelClassName(value string) control {
	return fc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (fc control) StaticOn(value string) control {
	return fc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (fc control) StaticPlaceholder(value string) control {
	return fc.set("staticPlaceholder", value)
}

// StaticSchema
func (fc control) StaticSchema(value string) control {
	return fc.set("staticSchema", value)
}

// Style 组件样式
func (fc control) Style(value any) control {
	return fc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (fc control) SubmitOnChange(value bool) control {
	return fc.set("submitOnChange", value)
}

// TestIdBuilder
func (fc control) TestIdBuilder(value string) control {
	return fc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (fc control) UseMobileUI(value bool) control {
	return fc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (fc control) ValidateApi(value string) control {
	return fc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证
func (fc control) ValidateOnChange(value bool) control {
	return fc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (fc control) ValidationErrors(value string) control {
	return fc.set("validationErrors", value)
}

// Validations
func (fc control) Validations(value string) control {
	return fc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的
func (fc control) Value(value string) control {
	return fc.set("value", value)
}

// Visible 是否显示
func (fc control) Visible(value bool) control {
	return fc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (fc control) VisibleOn(value string) control {
	return fc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (fc control) Width(value string) control {
	return fc.set("width", value)
}
