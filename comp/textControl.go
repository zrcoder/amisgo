package comp

// TextControl 文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/text
//
// @version 6.7.0
type TextControl Schema

// NewTextControl 创建一个新的 TextControl 实例
func NewTextControl() TextControl {
	return TextControl{}.set("type", "input-text")
}

// AddApi 添加时调用的接口
func (t TextControl) AddApi(value string) TextControl {
	return t.set("addApi", value)
}

// AddControls 新增时的表单项。
func (t TextControl) AddControls(value string) TextControl {
	return t.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (t TextControl) AddDialog(value string) TextControl {
	return t.set("addDialog", value)
}

// AddOn
func (t TextControl) AddOn(value string) TextControl {
	return t.set("addOn", value)
}

// AutoComplete 自动完成 API，当输入部分文字的时候，会将这些文字通过 ${term} 可以取到，发送给接口。 接口可以返回匹配到的选项，帮助用户输入。
func (t TextControl) AutoComplete(value string) TextControl {
	return t.set("autoComplete", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (t TextControl) AutoFill(value string) TextControl {
	return t.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (t TextControl) BorderMode(value string) TextControl {
	return t.set("borderMode", value)
}

// ClassName 容器 css 类名
func (t TextControl) ClassName(value string) TextControl {
	return t.set("className", value)
}

// ClearValueOnEmpty 在内容为空的时候清除值
func (t TextControl) ClearValueOnEmpty(value bool) TextControl {
	return t.set("clearValueOnEmpty", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (t TextControl) ClearValueOnHidden(value bool) TextControl {
	return t.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (t TextControl) Clearable(value bool) TextControl {
	return t.set("clearable", value)
}

// Creatable 是否可以新增
func (t TextControl) Creatable(value bool) TextControl {
	return t.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (t TextControl) CreateBtnLabel(value string) TextControl {
	return t.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (t TextControl) DeferApi(value string) TextControl {
	return t.set("deferApi", value)
}

// DeferField 懒加载字段
func (t TextControl) DeferField(value string) TextControl {
	return t.set("deferField", value)
}

// DeleteApi 选项删除 API
func (t TextControl) DeleteApi(value string) TextControl {
	return t.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (t TextControl) DeleteConfirmText(value string) TextControl {
	return t.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (t TextControl) Delimiter(value string) TextControl {
	return t.set("delimiter", value)
}

// Desc
func (t TextControl) Desc(value string) TextControl {
	return t.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (t TextControl) Description(value string) TextControl {
	return t.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (t TextControl) DescriptionClassName(value string) TextControl {
	return t.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (t TextControl) Disabled(value bool) TextControl {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t TextControl) DisabledOn(value string) TextControl {
	return t.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (t TextControl) EditApi(value string) TextControl {
	return t.set("editApi", value)
}

// EditControls 选项修改的表单项
func (t TextControl) EditControls(value string) TextControl {
	return t.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (t TextControl) EditDialog(value string) TextControl {
	return t.set("editDialog", value)
}

// Editable 是否可以编辑
func (t TextControl) Editable(value bool) TextControl {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t TextControl) EditorSetting(value string) TextControl {
	return t.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (t TextControl) ExtraName(value string) TextControl {
	return t.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (t TextControl) ExtractValue(value bool) TextControl {
	return t.set("extractValue", value)
}

// Hidden 是否隐藏
func (t TextControl) Hidden(value bool) TextControl {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t TextControl) HiddenOn(value string) TextControl {
	return t.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (t TextControl) Hint(value string) TextControl {
	return t.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。
func (t TextControl) Horizontal(value string) TextControl {
	return t.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t TextControl) Id(value string) TextControl {
	return t.set("id", value)
}

// InitAutoFill
func (t TextControl) InitAutoFill(value string) TextControl {
	return t.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (t TextControl) InitFetch(value bool) TextControl {
	return t.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (t TextControl) InitFetchOn(value string) TextControl {
	return t.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式。
func (t TextControl) Inline(value bool) TextControl {
	return t.set("inline", value)
}

// InputClassName 配置 input className
func (t TextControl) InputClassName(value string) TextControl {
	return t.set("inputClassName", value)
}

// InputControlClassName control节点的CSS类名
func (t TextControl) InputControlClassName(value string) TextControl {
	return t.set("inputControlClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (t TextControl) JoinValues(value bool) TextControl {
	return t.set("joinValues", value)
}

// Label 描述标题
func (t TextControl) Label(value string) TextControl {
	return t.set("label", value)
}

// LabelAlign 描述标题 可选值: right | left | top | inherit
func (t TextControl) LabelAlign(value string) TextControl {
	return t.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (t TextControl) LabelClassName(value string) TextControl {
	return t.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (t TextControl) LabelRemark(value string) TextControl {
	return t.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (t TextControl) LabelWidth(value string) TextControl {
	return t.set("labelWidth", value)
}

// MaxLength 限制文字最大输入个数
func (t TextControl) MaxLength(value string) TextControl {
	return t.set("maxLength", value)
}

// MinLength 限制文字最小输入个数
func (t TextControl) MinLength(value string) TextControl {
	return t.set("minLength", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (t TextControl) Mode(value string) TextControl {
	return t.set("mode", value)
}

// Multiple 是否为多选模式
func (t TextControl) Multiple(value bool) TextControl {
	return t.set("multiple", value)
}

// Name 表单项名称
func (t TextControl) Name(value string) TextControl {
	return t.set("name", value)
}

// Note
func (t TextControl) Note(value string) TextControl {
	return t.set("note", value)
}

// OnChange 触发的事件
func (t TextControl) OnChange(value string) TextControl {
	return t.set("onChange", value)
}

// OnFocus 触发的事件
func (t TextControl) OnFocus(value string) TextControl {
	return t.set("onFocus", value)
}

// OnBlur 触发的事件
func (t TextControl) OnBlur(value string) TextControl {
	return t.set("onBlur", value)
}

// OnInit 触发的事件
func (t TextControl) OnInit(value string) TextControl {
	return t.set("onInit", value)
}

// OnSearch 触发的事件
func (t TextControl) OnSearch(value string) TextControl {
	return t.set("onSearch", value)
}

// Option
func (t TextControl) Option(value string) TextControl {
	return t.set("option", value)
}

// Options 选项列表
func (t TextControl) Options(value string) TextControl {
	return t.set("options", value)
}

// OptionLabel 在多选模式下，选项的 label 位置。可选值: top | left | right
func (t TextControl) OptionLabel(value string) TextControl {
	return t.set("optionLabel", value)
}

// Placeholder 提示文字
func (t TextControl) Placeholder(value string) TextControl {
	return t.set("placeholder", value)
}

// ReadOnly 是否为只读模式
func (t TextControl) ReadOnly(value bool) TextControl {
	return t.set("readOnly", value)
}

// RefixApi 配置选择项的接口
func (t TextControl) RefixApi(value string) TextControl {
	return t.set("refixApi", value)
}

// ResetValue
func (t TextControl) ResetValue(value string) TextControl {
	return t.set("resetValue", value)
}

// Rules 校验规则
func (t TextControl) Rules(value string) TextControl {
	return t.set("rules", value)
}

// SearchApi 搜索接口
func (t TextControl) SearchApi(value string) TextControl {
	return t.set("searchApi", value)
}

// Source 数据来源
func (t TextControl) Source(value string) TextControl {
	return t.set("source", value)
}

// StaticClassName 用于静态展示的类名
func (t TextControl) StaticClassName(value string) TextControl {
	return t.set("staticClassName", value)
}

// StaticLabel 静态显示的标题
func (t TextControl) StaticLabel(value string) TextControl {
	return t.set("staticLabel", value)
}

// StaticOn 静态显示的条件
func (t TextControl) StaticOn(value string) TextControl {
	return t.set("staticOn", value)
}

// StaticValue 静态值
func (t TextControl) StaticValue(value string) TextControl {
	return t.set("staticValue", value)
}

// Tips 提示文本
func (t TextControl) Tips(value string) TextControl {
	return t.set("tips", value)
}

// ValidationErrors 校验不通过时的提示信息
func (t TextControl) ValidationErrors(value string) TextControl {
	return t.set("validationErrors", value)
}

// Value 表单项的值
func (t TextControl) Value(value interface{}) TextControl {
	return t.set("value", value)
}

// Visible 控制组件是否可见
func (t TextControl) Visible(value bool) TextControl {
	return t.set("visible", value)
}

// VisibleOn 是否可见表达式 (表达式，语法 `data.xxx > 5`。)
func (t TextControl) VisibleOn(value string) TextControl {
	return t.set("visibleOn", value)
}

// Width 控件宽度
func (t TextControl) Width(value string) TextControl {
	return t.set("width", value)
}

// OnInput 触发的事件
func (t TextControl) OnInput(value string) TextControl {
	return t.set("onInput", value)
}

// ValueType
func (t TextControl) ValueType(value string) TextControl {
	return t.set("valueType", value)
}

// ValueOnEvent
func (t TextControl) ValueOnEvent(value string) TextControl {
	return t.set("valueOnEvent", value)
}

// set 设置属性
func (t TextControl) set(key string, value interface{}) TextControl {
	t[key] = value
	return t
}
