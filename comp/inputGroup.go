package comp

// InputGroupControl 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-group
type InputGroupControl Schema

// NewInputGroupControl 创建一个新的 InputGroupControl 实例，并设置默认的 type
func NewInputGroupControl() InputGroupControl {
	return make(InputGroupControl).set("type", "input-group")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i InputGroupControl) set(key string, value interface{}) InputGroupControl {
	i[key] = value
	return i
}

// AutoFill 设置自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (i InputGroupControl) AutoFill(value string) InputGroupControl {
	return i.set("autoFill", value)
}

// Body 设置 FormItem 集合
func (i InputGroupControl) Body(value ...interface{}) InputGroupControl {
	return i.set("body", value)
}

// ClassName 设置容器 css 类名
func (i InputGroupControl) ClassName(value string) InputGroupControl {
	return i.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (i InputGroupControl) ClearValueOnHidden(value bool) InputGroupControl {
	return i.set("clearValueOnHidden", value)
}

// Desc 设置描述内容
func (i InputGroupControl) Desc(value string) InputGroupControl {
	return i.set("desc", value)
}

// Description 设置描述内容，支持 Html 片段
func (i InputGroupControl) Description(value string) InputGroupControl {
	return i.set("description", value)
}

// DescriptionClassName 设置配置描述上的 className
func (i InputGroupControl) DescriptionClassName(value string) InputGroupControl {
	return i.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (i InputGroupControl) Disabled(value bool) InputGroupControl {
	return i.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (i InputGroupControl) DisabledOn(value string) InputGroupControl {
	return i.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置，运行时可以忽略
func (i InputGroupControl) EditorSetting(value string) InputGroupControl {
	return i.set("editorSetting", value)
}

// ExtraName 设置额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (i InputGroupControl) ExtraName(value string) InputGroupControl {
	return i.set("extraName", value)
}

// Hidden 设置是否隐藏
func (i InputGroupControl) Hidden(value bool) InputGroupControl {
	return i.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (i InputGroupControl) HiddenOn(value string) InputGroupControl {
	return i.set("hiddenOn", value)
}

// Hint 设置输入提示，聚焦的时候显示
func (i InputGroupControl) Hint(value string) InputGroupControl {
	return i.set("hint", value)
}

// Horizontal 设置当配置为水平布局的时候，用来配置具体的左右分配
func (i InputGroupControl) Horizontal(value string) InputGroupControl {
	return i.set("horizontal", value)
}

// ID 设置组件唯一 id，主要用于日志采集
func (i InputGroupControl) ID(value string) InputGroupControl {
	return i.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (i InputGroupControl) InitAutoFill(value string) InputGroupControl {
	return i.set("initAutoFill", value)
}

// Inline 设置表单 control 是否为 inline 模式
func (i InputGroupControl) Inline(value bool) InputGroupControl {
	return i.set("inline", value)
}

// InputClassName 设置配置 input className
func (i InputGroupControl) InputClassName(value string) InputGroupControl {
	return i.set("inputClassName", value)
}

// Label 设置描述标题
func (i InputGroupControl) Label(value string) InputGroupControl {
	return i.set("label", value)
}

// LabelAlign 设置描述标题对齐方式
func (i InputGroupControl) LabelAlign(value string) InputGroupControl {
	return i.set("labelAlign", value)
}

// LabelClassName 设置配置 label className
func (i InputGroupControl) LabelClassName(value string) InputGroupControl {
	return i.set("labelClassName", value)
}

// LabelRemark 设置显示一个小图标, 鼠标放上去的时候显示提示内容
func (i InputGroupControl) LabelRemark(value string) InputGroupControl {
	return i.set("labelRemark", value)
}

// LabelWidth 设置 label 自定义宽度，默认单位为 px
func (i InputGroupControl) LabelWidth(value string) InputGroupControl {
	return i.set("labelWidth", value)
}

// Mode 设置配置当前表单项展示模式
func (i InputGroupControl) Mode(value string) InputGroupControl {
	return i.set("mode", value)
}

// Name 设置字段名，表单提交时的 key，支持多层级
func (i InputGroupControl) Name(value string) InputGroupControl {
	return i.set("name", value)
}

// OnEvent 设置事件动作配置
func (i InputGroupControl) OnEvent(value string) InputGroupControl {
	return i.set("onEvent", value)
}

// Placeholder 设置占位符
func (i InputGroupControl) Placeholder(value string) InputGroupControl {
	return i.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (i InputGroupControl) ReadOnly(value bool) InputGroupControl {
	return i.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (i InputGroupControl) ReadOnlyOn(value string) InputGroupControl {
	return i.set("readOnlyOn", value)
}

// Remark 设置显示一个小图标, 鼠标放上去的时候显示提示内容
func (i InputGroupControl) Remark(value string) InputGroupControl {
	return i.set("remark", value)
}

// Required 设置是否为必填
func (i InputGroupControl) Required(value bool) InputGroupControl {
	return i.set("required", value)
}

// Row 设置配置行
func (i InputGroupControl) Row(value string) InputGroupControl {
	return i.set("row", value)
}

// SaveImmediately 设置是否立即保存 (TableColumn中使用)
func (i InputGroupControl) SaveImmediately(value bool) InputGroupControl {
	return i.set("saveImmediately", value)
}

// Size 设置表单项大小
func (i InputGroupControl) Size(value string) InputGroupControl {
	return i.set("size", value)
}

// Static 设置是否静态展示
func (i InputGroupControl) Static(value bool) InputGroupControl {
	return i.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (i InputGroupControl) StaticClassName(value string) InputGroupControl {
	return i.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (i InputGroupControl) StaticInputClassName(value string) InputGroupControl {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (i InputGroupControl) StaticLabelClassName(value string) InputGroupControl {
	return i.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (i InputGroupControl) StaticOn(value string) InputGroupControl {
	return i.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (i InputGroupControl) StaticPlaceholder(value string) InputGroupControl {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (i InputGroupControl) StaticSchema(value string) InputGroupControl {
	return i.set("staticSchema", value)
}

// Style 设置组件样式
func (i InputGroupControl) Style(value string) InputGroupControl {
	return i.set("style", value)
}

// SubmitOnChange 设置当修改完的时候是否提交表单
func (i InputGroupControl) SubmitOnChange(value bool) InputGroupControl {
	return i.set("submitOnChange", value)
}

// TestIdBuilder 设置测试 id 构造器
func (i InputGroupControl) TestIdBuilder(value string) InputGroupControl {
	return i.set("testIdBuilder", value)
}

// UseMobileUI 设置可以组件级别用来关闭移动端样式
func (i InputGroupControl) UseMobileUI(value bool) InputGroupControl {
	return i.set("useMobileUI", value)
}

// ValidateApi 设置远端校验表单项接口
func (i InputGroupControl) ValidateApi(value string) InputGroupControl {
	return i.set("validateApi", value)
}

// ValidateOnChange 设置不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (i InputGroupControl) ValidateOnChange(value bool) InputGroupControl {
	return i.set("validateOnChange", value)
}

// ValidationConfig 设置校验提示信息配置
func (i InputGroupControl) ValidationConfig(value string) InputGroupControl {
	return i.set("validationConfig", value)
}

// ValidationErrors 设置验证失败的提示信息
func (i InputGroupControl) ValidationErrors(value string) InputGroupControl {
	return i.set("validationErrors", value)
}

// Validations 设置校验规则
func (i InputGroupControl) Validations(value string) InputGroupControl {
	return i.set("validations", value)
}

// Value 设置默认值，切记只能是静态值
func (i InputGroupControl) Value(value string) InputGroupControl {
	return i.set("value", value)
}

// Visible 设置是否显示
func (i InputGroupControl) Visible(value bool) InputGroupControl {
	return i.set("visible", value)
}

// VisibleOn 设置是否显示表达式
func (i InputGroupControl) VisibleOn(value string) InputGroupControl {
	return i.set("visibleOn", value)
}

// Width 设置在 Table 中调整宽度
func (i InputGroupControl) Width(value string) InputGroupControl {
	return i.set("width", value)
}
