package comp

// RatingControl 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/rating
//
// @author  slowlyo
// @version 6.7.0
type RatingControl Schema

// NewRatingControl 创建一个新的 RatingControl 实例
func NewRatingControl() RatingControl {
	return RatingControl{}.set("type", "input-rating")
}

func (rc RatingControl) set(key string, value interface{}) RatingControl {
	rc[key] = value
	return rc
}

// AllowClear 是否允许再次点击后清除
func (rc RatingControl) AllowClear(value bool) RatingControl {
	return rc.set("allowClear", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc RatingControl) AutoFill(value string) RatingControl {
	return rc.set("autoFill", value)
}

// Char 自定义字符
func (rc RatingControl) Char(value string) RatingControl {
	return rc.set("char", value)
}

// CharClassName 自定义字符类名
func (rc RatingControl) CharClassName(value string) RatingControl {
	return rc.set("charClassName", value)
}

// ClassName 容器 css 类名
func (rc RatingControl) ClassName(value string) RatingControl {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (rc RatingControl) ClearValueOnHidden(value bool) RatingControl {
	return rc.set("clearValueOnHidden", value)
}

// Colors 星星被选中的颜色
func (rc RatingControl) Colors(value string) RatingControl {
	return rc.set("colors", value)
}

// Count 分数
func (rc RatingControl) Count(value string) RatingControl {
	return rc.set("count", value)
}

// Desc
func (rc RatingControl) Desc(value string) RatingControl {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (rc RatingControl) Description(value string) RatingControl {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (rc RatingControl) DescriptionClassName(value string) RatingControl {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc RatingControl) Disabled(value bool) RatingControl {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (rc RatingControl) DisabledOn(value string) RatingControl {
	return rc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc RatingControl) EditorSetting(value string) RatingControl {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (rc RatingControl) ExtraName(value string) RatingControl {
	return rc.set("extraName", value)
}

// Half 允许半颗星
func (rc RatingControl) Half(value bool) RatingControl {
	return rc.set("half", value)
}

// Hidden 是否隐藏
func (rc RatingControl) Hidden(value bool) RatingControl {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (rc RatingControl) HiddenOn(value string) RatingControl {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc RatingControl) Hint(value string) RatingControl {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (rc RatingControl) Horizontal(value string) RatingControl {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id
func (rc RatingControl) Id(value string) RatingControl {
	return rc.set("id", value)
}

// InactiveColor 未被选中的星星的颜色
func (rc RatingControl) InactiveColor(value string) RatingControl {
	return rc.set("inactiveColor", value)
}

// InitAutoFill
func (rc RatingControl) InitAutoFill(value string) RatingControl {
	return rc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (rc RatingControl) Inline(value bool) RatingControl {
	return rc.set("inline", value)
}

// InputClassName 配置 input className
func (rc RatingControl) InputClassName(value string) RatingControl {
	return rc.set("inputClassName", value)
}

// Label 描述标题
func (rc RatingControl) Label(value string) RatingControl {
	return rc.set("label", value)
}

// LabelAlign 描述标题
func (rc RatingControl) LabelAlign(value string) RatingControl {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc RatingControl) LabelClassName(value string) RatingControl {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (rc RatingControl) LabelRemark(value string) RatingControl {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (rc RatingControl) LabelWidth(value string) RatingControl {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (rc RatingControl) Mode(value string) RatingControl {
	return rc.set("mode", value)
}

// Name 字段名
func (rc RatingControl) Name(value string) RatingControl {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc RatingControl) OnEvent(value string) RatingControl {
	return rc.set("onEvent", value)
}

// Placeholder 占位符
func (rc RatingControl) Placeholder(value string) RatingControl {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc RatingControl) ReadOnly(value bool) RatingControl {
	return rc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (rc RatingControl) ReadOnlyOn(value string) RatingControl {
	return rc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (rc RatingControl) Remark(value string) RatingControl {
	return rc.set("remark", value)
}

// Required 是否为必填
func (rc RatingControl) Required(value bool) RatingControl {
	return rc.set("required", value)
}

// Row
func (rc RatingControl) Row(value string) RatingControl {
	return rc.set("row", value)
}

// SaveImmediately 是否立即保存
func (rc RatingControl) SaveImmediately(value bool) RatingControl {
	return rc.set("saveImmediately", value)
}

// Size 表单项大小
func (rc RatingControl) Size(value string) RatingControl {
	return rc.set("size", value)
}

// Static 是否静态展示
func (rc RatingControl) Static(value bool) RatingControl {
	return rc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (rc RatingControl) StaticClassName(value string) RatingControl {
	return rc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (rc RatingControl) StaticInputClassName(value string) RatingControl {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (rc RatingControl) StaticLabelClassName(value string) RatingControl {
	return rc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (rc RatingControl) StaticOn(value string) RatingControl {
	return rc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rc RatingControl) StaticPlaceholder(value string) RatingControl {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc RatingControl) StaticSchema(value string) RatingControl {
	return rc.set("staticSchema", value)
}

// Style 组件样式
func (rc RatingControl) Style(value string) RatingControl {
	return rc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (rc RatingControl) SubmitOnChange(value bool) RatingControl {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc RatingControl) TestIdBuilder(value string) RatingControl {
	return rc.set("testIdBuilder", value)
}

// TextClassName 自定义文字类名
func (rc RatingControl) TextClassName(value string) RatingControl {
	return rc.set("textClassName", value)
}

// TextPosition 文字的位置
func (rc RatingControl) TextPosition(value string) RatingControl {
	return rc.set("textPosition", value)
}

// Texts 星星被选中时的提示文字
func (rc RatingControl) Texts(value string) RatingControl {
	return rc.set("texts", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rc RatingControl) UseMobileUI(value bool) RatingControl {
	return rc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (rc RatingControl) ValidateApi(value string) RatingControl {
	return rc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (rc RatingControl) ValidateOnChange(value bool) RatingControl {
	return rc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (rc RatingControl) ValidationErrors(value string) RatingControl {
	return rc.set("validationErrors", value)
}

// Validations
func (rc RatingControl) Validations(value string) RatingControl {
	return rc.set("validations", value)
}

// Value 默认值
func (rc RatingControl) Value(value string) RatingControl {
	return rc.set("value", value)
}

// Visible 是否显示
func (rc RatingControl) Visible(value bool) RatingControl {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (rc RatingControl) VisibleOn(value string) RatingControl {
	return rc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (rc RatingControl) Width(value string) RatingControl {
	return rc.set("width", value)
}
