package comp

// InputCityControl 城市选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/city
type InputCityControl Schema

// NewInputCityControl 创建一个新的 InputCityControl 实例，并设置默认的 type
func NewInputCityControl() InputCityControl {
	return make(InputCityControl).set("type", "input-city")
}

func (i InputCityControl) set(key string, value interface{}) InputCityControl {
	i[key] = value
	return i
}

// AllowCity 允许选择城市？
func (i InputCityControl) AllowCity(value bool) InputCityControl {
	return i.set("allowCity", value)
}

// AllowDistrict 允许选择地区？
func (i InputCityControl) AllowDistrict(value bool) InputCityControl {
	return i.set("allowDistrict", value)
}

// AllowStreet 允许选择街道？
func (i InputCityControl) AllowStreet(value bool) InputCityControl {
	return i.set("allowStreet", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (i InputCityControl) AutoFill(value string) InputCityControl {
	return i.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputCityControl) ClassName(value string) InputCityControl {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (i InputCityControl) ClearValueOnHidden(value bool) InputCityControl {
	return i.set("clearValueOnHidden", value)
}

// Delimiter 拼接的符号是啥？
func (i InputCityControl) Delimiter(value string) InputCityControl {
	return i.set("delimiter", value)
}

// Desc
func (i InputCityControl) Desc(value string) InputCityControl {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (i InputCityControl) Description(value string) InputCityControl {
	return i.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (i InputCityControl) DescriptionClassName(value string) InputCityControl {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i InputCityControl) Disabled(value bool) InputCityControl {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputCityControl) DisabledOn(value string) InputCityControl {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (i InputCityControl) EditorSetting(value string) InputCityControl {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (i InputCityControl) ExtraName(value string) InputCityControl {
	return i.set("extraName", value)
}

// ExtractValue 开启后只会存城市的 code 信息
func (i InputCityControl) ExtractValue(value bool) InputCityControl {
	return i.set("extractValue", value)
}

// Hidden 是否隐藏
func (i InputCityControl) Hidden(value bool) InputCityControl {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputCityControl) HiddenOn(value string) InputCityControl {
	return i.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (i InputCityControl) Hint(value string) InputCityControl {
	return i.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (i InputCityControl) Horizontal(value string) InputCityControl {
	return i.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (i InputCityControl) Id(value string) InputCityControl {
	return i.set("id", value)
}

// InitAutoFill
func (i InputCityControl) InitAutoFill(value string) InputCityControl {
	return i.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (i InputCityControl) Inline(value bool) InputCityControl {
	return i.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (i InputCityControl) InputClassName(value string) InputCityControl {
	return i.set("inputClassName", value)
}

// ItemClassName 下拉框className
func (i InputCityControl) ItemClassName(value string) InputCityControl {
	return i.set("itemClassName", value)
}

// JoinValues 是否将各个信息拼接成字符串。
func (i InputCityControl) JoinValues(value bool) InputCityControl {
	return i.set("joinValues", value)
}

// Label 描述标题
func (i InputCityControl) Label(value string) InputCityControl {
	return i.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (i InputCityControl) LabelAlign(value string) InputCityControl {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i InputCityControl) LabelClassName(value string) InputCityControl {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (i InputCityControl) LabelRemark(value string) InputCityControl {
	return i.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (i InputCityControl) LabelWidth(value string) InputCityControl {
	return i.set("labelWidth", value)
}

// LoadingConfig
func (i InputCityControl) LoadingConfig(value string) InputCityControl {
	return i.set("loadingConfig", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (i InputCityControl) Mode(value string) InputCityControl {
	return i.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (i InputCityControl) Name(value string) InputCityControl {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i InputCityControl) OnEvent(value string) InputCityControl {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i InputCityControl) Placeholder(value string) InputCityControl {
	return i.set("placeholder", value)
}

// ReadOnly 是否只读
func (i InputCityControl) ReadOnly(value bool) InputCityControl {
	return i.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (i InputCityControl) ReadOnlyOn(value string) InputCityControl {
	return i.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (i InputCityControl) Remark(value string) InputCityControl {
	return i.set("remark", value)
}

// Required 是否为必填
func (i InputCityControl) Required(value bool) InputCityControl {
	return i.set("required", value)
}

// Row
func (i InputCityControl) Row(value string) InputCityControl {
	return i.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (i InputCityControl) SaveImmediately(value bool) InputCityControl {
	return i.set("saveImmediately", value)
}

// Searchable 是否显示搜索框
func (i InputCityControl) Searchable(value bool) InputCityControl {
	return i.set("searchable", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (i InputCityControl) Size(value string) InputCityControl {
	return i.set("size", value)
}

// Static 是否静态展示
func (i InputCityControl) Static(value bool) InputCityControl {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputCityControl) StaticClassName(value string) InputCityControl {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputCityControl) StaticInputClassName(value string) InputCityControl {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputCityControl) StaticLabelClassName(value string) InputCityControl {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputCityControl) StaticOn(value string) InputCityControl {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i InputCityControl) StaticPlaceholder(value string) InputCityControl {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i InputCityControl) StaticSchema(value string) InputCityControl {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i InputCityControl) Style(value string) InputCityControl {
	return i.set("style", value)
}

// SubmitOnChange 当值变化后是否自动提交
func (i InputCityControl) SubmitOnChange(value bool) InputCityControl {
	return i.set("submitOnChange", value)
}

// TestIdBuilder
func (i InputCityControl) TestIdBuilder(value string) InputCityControl {
	return i.set("testIdBuilder", value)
}

// UseMobile 是否使用移动端适配
func (i InputCityControl) UseMobile(value bool) InputCityControl {
	return i.set("useMobile", value)
}

// ValidateApi 远端校验表单项接口
func (i InputCityControl) ValidateApi(value string) InputCityControl {
	return i.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (i InputCityControl) ValidateOnChange(value bool) InputCityControl {
	return i.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (i InputCityControl) ValidationErrors(value string) InputCityControl {
	return i.set("validationErrors", value)
}

// Validations
func (i InputCityControl) Validations(value string) InputCityControl {
	return i.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (i InputCityControl) Value(value interface{}) InputCityControl {
	return i.set("value", value)
}

// Visible 是否显示
func (i InputCityControl) Visible(value bool) InputCityControl {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputCityControl) VisibleOn(value string) InputCityControl {
	return i.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (i InputCityControl) Width(value string) InputCityControl {
	return i.set("width", value)
}
