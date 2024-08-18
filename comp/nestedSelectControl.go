package comp

// NestedSelectControl 嵌套选择控件
//
// @version 6.7.0
type NestedSelectControl Schema

// NewNestedSelectControl 创建一个新的 NestedSelectControl 实例
func NewNestedSelectControl() NestedSelectControl {
	return make(NestedSelectControl).set("type", "nested-select")
}

// set 设置字段值
func (nsc NestedSelectControl) set(key string, value interface{}) NestedSelectControl {
	nsc[key] = value
	return nsc
}

// AddApi 添加时调用的接口
func (nsc NestedSelectControl) AddApi(value string) NestedSelectControl {
	return nsc.set("addApi", value)
}

// AddControls 新增时的表单项
func (nsc NestedSelectControl) AddControls(value string) NestedSelectControl {
	return nsc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (nsc NestedSelectControl) AddDialog(value string) NestedSelectControl {
	return nsc.set("addDialog", value)
}

// AutoFill 自动填充
func (nsc NestedSelectControl) AutoFill(value string) NestedSelectControl {
	return nsc.set("autoFill", value)
}

// BorderMode 边框模式
func (nsc NestedSelectControl) BorderMode(value string) NestedSelectControl {
	return nsc.set("borderMode", value)
}

// Cascade 父子之间是否完全独立
func (nsc NestedSelectControl) Cascade(value bool) NestedSelectControl {
	return nsc.set("cascade", value)
}

// ClassName 容器 css 类名
func (nsc NestedSelectControl) ClassName(value string) NestedSelectControl {
	return nsc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (nsc NestedSelectControl) ClearValueOnHidden(value bool) NestedSelectControl {
	return nsc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (nsc NestedSelectControl) Clearable(value bool) NestedSelectControl {
	return nsc.set("clearable", value)
}

// Creatable 是否可以新增
func (nsc NestedSelectControl) Creatable(value bool) NestedSelectControl {
	return nsc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (nsc NestedSelectControl) CreateBtnLabel(value string) NestedSelectControl {
	return nsc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (nsc NestedSelectControl) DeferApi(value string) NestedSelectControl {
	return nsc.set("deferApi", value)
}

// DeferField 懒加载字段
func (nsc NestedSelectControl) DeferField(value string) NestedSelectControl {
	return nsc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (nsc NestedSelectControl) DeleteApi(value string) NestedSelectControl {
	return nsc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (nsc NestedSelectControl) DeleteConfirmText(value string) NestedSelectControl {
	return nsc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (nsc NestedSelectControl) Delimiter(value string) NestedSelectControl {
	return nsc.set("delimiter", value)
}

// Desc 描述
func (nsc NestedSelectControl) Desc(value string) NestedSelectControl {
	return nsc.set("desc", value)
}

// Description 描述内容
func (nsc NestedSelectControl) Description(value string) NestedSelectControl {
	return nsc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (nsc NestedSelectControl) DescriptionClassName(value string) NestedSelectControl {
	return nsc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (nsc NestedSelectControl) Disabled(value bool) NestedSelectControl {
	return nsc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (nsc NestedSelectControl) DisabledOn(value string) NestedSelectControl {
	return nsc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (nsc NestedSelectControl) EditApi(value string) NestedSelectControl {
	return nsc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (nsc NestedSelectControl) EditControls(value string) NestedSelectControl {
	return nsc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (nsc NestedSelectControl) EditDialog(value string) NestedSelectControl {
	return nsc.set("editDialog", value)
}

// Editable 是否可以编辑
func (nsc NestedSelectControl) Editable(value bool) NestedSelectControl {
	return nsc.set("editable", value)
}

// EditorSetting 编辑器配置
func (nsc NestedSelectControl) EditorSetting(value string) NestedSelectControl {
	return nsc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (nsc NestedSelectControl) ExtraName(value string) NestedSelectControl {
	return nsc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (nsc NestedSelectControl) ExtractValue(value bool) NestedSelectControl {
	return nsc.set("extractValue", value)
}

// Hidden 是否隐藏
func (nsc NestedSelectControl) Hidden(value bool) NestedSelectControl {
	return nsc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (nsc NestedSelectControl) HiddenOn(value string) NestedSelectControl {
	return nsc.set("hiddenOn", value)
}

// HideNodePathLabel 是否隐藏选择框中已选中节点的祖先节点的文本信息
func (nsc NestedSelectControl) HideNodePathLabel(value bool) NestedSelectControl {
	return nsc.set("hideNodePathLabel", value)
}

// Hint 输入提示
func (nsc NestedSelectControl) Hint(value string) NestedSelectControl {
	return nsc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候
func (nsc NestedSelectControl) Horizontal(value string) NestedSelectControl {
	return nsc.set("horizontal", value)
}

// Id 组件唯一 id
func (nsc NestedSelectControl) Id(value string) NestedSelectControl {
	return nsc.set("id", value)
}

// InitAutoFill 初始化自动填充
func (nsc NestedSelectControl) InitAutoFill(value string) NestedSelectControl {
	return nsc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (nsc NestedSelectControl) InitFetch(value bool) NestedSelectControl {
	return nsc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (nsc NestedSelectControl) InitFetchOn(value string) NestedSelectControl {
	return nsc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (nsc NestedSelectControl) Inline(value bool) NestedSelectControl {
	return nsc.set("inline", value)
}

// InputClassName 配置 input className
func (nsc NestedSelectControl) InputClassName(value string) NestedSelectControl {
	return nsc.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式
func (nsc NestedSelectControl) JoinValues(value bool) NestedSelectControl {
	return nsc.set("joinValues", value)
}

// Label 描述标题
func (nsc NestedSelectControl) Label(value string) NestedSelectControl {
	return nsc.set("label", value)
}

// LabelAlign 描述标题对齐
func (nsc NestedSelectControl) LabelAlign(value string) NestedSelectControl {
	return nsc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (nsc NestedSelectControl) LabelClassName(value string) NestedSelectControl {
	return nsc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标
func (nsc NestedSelectControl) LabelRemark(value string) NestedSelectControl {
	return nsc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (nsc NestedSelectControl) LabelWidth(value string) NestedSelectControl {
	return nsc.set("labelWidth", value)
}

// MaxTagCount 标签的最大展示数量
func (nsc NestedSelectControl) MaxTagCount(value string) NestedSelectControl {
	return nsc.set("maxTagCount", value)
}

// MenuClassName 弹框的 css 类
func (nsc NestedSelectControl) MenuClassName(value string) NestedSelectControl {
	return nsc.set("menuClassName", value)
}

// Mode 配置当前表单项展示模式
func (nsc NestedSelectControl) Mode(value string) NestedSelectControl {
	return nsc.set("mode", value)
}

// Multiple 是否为多选模式
func (nsc NestedSelectControl) Multiple(value bool) NestedSelectControl {
	return nsc.set("multiple", value)
}

// Name 字段名
func (nsc NestedSelectControl) Name(value string) NestedSelectControl {
	return nsc.set("name", value)
}

// OnEvent 事件动作配置
func (nsc NestedSelectControl) OnEvent(value string) NestedSelectControl {
	return nsc.set("onEvent", value)
}

// OnlyChildren 选父级的时候是否只把子节点的值包含在内
func (nsc NestedSelectControl) OnlyChildren(value bool) NestedSelectControl {
	return nsc.set("onlyChildren", value)
}

// Options 选项
func (nsc NestedSelectControl) Options(value interface{}) NestedSelectControl {
	return nsc.set("options", value)
}

// OptionTpl 选项模板
func (nsc NestedSelectControl) OptionTpl(value string) NestedSelectControl {
	return nsc.set("optionTpl", value)
}

// Placeholder 占位提示
func (nsc NestedSelectControl) Placeholder(value string) NestedSelectControl {
	return nsc.set("placeholder", value)
}

// Remark 备注
func (nsc NestedSelectControl) Remark(value string) NestedSelectControl {
	return nsc.set("remark", value)
}

// Render 模板渲染方式
func (nsc NestedSelectControl) Render(value string) NestedSelectControl {
	return nsc.set("render", value)
}

// SearchAble 是否可搜索
func (nsc NestedSelectControl) SearchAble(value bool) NestedSelectControl {
	return nsc.set("searchAble", value)
}

// ShowDesc 显示描述
func (nsc NestedSelectControl) ShowDesc(value bool) NestedSelectControl {
	return nsc.set("showDesc", value)
}

// Size 大小
func (nsc NestedSelectControl) Size(value string) NestedSelectControl {
	return nsc.set("size", value)
}

// Static 动态属性
func (nsc NestedSelectControl) Static(value bool) NestedSelectControl {
	return nsc.set("static", value)
}

// StaticLabel 静态展示内容
func (nsc NestedSelectControl) StaticLabel(value string) NestedSelectControl {
	return nsc.set("staticLabel", value)
}

// Style 控件样式
func (nsc NestedSelectControl) Style(value string) NestedSelectControl {
	return nsc.set("style", value)
}

// Tips 提示文字
func (nsc NestedSelectControl) Tips(value string) NestedSelectControl {
	return nsc.set("tips", value)
}

// Tree 子节点是否为树形结构
func (nsc NestedSelectControl) Tree(value bool) NestedSelectControl {
	return nsc.set("tree", value)
}

// Value 选中值
func (nsc NestedSelectControl) Value(value interface{}) NestedSelectControl {
	return nsc.set("value", value)
}

// ValueField 表单项值字段
func (nsc NestedSelectControl) ValueField(value string) NestedSelectControl {
	return nsc.set("valueField", value)
}

// Visible 控件是否可见
func (nsc NestedSelectControl) Visible(value bool) NestedSelectControl {
	return nsc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (nsc NestedSelectControl) VisibleOn(value string) NestedSelectControl {
	return nsc.set("visibleOn", value)
}

// WithChildren 是否包含子节点
func (nsc NestedSelectControl) WithChildren(value bool) NestedSelectControl {
	return nsc.set("withChildren", value)
}

// WithLabelChildren 显示标签的子节点
func (nsc NestedSelectControl) WithLabelChildren(value bool) NestedSelectControl {
	return nsc.set("withLabelChildren", value)
}

// WithSelect 用于选择节点是否选中
func (nsc NestedSelectControl) WithSelect(value bool) NestedSelectControl {
	return nsc.set("withSelect", value)
}
