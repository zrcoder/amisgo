package comp

// SubFormControl 子表单
//
// @version 6.7.0
type SubFormControl Schema

// NewSubFormControl 创建一个新的 SubFormControl 实例
func NewSubFormControl() SubFormControl {
	return SubFormControl{}.set("type", "input-sub-form")
}

func (s SubFormControl) set(key string, value interface{}) SubFormControl {
	s[key] = value
	return s
}

// AddButtonClassName 新增按钮 CSS 类名 (css类名，配置字符串，或者对象。)
func (s SubFormControl) AddButtonClassName(value string) SubFormControl {
	return s.set("addButtonClassName", value)
}

// AddButtonText 新增按钮文字
func (s SubFormControl) AddButtonText(value string) SubFormControl {
	return s.set("addButtonText", value)
}

// Addable 是否可新增
func (s SubFormControl) Addable(value bool) SubFormControl {
	return s.set("addable", value)
}

// AutoFill 自动填充
func (s SubFormControl) AutoFill(value string) SubFormControl {
	return s.set("autoFill", value)
}

// BtnLabel 按钮默认名称
func (s SubFormControl) BtnLabel(value string) SubFormControl {
	return s.set("btnLabel", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。)
func (s SubFormControl) ClassName(value string) SubFormControl {
	return s.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (s SubFormControl) ClearValueOnHidden(value bool) SubFormControl {
	return s.set("clearValueOnHidden", value)
}

// Desc 描述内容
func (s SubFormControl) Desc(value string) SubFormControl {
	return s.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (s SubFormControl) Description(value string) SubFormControl {
	return s.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (s SubFormControl) DescriptionClassName(value string) SubFormControl {
	return s.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (s SubFormControl) Disabled(value bool) SubFormControl {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s SubFormControl) DisabledOn(value string) SubFormControl {
	return s.set("disabledOn", value)
}

// Draggable 是否可拖拽排序
func (s SubFormControl) Draggable(value bool) SubFormControl {
	return s.set("draggable", value)
}

// DraggableTip 拖拽提示信息
func (s SubFormControl) DraggableTip(value string) SubFormControl {
	return s.set("draggableTip", value)
}

// EditorSetting 编辑器配置
func (s SubFormControl) EditorSetting(value string) SubFormControl {
	return s.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (s SubFormControl) ExtraName(value string) SubFormControl {
	return s.set("extraName", value)
}

// Form 子表单详情
func (s SubFormControl) Form(value string) SubFormControl {
	return s.set("form", value)
}

// Hidden 是否隐藏
func (s SubFormControl) Hidden(value bool) SubFormControl {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s SubFormControl) HiddenOn(value string) SubFormControl {
	return s.set("hiddenOn", value)
}

// Hint 输入提示
func (s SubFormControl) Hint(value string) SubFormControl {
	return s.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (s SubFormControl) Horizontal(value string) SubFormControl {
	return s.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s SubFormControl) Id(value string) SubFormControl {
	return s.set("id", value)
}

// InitAutoFill 初始化自动填充
func (s SubFormControl) InitAutoFill(value string) SubFormControl {
	return s.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (s SubFormControl) Inline(value bool) SubFormControl {
	return s.set("inline", value)
}

// InputClassName 配置 input className
func (s SubFormControl) InputClassName(value string) SubFormControl {
	return s.set("inputClassName", value)
}

// ItemClassName 值元素的类名
func (s SubFormControl) ItemClassName(value string) SubFormControl {
	return s.set("itemClassName", value)
}

// ItemsClassName 值列表元素的类名
func (s SubFormControl) ItemsClassName(value string) SubFormControl {
	return s.set("itemsClassName", value)
}

// Label 描述标题
func (s SubFormControl) Label(value string) SubFormControl {
	return s.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (s SubFormControl) LabelAlign(value string) SubFormControl {
	return s.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (s SubFormControl) LabelClassName(value string) SubFormControl {
	return s.set("labelClassName", value)
}

// LabelField 当值中存在这个字段，则按钮名称将使用此字段的值来展示
func (s SubFormControl) LabelField(value string) SubFormControl {
	return s.set("labelField", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (s SubFormControl) LabelRemark(value string) SubFormControl {
	return s.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (s SubFormControl) LabelWidth(value string) SubFormControl {
	return s.set("labelWidth", value)
}

// MaxLength 最多个数
func (s SubFormControl) MaxLength(value string) SubFormControl {
	return s.set("maxLength", value)
}

// MinLength 最少个数
func (s SubFormControl) MinLength(value string) SubFormControl {
	return s.set("minLength", value)
}

// Mode 配置当前表单项展示模式
func (s SubFormControl) Mode(value string) SubFormControl {
	return s.set("mode", value)
}

// Multiple 是否多选
func (s SubFormControl) Multiple(value bool) SubFormControl {
	return s.set("multiple", value)
}

// Name 字段名，表单提交时的 key
func (s SubFormControl) Name(value string) SubFormControl {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s SubFormControl) OnEvent(value string) SubFormControl {
	return s.set("onEvent", value)
}

// Placeholder 占位符
func (s SubFormControl) Placeholder(value string) SubFormControl {
	return s.set("placeholder", value)
}

// ReadOnly 是否只读
func (s SubFormControl) ReadOnly(value bool) SubFormControl {
	return s.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (s SubFormControl) ReadOnlyOn(value string) SubFormControl {
	return s.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (s SubFormControl) Remark(value string) SubFormControl {
	return s.set("remark", value)
}

// Removable 是否可删除
func (s SubFormControl) Removable(value bool) SubFormControl {
	return s.set("removable", value)
}

// Required 是否为必填
func (s SubFormControl) Required(value bool) SubFormControl {
	return s.set("required", value)
}

// Row 行
func (s SubFormControl) Row(value string) SubFormControl {
	return s.set("row", value)
}

// SaveImmediately 是否立即保存
func (s SubFormControl) SaveImmediately(value bool) SubFormControl {
	return s.set("saveImmediately", value)
}

// Scaffold
func (s SubFormControl) Scaffold(value string) SubFormControl {
	return s.set("scaffold", value)
}

// ShowErrorMsg 是否在左下角显示报错信息
func (s SubFormControl) ShowErrorMsg(value bool) SubFormControl {
	return s.set("showErrorMsg", value)
}

// Size 表单项大小
func (s SubFormControl) Size(value string) SubFormControl {
	return s.set("size", value)
}

// Static 是否静态展示
func (s SubFormControl) Static(value bool) SubFormControl {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (s SubFormControl) StaticClassName(value string) SubFormControl {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (s SubFormControl) StaticInputClassName(value string) SubFormControl {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (s SubFormControl) StaticLabelClassName(value string) SubFormControl {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (s SubFormControl) StaticOn(value string) SubFormControl {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s SubFormControl) StaticPlaceholder(value string) SubFormControl {
	return s.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式
func (s SubFormControl) StaticSchema(value string) SubFormControl {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s SubFormControl) Style(value string) SubFormControl {
	return s.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (s SubFormControl) SubmitOnChange(value bool) SubFormControl {
	return s.set("submitOnChange", value)
}

// TestIdBuilder 测试ID构建器
func (s SubFormControl) TestIdBuilder(value string) SubFormControl {
	return s.set("testIdBuilder", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (s SubFormControl) UseMobileUI(value bool) SubFormControl {
	return s.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (s SubFormControl) ValidateApi(value string) SubFormControl {
	return s.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (s SubFormControl) ValidateOnChange(value bool) SubFormControl {
	return s.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (s SubFormControl) ValidationErrors(value string) SubFormControl {
	return s.set("validationErrors", value)
}

// Validations 验证规则
func (s SubFormControl) Validations(value string) SubFormControl {
	return s.set("validations", value)
}

// Value 默认值
func (s SubFormControl) Value(value string) SubFormControl {
	return s.set("value", value)
}

// Visible 是否显示
func (s SubFormControl) Visible(value bool) SubFormControl {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式
func (s SubFormControl) VisibleOn(value string) SubFormControl {
	return s.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (s SubFormControl) Width(value string) SubFormControl {
	return s.set("width", value)
}
