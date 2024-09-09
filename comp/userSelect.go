package comp

// userSelect 移动端人员选择。

type userSelect schema

func UsersSelect() userSelect {
	return userSelect{}.set("type", "users-select")
}

func (usc userSelect) set(key string, value any) userSelect {
	usc[key] = value
	return usc
}

// AddApi 添加时调用的接口
func (usc userSelect) AddApi(value string) userSelect {
	return usc.set("addApi", value)
}

// AddControls 新增时的表单项
func (usc userSelect) AddControls(value string) userSelect {
	return usc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (usc userSelect) AddDialog(value string) userSelect {
	return usc.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (usc userSelect) AutoFill(value string) userSelect {
	return usc.set("autoFill", value)
}

// ClassName 容器 css 类名
func (usc userSelect) ClassName(value string) userSelect {
	return usc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (usc userSelect) ClearValueOnHidden(value bool) userSelect {
	return usc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (usc userSelect) Clearable(value bool) userSelect {
	return usc.set("clearable", value)
}

// Creatable 是否可以新增
func (usc userSelect) Creatable(value bool) userSelect {
	return usc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (usc userSelect) CreateBtnLabel(value string) userSelect {
	return usc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (usc userSelect) DeferApi(value string) userSelect {
	return usc.set("deferApi", value)
}

// DeferField 懒加载字段
func (usc userSelect) DeferField(value string) userSelect {
	return usc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (usc userSelect) DeleteApi(value string) userSelect {
	return usc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (usc userSelect) DeleteConfirmText(value string) userSelect {
	return usc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (usc userSelect) Delimiter(value string) userSelect {
	return usc.set("delimiter", value)
}

// Desc
func (usc userSelect) Desc(value string) userSelect {
	return usc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (usc userSelect) Description(value string) userSelect {
	return usc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (usc userSelect) DescriptionClassName(value string) userSelect {
	return usc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (usc userSelect) Disabled(value bool) userSelect {
	return usc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (usc userSelect) DisabledOn(value string) userSelect {
	return usc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (usc userSelect) EditApi(value string) userSelect {
	return usc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (usc userSelect) EditControls(value string) userSelect {
	return usc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (usc userSelect) EditDialog(value string) userSelect {
	return usc.set("editDialog", value)
}

// Editable 是否可以编辑
func (usc userSelect) Editable(value bool) userSelect {
	return usc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (usc userSelect) EditorSetting(value string) userSelect {
	return usc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (usc userSelect) ExtraName(value string) userSelect {
	return usc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值
func (usc userSelect) ExtractValue(value bool) userSelect {
	return usc.set("extractValue", value)
}

// Hidden 是否隐藏
func (usc userSelect) Hidden(value bool) userSelect {
	return usc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (usc userSelect) HiddenOn(value string) userSelect {
	return usc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (usc userSelect) Hint(value string) userSelect {
	return usc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (usc userSelect) Horizontal(value string) userSelect {
	return usc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (usc userSelect) Id(value string) userSelect {
	return usc.set("id", value)
}

// InitAutoFill
func (usc userSelect) InitAutoFill(value string) userSelect {
	return usc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (usc userSelect) InitFetch(value bool) userSelect {
	return usc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (usc userSelect) InitFetchOn(value string) userSelect {
	return usc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (usc userSelect) Inline(value bool) userSelect {
	return usc.set("inline", value)
}

// InputClassName 配置 input className
func (usc userSelect) InputClassName(value string) userSelect {
	return usc.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式下的值处理
func (usc userSelect) JoinValues(value bool) userSelect {
	return usc.set("joinValues", value)
}

// Label 描述标题
func (usc userSelect) Label(value string) userSelect {
	return usc.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (usc userSelect) LabelAlign(value string) userSelect {
	return usc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (usc userSelect) LabelClassName(value string) userSelect {
	return usc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (usc userSelect) LabelRemark(value string) userSelect {
	return usc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (usc userSelect) LabelWidth(value string) userSelect {
	return usc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (usc userSelect) Mode(value string) userSelect {
	return usc.set("mode", value)
}

// Multiple 是否为多选模式
func (usc userSelect) Multiple(value bool) userSelect {
	return usc.set("multiple", value)
}

// Name 字段名，表单提交时的 key
func (usc userSelect) Name(value string) userSelect {
	return usc.set("name", value)
}

// OnEvent 事件动作配置
func (usc userSelect) OnEvent(value any) userSelect {
	return usc.set("onEvent", value)
}

// Options 选项集合
func (usc userSelect) Options(value ...option) userSelect {
	return usc.set("options", value)
}

// Placeholder 占位符
func (usc userSelect) Placeholder(value string) userSelect {
	return usc.set("placeholder", value)
}

// ReadOnly 是否只读
func (usc userSelect) ReadOnly(value bool) userSelect {
	return usc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (usc userSelect) ReadOnlyOn(value string) userSelect {
	return usc.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (usc userSelect) Remark(value string) userSelect {
	return usc.set("remark", value)
}

// Removable 是否可删除
func (usc userSelect) Removable(value bool) userSelect {
	return usc.set("removable", value)
}

// Required 是否为必填
func (usc userSelect) Required(value bool) userSelect {
	return usc.set("required", value)
}

// ResetValue 点清除按钮时，将表单项重置成的值
func (usc userSelect) ResetValue(value string) userSelect {
	return usc.set("resetValue", value)
}

// ShowErrorMsg 验证失败时是否展示错误信息
func (usc userSelect) ShowErrorMsg(value bool) userSelect {
	return usc.set("showErrorMsg", value)
}

// Size 输入框大小
func (usc userSelect) Size(value string) userSelect {
	return usc.set("size", value)
}

// StaticClassName 配置静态内容 className
func (usc userSelect) StaticClassName(value string) userSelect {
	return usc.set("staticClassName", value)
}

// StaticLabel 静态展示时 label
func (usc userSelect) StaticLabel(value string) userSelect {
	return usc.set("staticLabel", value)
}

// StaticPlaceholder 静态展示时 placeholder
func (usc userSelect) StaticPlaceholder(value string) userSelect {
	return usc.set("staticPlaceholder", value)
}

// StaticOptions 静态选项集合
func (usc userSelect) StaticOptions(value ...option) userSelect {
	return usc.set("staticOptions", value)
}

// Width 输入框宽度
func (usc userSelect) Width(value string) userSelect {
	return usc.set("width", value)
}
