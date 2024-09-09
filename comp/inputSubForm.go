package comp

// inputSubForm 子表单

type inputSubForm schema

func InputSubForm() inputSubForm {
	return inputSubForm{}.set("type", "input-sub-form")
}

func (s inputSubForm) set(key string, value any) inputSubForm {
	s[key] = value
	return s
}

// AddButtonClassName 新增按钮 CSS 类名 (css类名，配置字符串，或者对象。)
func (s inputSubForm) AddButtonClassName(value string) inputSubForm {
	return s.set("addButtonClassName", value)
}

// AddButtonText 新增按钮文字
func (s inputSubForm) AddButtonText(value string) inputSubForm {
	return s.set("addButtonText", value)
}

// Addable 是否可新增
func (s inputSubForm) Addable(value bool) inputSubForm {
	return s.set("addable", value)
}

// AutoFill 自动填充
func (s inputSubForm) AutoFill(value string) inputSubForm {
	return s.set("autoFill", value)
}

// BtnLabel 按钮默认名称
func (s inputSubForm) BtnLabel(value string) inputSubForm {
	return s.set("btnLabel", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。)
func (s inputSubForm) ClassName(value string) inputSubForm {
	return s.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (s inputSubForm) ClearValueOnHidden(value bool) inputSubForm {
	return s.set("clearValueOnHidden", value)
}

// Desc 描述内容
func (s inputSubForm) Desc(value string) inputSubForm {
	return s.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (s inputSubForm) Description(value string) inputSubForm {
	return s.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (s inputSubForm) DescriptionClassName(value string) inputSubForm {
	return s.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (s inputSubForm) Disabled(value bool) inputSubForm {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s inputSubForm) DisabledOn(value string) inputSubForm {
	return s.set("disabledOn", value)
}

// Draggable 是否可拖拽排序
func (s inputSubForm) Draggable(value bool) inputSubForm {
	return s.set("draggable", value)
}

// DraggableTip 拖拽提示信息
func (s inputSubForm) DraggableTip(value string) inputSubForm {
	return s.set("draggableTip", value)
}

// EditorSetting 编辑器配置
func (s inputSubForm) EditorSetting(value string) inputSubForm {
	return s.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (s inputSubForm) ExtraName(value string) inputSubForm {
	return s.set("extraName", value)
}

// Form 子表单详情
func (s inputSubForm) Form(value string) inputSubForm {
	return s.set("form", value)
}

// Hidden 是否隐藏
func (s inputSubForm) Hidden(value bool) inputSubForm {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s inputSubForm) HiddenOn(value string) inputSubForm {
	return s.set("hiddenOn", value)
}

// Hint 输入提示
func (s inputSubForm) Hint(value string) inputSubForm {
	return s.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (s inputSubForm) Horizontal(value string) inputSubForm {
	return s.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s inputSubForm) Id(value string) inputSubForm {
	return s.set("id", value)
}

// InitAutoFill 初始化自动填充
func (s inputSubForm) InitAutoFill(value string) inputSubForm {
	return s.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (s inputSubForm) Inline(value bool) inputSubForm {
	return s.set("inline", value)
}

// InputClassName 配置 input className
func (s inputSubForm) InputClassName(value string) inputSubForm {
	return s.set("inputClassName", value)
}

// ItemClassName 值元素的类名
func (s inputSubForm) ItemClassName(value string) inputSubForm {
	return s.set("itemClassName", value)
}

// ItemsClassName 值列表元素的类名
func (s inputSubForm) ItemsClassName(value string) inputSubForm {
	return s.set("itemsClassName", value)
}

// Label 描述标题
func (s inputSubForm) Label(value string) inputSubForm {
	return s.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (s inputSubForm) LabelAlign(value string) inputSubForm {
	return s.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (s inputSubForm) LabelClassName(value string) inputSubForm {
	return s.set("labelClassName", value)
}

// LabelField 当值中存在这个字段，则按钮名称将使用此字段的值来展示
func (s inputSubForm) LabelField(value string) inputSubForm {
	return s.set("labelField", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (s inputSubForm) LabelRemark(value string) inputSubForm {
	return s.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (s inputSubForm) LabelWidth(value string) inputSubForm {
	return s.set("labelWidth", value)
}

// MaxLength 最多个数
func (s inputSubForm) MaxLength(value string) inputSubForm {
	return s.set("maxLength", value)
}

// MinLength 最少个数
func (s inputSubForm) MinLength(value string) inputSubForm {
	return s.set("minLength", value)
}

// Mode 配置当前表单项展示模式
func (s inputSubForm) Mode(value string) inputSubForm {
	return s.set("mode", value)
}

// Multiple 是否多选
func (s inputSubForm) Multiple(value bool) inputSubForm {
	return s.set("multiple", value)
}

// Name 字段名，表单提交时的 key
func (s inputSubForm) Name(value string) inputSubForm {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s inputSubForm) OnEvent(value any) inputSubForm {
	return s.set("onEvent", value)
}

// Placeholder 占位符
func (s inputSubForm) Placeholder(value string) inputSubForm {
	return s.set("placeholder", value)
}

// ReadOnly 是否只读
func (s inputSubForm) ReadOnly(value bool) inputSubForm {
	return s.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (s inputSubForm) ReadOnlyOn(value string) inputSubForm {
	return s.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (s inputSubForm) Remark(value string) inputSubForm {
	return s.set("remark", value)
}

// Removable 是否可删除
func (s inputSubForm) Removable(value bool) inputSubForm {
	return s.set("removable", value)
}

// Required 是否为必填
func (s inputSubForm) Required(value bool) inputSubForm {
	return s.set("required", value)
}

// Row 行
func (s inputSubForm) Row(value string) inputSubForm {
	return s.set("row", value)
}

// SaveImmediately 是否立即保存
func (s inputSubForm) SaveImmediately(value bool) inputSubForm {
	return s.set("saveImmediately", value)
}

// Scaffold
func (s inputSubForm) Scaffold(value string) inputSubForm {
	return s.set("scaffold", value)
}

// ShowErrorMsg 是否在左下角显示报错信息
func (s inputSubForm) ShowErrorMsg(value bool) inputSubForm {
	return s.set("showErrorMsg", value)
}

// Size 表单项大小
func (s inputSubForm) Size(value string) inputSubForm {
	return s.set("size", value)
}

// Static 是否静态展示
func (s inputSubForm) Static(value bool) inputSubForm {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (s inputSubForm) StaticClassName(value string) inputSubForm {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (s inputSubForm) StaticInputClassName(value string) inputSubForm {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (s inputSubForm) StaticLabelClassName(value string) inputSubForm {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (s inputSubForm) StaticOn(value string) inputSubForm {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s inputSubForm) StaticPlaceholder(value string) inputSubForm {
	return s.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式
func (s inputSubForm) StaticSchema(value string) inputSubForm {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s inputSubForm) Style(value any) inputSubForm {
	return s.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (s inputSubForm) SubmitOnChange(value bool) inputSubForm {
	return s.set("submitOnChange", value)
}

// TestIdBuilder 测试ID构建器
func (s inputSubForm) TestIdBuilder(value string) inputSubForm {
	return s.set("testIdBuilder", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (s inputSubForm) UseMobileUI(value bool) inputSubForm {
	return s.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (s inputSubForm) ValidateApi(value string) inputSubForm {
	return s.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (s inputSubForm) ValidateOnChange(value bool) inputSubForm {
	return s.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (s inputSubForm) ValidationErrors(value string) inputSubForm {
	return s.set("validationErrors", value)
}

// Validations 验证规则
func (s inputSubForm) Validations(value string) inputSubForm {
	return s.set("validations", value)
}

// Value 默认值
func (s inputSubForm) Value(value string) inputSubForm {
	return s.set("value", value)
}

// Visible 是否显示
func (s inputSubForm) Visible(value bool) inputSubForm {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式
func (s inputSubForm) VisibleOn(value string) inputSubForm {
	return s.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (s inputSubForm) Width(value string) inputSubForm {
	return s.set("width", value)
}
