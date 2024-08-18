package comp

// ListControl 列表控件
//
// @version 6.7.0
type ListControl Schema

// NewListControl 创建一个新的 ListControl 实例
func NewListControl() ListControl {
	return make(ListControl).set("type", "list-select")
}

// set 设置字段值
func (lc ListControl) set(key string, value interface{}) ListControl {
	lc[key] = value
	return lc
}

// ActiveItemSchema 激活态自定义展示模板
func (lc ListControl) ActiveItemSchema(value string) ListControl {
	return lc.set("activeItemSchema", value)
}

// AddApi 添加时调用的接口
func (lc ListControl) AddApi(value string) ListControl {
	return lc.set("addApi", value)
}

// AddControls 新增时的表单项
func (lc ListControl) AddControls(value string) ListControl {
	return lc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (lc ListControl) AddDialog(value string) ListControl {
	return lc.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (lc ListControl) AutoFill(value string) ListControl {
	return lc.set("autoFill", value)
}

// ClassName 容器 css 类名
func (lc ListControl) ClassName(value string) ListControl {
	return lc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (lc ListControl) ClearValueOnHidden(value bool) ListControl {
	return lc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (lc ListControl) Clearable(value bool) ListControl {
	return lc.set("clearable", value)
}

// Creatable 是否可以新增
func (lc ListControl) Creatable(value bool) ListControl {
	return lc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (lc ListControl) CreateBtnLabel(value string) ListControl {
	return lc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (lc ListControl) DeferApi(value string) ListControl {
	return lc.set("deferApi", value)
}

// DeferField 懒加载字段
func (lc ListControl) DeferField(value string) ListControl {
	return lc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (lc ListControl) DeleteApi(value string) ListControl {
	return lc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (lc ListControl) DeleteConfirmText(value string) ListControl {
	return lc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (lc ListControl) Delimiter(value string) ListControl {
	return lc.set("delimiter", value)
}

// Desc 描述内容
func (lc ListControl) Desc(value string) ListControl {
	return lc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (lc ListControl) Description(value string) ListControl {
	return lc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (lc ListControl) DescriptionClassName(value string) ListControl {
	return lc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (lc ListControl) Disabled(value bool) ListControl {
	return lc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (lc ListControl) DisabledOn(value string) ListControl {
	return lc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (lc ListControl) EditApi(value string) ListControl {
	return lc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (lc ListControl) EditControls(value string) ListControl {
	return lc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (lc ListControl) EditDialog(value string) ListControl {
	return lc.set("editDialog", value)
}

// Editable 是否可以编辑
func (lc ListControl) Editable(value bool) ListControl {
	return lc.set("editable", value)
}

// EditorSetting 编辑器配置
func (lc ListControl) EditorSetting(value string) ListControl {
	return lc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (lc ListControl) ExtraName(value string) ListControl {
	return lc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (lc ListControl) ExtractValue(value bool) ListControl {
	return lc.set("extractValue", value)
}

// Hidden 是否隐藏
func (lc ListControl) Hidden(value bool) ListControl {
	return lc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (lc ListControl) HiddenOn(value string) ListControl {
	return lc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (lc ListControl) Hint(value string) ListControl {
	return lc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (lc ListControl) Horizontal(value string) ListControl {
	return lc.set("horizontal", value)
}

// Id 组件唯一 id
func (lc ListControl) Id(value string) ListControl {
	return lc.set("id", value)
}

// ImageClassName 图片div类名
func (lc ListControl) ImageClassName(value string) ListControl {
	return lc.set("imageClassName", value)
}

// InitAutoFill 初始化自动填充
func (lc ListControl) InitAutoFill(value string) ListControl {
	return lc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (lc ListControl) InitFetch(value bool) ListControl {
	return lc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (lc ListControl) InitFetchOn(value string) ListControl {
	return lc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (lc ListControl) Inline(value bool) ListControl {
	return lc.set("inline", value)
}

// InputClassName 配置 input className
func (lc ListControl) InputClassName(value string) ListControl {
	return lc.set("inputClassName", value)
}

// ItemSchema 可以自定义展示模板
func (lc ListControl) ItemSchema(value string) ListControl {
	return lc.set("itemSchema", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交
func (lc ListControl) JoinValues(value bool) ListControl {
	return lc.set("joinValues", value)
}

// Label 描述标题
func (lc ListControl) Label(value string) ListControl {
	return lc.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (lc ListControl) LabelAlign(value string) ListControl {
	return lc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (lc ListControl) LabelClassName(value string) ListControl {
	return lc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (lc ListControl) LabelRemark(value string) ListControl {
	return lc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (lc ListControl) LabelWidth(value string) ListControl {
	return lc.set("labelWidth", value)
}

// ListClassName 支持配置 list div 的 css 类名
func (lc ListControl) ListClassName(value string) ListControl {
	return lc.set("listClassName", value)
}

// Mode 配置当前表单项展示模式
func (lc ListControl) Mode(value string) ListControl {
	return lc.set("mode", value)
}

// Multiple 是否为多选模式
func (lc ListControl) Multiple(value bool) ListControl {
	return lc.set("multiple", value)
}

// Name 字段名，表单提交时的 key
func (lc ListControl) Name(value string) ListControl {
	return lc.set("name", value)
}

// OnEvent 事件动作配置
func (lc ListControl) OnEvent(value string) ListControl {
	return lc.set("onEvent", value)
}

// Options 选项集合
func (lc ListControl) Options(value string) ListControl {
	return lc.set("options", value)
}

// Placeholder 占位符
func (lc ListControl) Placeholder(value string) ListControl {
	return lc.set("placeholder", value)
}

// ReadOnly 是否只读
func (lc ListControl) ReadOnly(value bool) ListControl {
	return lc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (lc ListControl) ReadOnlyOn(value string) ListControl {
	return lc.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (lc ListControl) Remark(value string) ListControl {
	return lc.set("remark", value)
}

// Removable 是否可删除
func (lc ListControl) Removable(value bool) ListControl {
	return lc.set("removable", value)
}

// Required 是否为必填
func (lc ListControl) Required(value bool) ListControl {
	return lc.set("required", value)
}

// ResetValueOnHidden 是否在隐藏时重置值
func (lc ListControl) ResetValueOnHidden(value bool) ListControl {
	return lc.set("resetValueOnHidden", value)
}

// Size 配置尺寸
func (lc ListControl) Size(value string) ListControl {
	return lc.set("size", value)
}

// Source 配置选项数据源接口
func (lc ListControl) Source(value string) ListControl {
	return lc.set("source", value)
}

// Value 配置表单项初始值
func (lc ListControl) Value(value string) ListControl {
	return lc.set("value", value)
}

// ValueField 设置选项 value 字段名
func (lc ListControl) ValueField(value string) ListControl {
	return lc.set("valueField", value)
}

// Visible 配置是否可见
func (lc ListControl) Visible(value bool) ListControl {
	return lc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (lc ListControl) VisibleOn(value string) ListControl {
	return lc.set("visibleOn", value)
}

// Width 配置控件宽度
func (lc ListControl) Width(value string) ListControl {
	return lc.set("width", value)
}
