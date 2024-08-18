package comp

// UserSelectControl 移动端人员选择。
//
// @author  slowlyo
// @version 6.7.0
type UserSelectControl Schema

// NewUserSelectControl 创建一个新的 UserSelectControl 实例
func NewUserSelectControl() UserSelectControl {
	return UserSelectControl{}.set("type", "users-select")
}

func (usc UserSelectControl) set(key string, value interface{}) UserSelectControl {
	usc[key] = value
	return usc
}

// AddApi 添加时调用的接口
func (usc UserSelectControl) AddApi(value string) UserSelectControl {
	return usc.set("addApi", value)
}

// AddControls 新增时的表单项
func (usc UserSelectControl) AddControls(value string) UserSelectControl {
	return usc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (usc UserSelectControl) AddDialog(value string) UserSelectControl {
	return usc.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (usc UserSelectControl) AutoFill(value string) UserSelectControl {
	return usc.set("autoFill", value)
}

// ClassName 容器 css 类名
func (usc UserSelectControl) ClassName(value string) UserSelectControl {
	return usc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (usc UserSelectControl) ClearValueOnHidden(value bool) UserSelectControl {
	return usc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (usc UserSelectControl) Clearable(value bool) UserSelectControl {
	return usc.set("clearable", value)
}

// Creatable 是否可以新增
func (usc UserSelectControl) Creatable(value bool) UserSelectControl {
	return usc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (usc UserSelectControl) CreateBtnLabel(value string) UserSelectControl {
	return usc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (usc UserSelectControl) DeferApi(value string) UserSelectControl {
	return usc.set("deferApi", value)
}

// DeferField 懒加载字段
func (usc UserSelectControl) DeferField(value string) UserSelectControl {
	return usc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (usc UserSelectControl) DeleteApi(value string) UserSelectControl {
	return usc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (usc UserSelectControl) DeleteConfirmText(value string) UserSelectControl {
	return usc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (usc UserSelectControl) Delimiter(value string) UserSelectControl {
	return usc.set("delimiter", value)
}

// Desc
func (usc UserSelectControl) Desc(value string) UserSelectControl {
	return usc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (usc UserSelectControl) Description(value string) UserSelectControl {
	return usc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (usc UserSelectControl) DescriptionClassName(value string) UserSelectControl {
	return usc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (usc UserSelectControl) Disabled(value bool) UserSelectControl {
	return usc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (usc UserSelectControl) DisabledOn(value string) UserSelectControl {
	return usc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (usc UserSelectControl) EditApi(value string) UserSelectControl {
	return usc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (usc UserSelectControl) EditControls(value string) UserSelectControl {
	return usc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (usc UserSelectControl) EditDialog(value string) UserSelectControl {
	return usc.set("editDialog", value)
}

// Editable 是否可以编辑
func (usc UserSelectControl) Editable(value bool) UserSelectControl {
	return usc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (usc UserSelectControl) EditorSetting(value string) UserSelectControl {
	return usc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (usc UserSelectControl) ExtraName(value string) UserSelectControl {
	return usc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值
func (usc UserSelectControl) ExtractValue(value bool) UserSelectControl {
	return usc.set("extractValue", value)
}

// Hidden 是否隐藏
func (usc UserSelectControl) Hidden(value bool) UserSelectControl {
	return usc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (usc UserSelectControl) HiddenOn(value string) UserSelectControl {
	return usc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (usc UserSelectControl) Hint(value string) UserSelectControl {
	return usc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (usc UserSelectControl) Horizontal(value string) UserSelectControl {
	return usc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (usc UserSelectControl) Id(value string) UserSelectControl {
	return usc.set("id", value)
}

// InitAutoFill
func (usc UserSelectControl) InitAutoFill(value string) UserSelectControl {
	return usc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (usc UserSelectControl) InitFetch(value bool) UserSelectControl {
	return usc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (usc UserSelectControl) InitFetchOn(value string) UserSelectControl {
	return usc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (usc UserSelectControl) Inline(value bool) UserSelectControl {
	return usc.set("inline", value)
}

// InputClassName 配置 input className
func (usc UserSelectControl) InputClassName(value string) UserSelectControl {
	return usc.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式下的值处理
func (usc UserSelectControl) JoinValues(value bool) UserSelectControl {
	return usc.set("joinValues", value)
}

// Label 描述标题
func (usc UserSelectControl) Label(value string) UserSelectControl {
	return usc.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (usc UserSelectControl) LabelAlign(value string) UserSelectControl {
	return usc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (usc UserSelectControl) LabelClassName(value string) UserSelectControl {
	return usc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (usc UserSelectControl) LabelRemark(value string) UserSelectControl {
	return usc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (usc UserSelectControl) LabelWidth(value string) UserSelectControl {
	return usc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (usc UserSelectControl) Mode(value string) UserSelectControl {
	return usc.set("mode", value)
}

// Multiple 是否为多选模式
func (usc UserSelectControl) Multiple(value bool) UserSelectControl {
	return usc.set("multiple", value)
}

// Name 字段名，表单提交时的 key
func (usc UserSelectControl) Name(value string) UserSelectControl {
	return usc.set("name", value)
}

// OnEvent 事件动作配置
func (usc UserSelectControl) OnEvent(value string) UserSelectControl {
	return usc.set("onEvent", value)
}

// Options 选项集合
func (usc UserSelectControl) Options(value string) UserSelectControl {
	return usc.set("options", value)
}

// Placeholder 占位符
func (usc UserSelectControl) Placeholder(value string) UserSelectControl {
	return usc.set("placeholder", value)
}

// ReadOnly 是否只读
func (usc UserSelectControl) ReadOnly(value bool) UserSelectControl {
	return usc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (usc UserSelectControl) ReadOnlyOn(value string) UserSelectControl {
	return usc.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (usc UserSelectControl) Remark(value string) UserSelectControl {
	return usc.set("remark", value)
}

// Removable 是否可删除
func (usc UserSelectControl) Removable(value bool) UserSelectControl {
	return usc.set("removable", value)
}

// Required 是否为必填
func (usc UserSelectControl) Required(value bool) UserSelectControl {
	return usc.set("required", value)
}

// ResetValue 点清除按钮时，将表单项重置成的值
func (usc UserSelectControl) ResetValue(value string) UserSelectControl {
	return usc.set("resetValue", value)
}

// ShowErrorMsg 验证失败时是否展示错误信息
func (usc UserSelectControl) ShowErrorMsg(value bool) UserSelectControl {
	return usc.set("showErrorMsg", value)
}

// Size 输入框大小
func (usc UserSelectControl) Size(value string) UserSelectControl {
	return usc.set("size", value)
}

// StaticClassName 配置静态内容 className
func (usc UserSelectControl) StaticClassName(value string) UserSelectControl {
	return usc.set("staticClassName", value)
}

// StaticLabel 静态展示时 label
func (usc UserSelectControl) StaticLabel(value string) UserSelectControl {
	return usc.set("staticLabel", value)
}

// StaticPlaceholder 静态展示时 placeholder
func (usc UserSelectControl) StaticPlaceholder(value string) UserSelectControl {
	return usc.set("staticPlaceholder", value)
}

// StaticOptions 静态选项集合
func (usc UserSelectControl) StaticOptions(value string) UserSelectControl {
	return usc.set("staticOptions", value)
}

// Width 输入框宽度
func (usc UserSelectControl) Width(value string) UserSelectControl {
	return usc.set("width", value)
}
