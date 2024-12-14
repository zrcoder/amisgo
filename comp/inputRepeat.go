package comp

// inputRepeat 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/repeat

type inputRepeat schema

// InputRepeat 创建一个新的 RepeatControl 实例
func InputRepeat() inputRepeat {
	return inputRepeat{}.set("type", "input-repeat")
}

func (rc inputRepeat) set(key string, value any) inputRepeat {
	rc[key] = value
	return rc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc inputRepeat) AutoFill(value string) inputRepeat {
	return rc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc inputRepeat) ClassName(value string) inputRepeat {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (rc inputRepeat) ClearValueOnHidden(value bool) inputRepeat {
	return rc.set("clearValueOnHidden", value)
}

// Desc
func (rc inputRepeat) Desc(value string) inputRepeat {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (rc inputRepeat) Description(value string) inputRepeat {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (rc inputRepeat) DescriptionClassName(value string) inputRepeat {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc inputRepeat) Disabled(value bool) inputRepeat {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rc inputRepeat) DisabledOn(value string) inputRepeat {
	return rc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc inputRepeat) EditorSetting(value string) inputRepeat {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (rc inputRepeat) ExtraName(value string) inputRepeat {
	return rc.set("extraName", value)
}

// Hidden 是否隐藏
func (rc inputRepeat) Hidden(value bool) inputRepeat {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rc inputRepeat) HiddenOn(value string) inputRepeat {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc inputRepeat) Hint(value string) inputRepeat {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。
func (rc inputRepeat) Horizontal(value string) inputRepeat {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rc inputRepeat) ID(value string) inputRepeat {
	return rc.set("id", value)
}

// InitAutoFill
func (rc inputRepeat) InitAutoFill(value string) inputRepeat {
	return rc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (rc inputRepeat) Inline(value bool) inputRepeat {
	return rc.set("inline", value)
}

// InputClassName 配置 input className
func (rc inputRepeat) InputClassName(value string) inputRepeat {
	return rc.set("inputClassName", value)
}

// Label 描述标题
func (rc inputRepeat) Label(value string) inputRepeat {
	return rc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (rc inputRepeat) LabelAlign(value string) inputRepeat {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc inputRepeat) LabelClassName(value string) inputRepeat {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (rc inputRepeat) LabelRemark(value string) inputRepeat {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (rc inputRepeat) LabelWidth(value string) inputRepeat {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (rc inputRepeat) Mode(value string) inputRepeat {
	return rc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (rc inputRepeat) Name(value string) inputRepeat {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc inputRepeat) OnEvent(value any) inputRepeat {
	return rc.set("onEvent", value)
}

// Options
func (rc inputRepeat) Options(value ...option) inputRepeat {
	return rc.set("options", value)
}

// Placeholder 占位符
func (rc inputRepeat) Placeholder(value string) inputRepeat {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc inputRepeat) ReadOnly(value bool) inputRepeat {
	return rc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (rc inputRepeat) ReadOnlyOn(value string) inputRepeat {
	return rc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (rc inputRepeat) Remark(value string) inputRepeat {
	return rc.set("remark", value)
}

// Required 是否为必填
func (rc inputRepeat) Required(value bool) inputRepeat {
	return rc.set("required", value)
}

// Row
func (rc inputRepeat) Row(value string) inputRepeat {
	return rc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (rc inputRepeat) SaveImmediately(value bool) inputRepeat {
	return rc.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (rc inputRepeat) Size(value string) inputRepeat {
	return rc.set("size", value)
}

// Static 是否静态展示
func (rc inputRepeat) Static(value bool) inputRepeat {
	return rc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc inputRepeat) StaticClassName(value string) inputRepeat {
	return rc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc inputRepeat) StaticInputClassName(value string) inputRepeat {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc inputRepeat) StaticLabelClassName(value string) inputRepeat {
	return rc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc inputRepeat) StaticOn(value string) inputRepeat {
	return rc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rc inputRepeat) StaticPlaceholder(value string) inputRepeat {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc inputRepeat) StaticSchema(value string) inputRepeat {
	return rc.set("staticSchema", value)
}

// Style 组件样式
func (rc inputRepeat) Style(value any) inputRepeat {
	return rc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (rc inputRepeat) SubmitOnChange(value bool) inputRepeat {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc inputRepeat) TestIdBuilder(value string) inputRepeat {
	return rc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rc inputRepeat) UseMobileUI(value bool) inputRepeat {
	return rc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (rc inputRepeat) ValidateApi(value string) inputRepeat {
	return rc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (rc inputRepeat) ValidateOnChange(value bool) inputRepeat {
	return rc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (rc inputRepeat) ValidationErrors(value string) inputRepeat {
	return rc.set("validationErrors", value)
}

// Validations
func (rc inputRepeat) Validations(value string) inputRepeat {
	return rc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (rc inputRepeat) Value(value string) inputRepeat {
	return rc.set("value", value)
}

// Visible 是否显示
func (rc inputRepeat) Visible(value bool) inputRepeat {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc inputRepeat) VisibleOn(value string) inputRepeat {
	return rc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (rc inputRepeat) Width(value string) inputRepeat {
	return rc.set("width", value)
}
