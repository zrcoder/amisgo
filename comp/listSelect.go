package comp

// listSelect 列表控件

type listSelect Schema

func ListSelect() listSelect {
	return make(listSelect).set("type", "list-select")
}

// set 设置字段值
func (lc listSelect) set(key string, value any) listSelect {
	lc[key] = value
	return lc
}

// ActiveItemSchema 激活态自定义展示模板
func (lc listSelect) ActiveItemSchema(value string) listSelect {
	return lc.set("activeItemSchema", value)
}

// AddApi 添加时调用的接口
func (lc listSelect) AddApi(value string) listSelect {
	return lc.set("addApi", value)
}

// AddControls 新增时的表单项
func (lc listSelect) AddControls(value string) listSelect {
	return lc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (lc listSelect) AddDialog(value string) listSelect {
	return lc.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (lc listSelect) AutoFill(value string) listSelect {
	return lc.set("autoFill", value)
}

// ClassName 容器 css 类名
func (lc listSelect) ClassName(value string) listSelect {
	return lc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (lc listSelect) ClearValueOnHidden(value bool) listSelect {
	return lc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (lc listSelect) Clearable(value bool) listSelect {
	return lc.set("clearable", value)
}

// Creatable 是否可以新增
func (lc listSelect) Creatable(value bool) listSelect {
	return lc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (lc listSelect) CreateBtnLabel(value string) listSelect {
	return lc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (lc listSelect) DeferApi(value string) listSelect {
	return lc.set("deferApi", value)
}

// DeferField 懒加载字段
func (lc listSelect) DeferField(value string) listSelect {
	return lc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (lc listSelect) DeleteApi(value string) listSelect {
	return lc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (lc listSelect) DeleteConfirmText(value string) listSelect {
	return lc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (lc listSelect) Delimiter(value string) listSelect {
	return lc.set("delimiter", value)
}

// Desc 描述内容
func (lc listSelect) Desc(value string) listSelect {
	return lc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (lc listSelect) Description(value string) listSelect {
	return lc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (lc listSelect) DescriptionClassName(value string) listSelect {
	return lc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (lc listSelect) Disabled(value bool) listSelect {
	return lc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (lc listSelect) DisabledOn(value string) listSelect {
	return lc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (lc listSelect) EditApi(value string) listSelect {
	return lc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (lc listSelect) EditControls(value string) listSelect {
	return lc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (lc listSelect) EditDialog(value string) listSelect {
	return lc.set("editDialog", value)
}

// Editable 是否可以编辑
func (lc listSelect) Editable(value bool) listSelect {
	return lc.set("editable", value)
}

// EditorSetting 编辑器配置
func (lc listSelect) EditorSetting(value string) listSelect {
	return lc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (lc listSelect) ExtraName(value string) listSelect {
	return lc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (lc listSelect) ExtractValue(value bool) listSelect {
	return lc.set("extractValue", value)
}

// Hidden 是否隐藏
func (lc listSelect) Hidden(value bool) listSelect {
	return lc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (lc listSelect) HiddenOn(value string) listSelect {
	return lc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (lc listSelect) Hint(value string) listSelect {
	return lc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (lc listSelect) Horizontal(value string) listSelect {
	return lc.set("horizontal", value)
}

// Id 组件唯一 id
func (lc listSelect) ID(value string) listSelect {
	return lc.set("id", value)
}

// ImageClassName 图片div类名
func (lc listSelect) ImageClassName(value string) listSelect {
	return lc.set("imageClassName", value)
}

// InitAutoFill 初始化自动填充
func (lc listSelect) InitAutoFill(value string) listSelect {
	return lc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (lc listSelect) InitFetch(value bool) listSelect {
	return lc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (lc listSelect) InitFetchOn(value string) listSelect {
	return lc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (lc listSelect) Inline(value bool) listSelect {
	return lc.set("inline", value)
}

// InputClassName 配置 input className
func (lc listSelect) InputClassName(value string) listSelect {
	return lc.set("inputClassName", value)
}

// ItemSchema 可以自定义展示模板
func (lc listSelect) ItemSchema(value string) listSelect {
	return lc.set("itemSchema", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交
func (lc listSelect) JoinValues(value bool) listSelect {
	return lc.set("joinValues", value)
}

// Label 描述标题
func (lc listSelect) Label(value string) listSelect {
	return lc.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (lc listSelect) LabelAlign(value string) listSelect {
	return lc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (lc listSelect) LabelClassName(value string) listSelect {
	return lc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (lc listSelect) LabelRemark(value string) listSelect {
	return lc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (lc listSelect) LabelWidth(value string) listSelect {
	return lc.set("labelWidth", value)
}

// ListClassName 支持配置 list div 的 css 类名
func (lc listSelect) ListClassName(value string) listSelect {
	return lc.set("listClassName", value)
}

// Mode 配置当前表单项展示模式
func (lc listSelect) Mode(value string) listSelect {
	return lc.set("mode", value)
}

// Multiple 是否为多选模式
func (lc listSelect) Multiple(value bool) listSelect {
	return lc.set("multiple", value)
}

// Name 字段名，表单提交时的 key
func (lc listSelect) Name(value string) listSelect {
	return lc.set("name", value)
}

// OnEvent 事件动作配置
func (lc listSelect) OnEvent(value any) listSelect {
	return lc.set("onEvent", value)
}

// Options 选项集合
func (lc listSelect) Options(value ...MOption) listSelect {
	return lc.set("options", value)
}

// Placeholder 占位符
func (lc listSelect) Placeholder(value string) listSelect {
	return lc.set("placeholder", value)
}

// ReadOnly 是否只读
func (lc listSelect) ReadOnly(value bool) listSelect {
	return lc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (lc listSelect) ReadOnlyOn(value string) listSelect {
	return lc.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (lc listSelect) Remark(value string) listSelect {
	return lc.set("remark", value)
}

// Removable 是否可删除
func (lc listSelect) Removable(value bool) listSelect {
	return lc.set("removable", value)
}

// Required 是否为必填
func (lc listSelect) Required(value bool) listSelect {
	return lc.set("required", value)
}

// ResetValueOnHidden 是否在隐藏时重置值
func (lc listSelect) ResetValueOnHidden(value bool) listSelect {
	return lc.set("resetValueOnHidden", value)
}

// Size 配置尺寸
func (lc listSelect) Size(value string) listSelect {
	return lc.set("size", value)
}

// Source 配置选项数据源接口
func (lc listSelect) Source(value string) listSelect {
	return lc.set("source", value)
}

// Value 配置表单项初始值
func (lc listSelect) Value(value string) listSelect {
	return lc.set("value", value)
}

// ValueField 设置选项 value 字段名
func (lc listSelect) ValueField(value string) listSelect {
	return lc.set("valueField", value)
}

// Visible 配置是否可见
func (lc listSelect) Visible(value bool) listSelect {
	return lc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (lc listSelect) VisibleOn(value string) listSelect {
	return lc.set("visibleOn", value)
}

// Width 配置控件宽度
func (lc listSelect) Width(value string) listSelect {
	return lc.set("width", value)
}
