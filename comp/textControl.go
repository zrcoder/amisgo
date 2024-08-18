package comp

// InputText 文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/text
//
// @version 6.7.0
type InputText Schema

// NewInputText 创建一个新的 TextControl 实例
func NewInputText() InputText {
	return InputText{}.set("type", "input-text")
}

// Type 默认 input-text，可设置为 divider、input-url、input-email、input-password
func (t InputText) Type(value string) InputText {
	return t.set("type", value)
}

// AddApi 添加时调用的接口
func (t InputText) AddApi(value string) InputText {
	return t.set("addApi", value)
}

// AddControls 新增时的表单项。
func (t InputText) AddControls(value string) InputText {
	return t.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (t InputText) AddDialog(value string) InputText {
	return t.set("addDialog", value)
}

// AddOn
func (t InputText) AddOn(value string) InputText {
	return t.set("addOn", value)
}

// AutoComplete 自动完成 API，当输入部分文字的时候，会将这些文字通过 ${term} 可以取到，发送给接口。 接口可以返回匹配到的选项，帮助用户输入。
func (t InputText) AutoComplete(value string) InputText {
	return t.set("autoComplete", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (t InputText) AutoFill(value string) InputText {
	return t.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (t InputText) BorderMode(value string) InputText {
	return t.set("borderMode", value)
}

// ClassName 容器 css 类名
func (t InputText) ClassName(value string) InputText {
	return t.set("className", value)
}

// ClearValueOnEmpty 在内容为空的时候清除值
func (t InputText) ClearValueOnEmpty(value bool) InputText {
	return t.set("clearValueOnEmpty", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (t InputText) ClearValueOnHidden(value bool) InputText {
	return t.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (t InputText) Clearable(value bool) InputText {
	return t.set("clearable", value)
}

// Creatable 是否可以新增
func (t InputText) Creatable(value bool) InputText {
	return t.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (t InputText) CreateBtnLabel(value string) InputText {
	return t.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (t InputText) DeferApi(value string) InputText {
	return t.set("deferApi", value)
}

// DeferField 懒加载字段
func (t InputText) DeferField(value string) InputText {
	return t.set("deferField", value)
}

// DeleteApi 选项删除 API
func (t InputText) DeleteApi(value string) InputText {
	return t.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (t InputText) DeleteConfirmText(value string) InputText {
	return t.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (t InputText) Delimiter(value string) InputText {
	return t.set("delimiter", value)
}

// Desc
func (t InputText) Desc(value string) InputText {
	return t.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (t InputText) Description(value string) InputText {
	return t.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (t InputText) DescriptionClassName(value string) InputText {
	return t.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (t InputText) Disabled(value bool) InputText {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t InputText) DisabledOn(value string) InputText {
	return t.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (t InputText) EditApi(value string) InputText {
	return t.set("editApi", value)
}

// EditControls 选项修改的表单项
func (t InputText) EditControls(value string) InputText {
	return t.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (t InputText) EditDialog(value string) InputText {
	return t.set("editDialog", value)
}

// Editable 是否可以编辑
func (t InputText) Editable(value bool) InputText {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t InputText) EditorSetting(value string) InputText {
	return t.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (t InputText) ExtraName(value string) InputText {
	return t.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (t InputText) ExtractValue(value bool) InputText {
	return t.set("extractValue", value)
}

// Hidden 是否隐藏
func (t InputText) Hidden(value bool) InputText {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t InputText) HiddenOn(value string) InputText {
	return t.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (t InputText) Hint(value string) InputText {
	return t.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。
func (t InputText) Horizontal(value string) InputText {
	return t.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t InputText) Id(value string) InputText {
	return t.set("id", value)
}

// InitAutoFill
func (t InputText) InitAutoFill(value string) InputText {
	return t.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (t InputText) InitFetch(value bool) InputText {
	return t.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (t InputText) InitFetchOn(value string) InputText {
	return t.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式。
func (t InputText) Inline(value bool) InputText {
	return t.set("inline", value)
}

// InputClassName 配置 input className
func (t InputText) InputClassName(value string) InputText {
	return t.set("inputClassName", value)
}

// InputControlClassName control节点的CSS类名
func (t InputText) InputControlClassName(value string) InputText {
	return t.set("inputControlClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (t InputText) JoinValues(value bool) InputText {
	return t.set("joinValues", value)
}

// Label 描述标题
func (t InputText) Label(value string) InputText {
	return t.set("label", value)
}

// LabelAlign 描述标题 可选值: right | left | top | inherit
func (t InputText) LabelAlign(value string) InputText {
	return t.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (t InputText) LabelClassName(value string) InputText {
	return t.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (t InputText) LabelRemark(value string) InputText {
	return t.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (t InputText) LabelWidth(value string) InputText {
	return t.set("labelWidth", value)
}

// MaxLength 限制文字最大输入个数
func (t InputText) MaxLength(value string) InputText {
	return t.set("maxLength", value)
}

// MinLength 限制文字最小输入个数
func (t InputText) MinLength(value string) InputText {
	return t.set("minLength", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (t InputText) Mode(value string) InputText {
	return t.set("mode", value)
}

// Multiple 是否为多选模式
func (t InputText) Multiple(value bool) InputText {
	return t.set("multiple", value)
}

// Name 表单项名称
func (t InputText) Name(value string) InputText {
	return t.set("name", value)
}

// Note
func (t InputText) Note(value string) InputText {
	return t.set("note", value)
}

// OnChange 触发的事件
func (t InputText) OnChange(value string) InputText {
	return t.set("onChange", value)
}

// OnFocus 触发的事件
func (t InputText) OnFocus(value string) InputText {
	return t.set("onFocus", value)
}

// OnBlur 触发的事件
func (t InputText) OnBlur(value string) InputText {
	return t.set("onBlur", value)
}

// OnInit 触发的事件
func (t InputText) OnInit(value string) InputText {
	return t.set("onInit", value)
}

// OnSearch 触发的事件
func (t InputText) OnSearch(value string) InputText {
	return t.set("onSearch", value)
}

// Option
func (t InputText) Option(value string) InputText {
	return t.set("option", value)
}

// Options 选项列表
func (t InputText) Options(value string) InputText {
	return t.set("options", value)
}

// OptionLabel 在多选模式下，选项的 label 位置。可选值: top | left | right
func (t InputText) OptionLabel(value string) InputText {
	return t.set("optionLabel", value)
}

// Placeholder 提示文字
func (t InputText) Placeholder(value string) InputText {
	return t.set("placeholder", value)
}

// ReadOnly 是否为只读模式
func (t InputText) ReadOnly(value bool) InputText {
	return t.set("readOnly", value)
}

// RefixApi 配置选择项的接口
func (t InputText) RefixApi(value string) InputText {
	return t.set("refixApi", value)
}

// ResetValue
func (t InputText) ResetValue(value string) InputText {
	return t.set("resetValue", value)
}

// Rules 校验规则
func (t InputText) Rules(value string) InputText {
	return t.set("rules", value)
}

// SearchApi 搜索接口
func (t InputText) SearchApi(value string) InputText {
	return t.set("searchApi", value)
}

// Source 数据来源
func (t InputText) Source(value string) InputText {
	return t.set("source", value)
}

// StaticClassName 用于静态展示的类名
func (t InputText) StaticClassName(value string) InputText {
	return t.set("staticClassName", value)
}

// StaticLabel 静态显示的标题
func (t InputText) StaticLabel(value string) InputText {
	return t.set("staticLabel", value)
}

// StaticOn 静态显示的条件
func (t InputText) StaticOn(value string) InputText {
	return t.set("staticOn", value)
}

// StaticValue 静态值
func (t InputText) StaticValue(value string) InputText {
	return t.set("staticValue", value)
}

// Tips 提示文本
func (t InputText) Tips(value string) InputText {
	return t.set("tips", value)
}

// ValidationErrors 校验不通过时的提示信息
func (t InputText) ValidationErrors(value string) InputText {
	return t.set("validationErrors", value)
}

// Value 表单项的值
func (t InputText) Value(value interface{}) InputText {
	return t.set("value", value)
}

// Visible 控制组件是否可见
func (t InputText) Visible(value bool) InputText {
	return t.set("visible", value)
}

// VisibleOn 是否可见表达式 (表达式，语法 `data.xxx > 5`。)
func (t InputText) VisibleOn(value string) InputText {
	return t.set("visibleOn", value)
}

// Width 控件宽度
func (t InputText) Width(value string) InputText {
	return t.set("width", value)
}

// OnInput 触发的事件
func (t InputText) OnInput(value string) InputText {
	return t.set("onInput", value)
}

// ValueType
func (t InputText) ValueType(value string) InputText {
	return t.set("valueType", value)
}

// ValueOnEvent
func (t InputText) ValueOnEvent(value string) InputText {
	return t.set("valueOnEvent", value)
}

// set 设置属性
func (t InputText) set(key string, value interface{}) InputText {
	t[key] = value
	return t
}
