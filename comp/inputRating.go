package comp

// inputRating 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/rating

// @version 6.7.0
type inputRating schema

func InputRating() inputRating {
	return inputRating{}.set("type", "input-rating")
}

func (rc inputRating) set(key string, value any) inputRating {
	rc[key] = value
	return rc
}

// AllowClear 是否允许再次点击后清除
func (rc inputRating) AllowClear(value bool) inputRating {
	return rc.set("allowClear", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc inputRating) AutoFill(value string) inputRating {
	return rc.set("autoFill", value)
}

// Char 自定义字符
func (rc inputRating) Char(value string) inputRating {
	return rc.set("char", value)
}

// CharClassName 自定义字符类名
func (rc inputRating) CharClassName(value string) inputRating {
	return rc.set("charClassName", value)
}

// ClassName 容器 css 类名
func (rc inputRating) ClassName(value string) inputRating {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (rc inputRating) ClearValueOnHidden(value bool) inputRating {
	return rc.set("clearValueOnHidden", value)
}

// Colors 星星被选中的颜色
func (rc inputRating) Colors(value string) inputRating {
	return rc.set("colors", value)
}

// Count 分数
func (rc inputRating) Count(value string) inputRating {
	return rc.set("count", value)
}

// Desc
func (rc inputRating) Desc(value string) inputRating {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (rc inputRating) Description(value string) inputRating {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (rc inputRating) DescriptionClassName(value string) inputRating {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc inputRating) Disabled(value bool) inputRating {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (rc inputRating) DisabledOn(value string) inputRating {
	return rc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc inputRating) EditorSetting(value string) inputRating {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (rc inputRating) ExtraName(value string) inputRating {
	return rc.set("extraName", value)
}

// Half 允许半颗星
func (rc inputRating) Half(value bool) inputRating {
	return rc.set("half", value)
}

// Hidden 是否隐藏
func (rc inputRating) Hidden(value bool) inputRating {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (rc inputRating) HiddenOn(value string) inputRating {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc inputRating) Hint(value string) inputRating {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (rc inputRating) Horizontal(value string) inputRating {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id
func (rc inputRating) Id(value string) inputRating {
	return rc.set("id", value)
}

// InactiveColor 未被选中的星星的颜色
func (rc inputRating) InactiveColor(value string) inputRating {
	return rc.set("inactiveColor", value)
}

// InitAutoFill
func (rc inputRating) InitAutoFill(value string) inputRating {
	return rc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (rc inputRating) Inline(value bool) inputRating {
	return rc.set("inline", value)
}

// InputClassName 配置 input className
func (rc inputRating) InputClassName(value string) inputRating {
	return rc.set("inputClassName", value)
}

// Label 描述标题
func (rc inputRating) Label(value string) inputRating {
	return rc.set("label", value)
}

// LabelAlign 描述标题
func (rc inputRating) LabelAlign(value string) inputRating {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc inputRating) LabelClassName(value string) inputRating {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (rc inputRating) LabelRemark(value string) inputRating {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (rc inputRating) LabelWidth(value string) inputRating {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (rc inputRating) Mode(value string) inputRating {
	return rc.set("mode", value)
}

// Name 字段名
func (rc inputRating) Name(value string) inputRating {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc inputRating) OnEvent(value any) inputRating {
	return rc.set("onEvent", value)
}

// Placeholder 占位符
func (rc inputRating) Placeholder(value string) inputRating {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc inputRating) ReadOnly(value bool) inputRating {
	return rc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (rc inputRating) ReadOnlyOn(value string) inputRating {
	return rc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (rc inputRating) Remark(value string) inputRating {
	return rc.set("remark", value)
}

// Required 是否为必填
func (rc inputRating) Required(value bool) inputRating {
	return rc.set("required", value)
}

// Row
func (rc inputRating) Row(value string) inputRating {
	return rc.set("row", value)
}

// SaveImmediately 是否立即保存
func (rc inputRating) SaveImmediately(value bool) inputRating {
	return rc.set("saveImmediately", value)
}

// Size 表单项大小
func (rc inputRating) Size(value string) inputRating {
	return rc.set("size", value)
}

// Static 是否静态展示
func (rc inputRating) Static(value bool) inputRating {
	return rc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (rc inputRating) StaticClassName(value string) inputRating {
	return rc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (rc inputRating) StaticInputClassName(value string) inputRating {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (rc inputRating) StaticLabelClassName(value string) inputRating {
	return rc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (rc inputRating) StaticOn(value string) inputRating {
	return rc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rc inputRating) StaticPlaceholder(value string) inputRating {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc inputRating) StaticSchema(value string) inputRating {
	return rc.set("staticSchema", value)
}

// Style 组件样式
func (rc inputRating) Style(value any) inputRating {
	return rc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (rc inputRating) SubmitOnChange(value bool) inputRating {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc inputRating) TestIdBuilder(value string) inputRating {
	return rc.set("testIdBuilder", value)
}

// TextClassName 自定义文字类名
func (rc inputRating) TextClassName(value string) inputRating {
	return rc.set("textClassName", value)
}

// TextPosition 文字的位置
func (rc inputRating) TextPosition(value string) inputRating {
	return rc.set("textPosition", value)
}

// Texts 星星被选中时的提示文字
func (rc inputRating) Texts(value string) inputRating {
	return rc.set("texts", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rc inputRating) UseMobileUI(value bool) inputRating {
	return rc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (rc inputRating) ValidateApi(value string) inputRating {
	return rc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (rc inputRating) ValidateOnChange(value bool) inputRating {
	return rc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (rc inputRating) ValidationErrors(value string) inputRating {
	return rc.set("validationErrors", value)
}

// Validations
func (rc inputRating) Validations(value string) inputRating {
	return rc.set("validations", value)
}

// Value 默认值
func (rc inputRating) Value(value string) inputRating {
	return rc.set("value", value)
}

// Visible 是否显示
func (rc inputRating) Visible(value bool) inputRating {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (rc inputRating) VisibleOn(value string) inputRating {
	return rc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (rc inputRating) Width(value string) inputRating {
	return rc.set("width", value)
}
