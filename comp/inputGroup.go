package comp

// inputGroup 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-group
type inputGroup schema

func InputGroup() inputGroup {
	return make(inputGroup).set("type", "input-group")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i inputGroup) set(key string, value any) inputGroup {
	i[key] = value
	return i
}

// AutoFill 设置自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (i inputGroup) AutoFill(value string) inputGroup {
	return i.set("autoFill", value)
}

// Body 设置 FormItem 集合
func (i inputGroup) Body(value ...any) inputGroup {
	return i.set("body", value)
}

// ClassName 设置容器 css 类名
func (i inputGroup) ClassName(value string) inputGroup {
	return i.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (i inputGroup) ClearValueOnHidden(value bool) inputGroup {
	return i.set("clearValueOnHidden", value)
}

// Desc 设置描述内容
func (i inputGroup) Desc(value string) inputGroup {
	return i.set("desc", value)
}

// Description 设置描述内容，支持 Html 片段
func (i inputGroup) Description(value string) inputGroup {
	return i.set("description", value)
}

// DescriptionClassName 设置配置描述上的 className
func (i inputGroup) DescriptionClassName(value string) inputGroup {
	return i.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (i inputGroup) Disabled(value bool) inputGroup {
	return i.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (i inputGroup) DisabledOn(value string) inputGroup {
	return i.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置，运行时可以忽略
func (i inputGroup) EditorSetting(value string) inputGroup {
	return i.set("editorSetting", value)
}

// ExtraName 设置额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (i inputGroup) ExtraName(value string) inputGroup {
	return i.set("extraName", value)
}

// Hidden 设置是否隐藏
func (i inputGroup) Hidden(value bool) inputGroup {
	return i.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (i inputGroup) HiddenOn(value string) inputGroup {
	return i.set("hiddenOn", value)
}

// Hint 设置输入提示，聚焦的时候显示
func (i inputGroup) Hint(value string) inputGroup {
	return i.set("hint", value)
}

// Horizontal 设置当配置为水平布局的时候，用来配置具体的左右分配
func (i inputGroup) Horizontal(value string) inputGroup {
	return i.set("horizontal", value)
}

// ID 设置组件唯一 id，主要用于日志采集
func (i inputGroup) ID(value string) inputGroup {
	return i.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (i inputGroup) InitAutoFill(value string) inputGroup {
	return i.set("initAutoFill", value)
}

// Inline 设置表单 control 是否为 inline 模式
func (i inputGroup) Inline(value bool) inputGroup {
	return i.set("inline", value)
}

// InputClassName 设置配置 input className
func (i inputGroup) InputClassName(value string) inputGroup {
	return i.set("inputClassName", value)
}

// Label 设置描述标题
func (i inputGroup) Label(value string) inputGroup {
	return i.set("label", value)
}

// LabelAlign 设置描述标题对齐方式
func (i inputGroup) LabelAlign(value string) inputGroup {
	return i.set("labelAlign", value)
}

// LabelClassName 设置配置 label className
func (i inputGroup) LabelClassName(value string) inputGroup {
	return i.set("labelClassName", value)
}

// LabelRemark 设置显示一个小图标, 鼠标放上去的时候显示提示内容
func (i inputGroup) LabelRemark(value string) inputGroup {
	return i.set("labelRemark", value)
}

// LabelWidth 设置 label 自定义宽度，默认单位为 px
func (i inputGroup) LabelWidth(value string) inputGroup {
	return i.set("labelWidth", value)
}

// Mode 设置配置当前表单项展示模式
func (i inputGroup) Mode(value string) inputGroup {
	return i.set("mode", value)
}

// Name 设置字段名，表单提交时的 key，支持多层级
func (i inputGroup) Name(value string) inputGroup {
	return i.set("name", value)
}

// OnEvent 设置事件动作配置
func (i inputGroup) OnEvent(value any) inputGroup {
	return i.set("onEvent", value)
}

// Placeholder 设置占位符
func (i inputGroup) Placeholder(value string) inputGroup {
	return i.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (i inputGroup) ReadOnly(value bool) inputGroup {
	return i.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (i inputGroup) ReadOnlyOn(value string) inputGroup {
	return i.set("readOnlyOn", value)
}

// Remark 设置显示一个小图标, 鼠标放上去的时候显示提示内容
func (i inputGroup) Remark(value string) inputGroup {
	return i.set("remark", value)
}

// Required 设置是否为必填
func (i inputGroup) Required(value bool) inputGroup {
	return i.set("required", value)
}

// Row 设置配置行
func (i inputGroup) Row(value string) inputGroup {
	return i.set("row", value)
}

// SaveImmediately 设置是否立即保存 (TableColumn中使用)
func (i inputGroup) SaveImmediately(value bool) inputGroup {
	return i.set("saveImmediately", value)
}

// Size 设置表单项大小
func (i inputGroup) Size(value string) inputGroup {
	return i.set("size", value)
}

// Static 设置是否静态展示
func (i inputGroup) Static(value bool) inputGroup {
	return i.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (i inputGroup) StaticClassName(value string) inputGroup {
	return i.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (i inputGroup) StaticInputClassName(value string) inputGroup {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (i inputGroup) StaticLabelClassName(value string) inputGroup {
	return i.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (i inputGroup) StaticOn(value string) inputGroup {
	return i.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (i inputGroup) StaticPlaceholder(value string) inputGroup {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (i inputGroup) StaticSchema(value string) inputGroup {
	return i.set("staticSchema", value)
}

// Style 设置组件样式
func (i inputGroup) Style(value any) inputGroup {
	return i.set("style", value)
}

// SubmitOnChange 设置当修改完的时候是否提交表单
func (i inputGroup) SubmitOnChange(value bool) inputGroup {
	return i.set("submitOnChange", value)
}

// TestIdBuilder 设置测试 id 构造器
func (i inputGroup) TestIdBuilder(value string) inputGroup {
	return i.set("testIdBuilder", value)
}

// UseMobileUI 设置可以组件级别用来关闭移动端样式
func (i inputGroup) UseMobileUI(value bool) inputGroup {
	return i.set("useMobileUI", value)
}

// ValidateApi 设置远端校验表单项接口
func (i inputGroup) ValidateApi(value string) inputGroup {
	return i.set("validateApi", value)
}

// ValidateOnChange 设置不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (i inputGroup) ValidateOnChange(value bool) inputGroup {
	return i.set("validateOnChange", value)
}

// ValidationConfig 设置校验提示信息配置
func (i inputGroup) ValidationConfig(value string) inputGroup {
	return i.set("validationConfig", value)
}

// ValidationErrors 设置验证失败的提示信息
func (i inputGroup) ValidationErrors(value string) inputGroup {
	return i.set("validationErrors", value)
}

// Validations 设置校验规则
func (i inputGroup) Validations(value string) inputGroup {
	return i.set("validations", value)
}

// Value 设置默认值，切记只能是静态值
func (i inputGroup) Value(value string) inputGroup {
	return i.set("value", value)
}

// Visible 设置是否显示
func (i inputGroup) Visible(value bool) inputGroup {
	return i.set("visible", value)
}

// VisibleOn 设置是否显示表达式
func (i inputGroup) VisibleOn(value string) inputGroup {
	return i.set("visibleOn", value)
}

// Width 设置在 Table 中调整宽度
func (i inputGroup) Width(value string) inputGroup {
	return i.set("width", value)
}
