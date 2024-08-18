package comp

// TableControl 表格控件
//
// @version 6.7.0
type TableControl Schema

// NewTableControl 创建一个新的 TableControl 实例
func NewTableControl() TableControl {
	return TableControl{}.set("type", "input-table")
}

func (tc TableControl) set(key string, value interface{}) TableControl {
	tc[key] = value
	return tc
}

// AddApi 新增 API (新增 API)
func (tc TableControl) AddApi(value string) TableControl {
	return tc.set("addApi", value)
}

// AddBtnIcon 新增按钮图标
func (tc TableControl) AddBtnIcon(value string) TableControl {
	return tc.set("addBtnIcon", value)
}

// AddBtnLabel 新增按钮文字
func (tc TableControl) AddBtnLabel(value string) TableControl {
	return tc.set("addBtnLabel", value)
}

// Addable 可新增
func (tc TableControl) Addable(value bool) TableControl {
	return tc.set("addable", value)
}

// AffixFooter 是否固底
func (tc TableControl) AffixFooter(value bool) TableControl {
	return tc.set("affixFooter", value)
}

// AffixHeader 是否固定表头
func (tc TableControl) AffixHeader(value bool) TableControl {
	return tc.set("affixHeader", value)
}

// AffixRow 底部总结行
func (tc TableControl) AffixRow(value string) TableControl {
	return tc.set("affixRow", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tc TableControl) AutoFill(value string) TableControl {
	return tc.set("autoFill", value)
}

// AutoFillHeight 表格自动计算高度
func (tc TableControl) AutoFillHeight(value string) TableControl {
	return tc.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
func (tc TableControl) AutoGenerateFilter(value string) TableControl {
	return tc.set("autoGenerateFilter", value)
}

// CanAccessSuperData 是否可以访问父级数据，正常 combo 已经关联到数组成员，是不能访问父级数据的。
func (tc TableControl) CanAccessSuperData(value bool) TableControl {
	return tc.set("canAccessSuperData", value)
}

// CancelBtnIcon 取消按钮图标
func (tc TableControl) CancelBtnIcon(value string) TableControl {
	return tc.set("cancelBtnIcon", value)
}

// CancelBtnLabel 取消按钮文字
func (tc TableControl) CancelBtnLabel(value string) TableControl {
	return tc.set("cancelBtnLabel", value)
}

// ChildrenAddable 是否可以新增子项
func (tc TableControl) ChildrenAddable(value bool) TableControl {
	return tc.set("childrenAddable", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tc TableControl) ClassName(value string) TableControl {
	return tc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tc TableControl) ClearValueOnHidden(value bool) TableControl {
	return tc.set("clearValueOnHidden", value)
}

// Columns 表格的列信息
func (tc TableControl) Columns(value string) TableControl {
	return tc.set("columns", value)
}

// ColumnsTogglable 展示列显示开关，自动即：列数量大于或等于5个时自动开启
func (tc TableControl) ColumnsTogglable(value bool) TableControl {
	return tc.set("columnsTogglable", value)
}

// CombineFromIndex 合并单元格配置，配置从第几列开始合并。
func (tc TableControl) CombineFromIndex(value string) TableControl {
	return tc.set("combineFromIndex", value)
}

// CombineNum 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格。
func (tc TableControl) CombineNum(value string) TableControl {
	return tc.set("combineNum", value)
}

// ConfirmBtnIcon 确认按钮图标
func (tc TableControl) ConfirmBtnIcon(value string) TableControl {
	return tc.set("confirmBtnIcon", value)
}

// ConfirmBtnLabel 确认按钮文字
func (tc TableControl) ConfirmBtnLabel(value string) TableControl {
	return tc.set("confirmBtnLabel", value)
}

// CopyAddBtn 是否显示复制按钮
func (tc TableControl) CopyAddBtn(value bool) TableControl {
	return tc.set("copyAddBtn", value)
}

// CopyBtnIcon 复制按钮图标
func (tc TableControl) CopyBtnIcon(value string) TableControl {
	return tc.set("copyBtnIcon", value)
}

// CopyBtnLabel 复制按钮文字
func (tc TableControl) CopyBtnLabel(value string) TableControl {
	return tc.set("copyBtnLabel", value)
}

// CopyData 复制的时候用来配置复制映射的数据。默认值是 {&:$$}，相当与复制整个行数据 通常有时候需要用来标记是复制过来的，也可能需要删掉一下主键字段。
func (tc TableControl) CopyData(value string) TableControl {
	return tc.set("copyData", value)
}

// Copyable 可复制新增
func (tc TableControl) Copyable(value bool) TableControl {
	return tc.set("copyable", value)
}

// DeferApi 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。 (懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。)
func (tc TableControl) DeferApi(value string) TableControl {
	return tc.set("deferApi", value)
}

// DeleteApi 删除的 API (删除的 API)
func (tc TableControl) DeleteApi(value string) TableControl {
	return tc.set("deleteApi", value)
}

// DeleteBtnIcon 删除按钮图标
func (tc TableControl) DeleteBtnIcon(value string) TableControl {
	return tc.set("deleteBtnIcon", value)
}

// DeleteBtnLabel 删除按钮文字
func (tc TableControl) DeleteBtnLabel(value string) TableControl {
	return tc.set("deleteBtnLabel", value)
}

// DeleteConfirmText 删除确认文字
func (tc TableControl) DeleteConfirmText(value string) TableControl {
	return tc.set("deleteConfirmText", value)
}

// Desc
func (tc TableControl) Desc(value string) TableControl {
	return tc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tc TableControl) Description(value string) TableControl {
	return tc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tc TableControl) DescriptionClassName(value string) TableControl {
	return tc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tc TableControl) Disabled(value bool) TableControl {
	return tc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TableControl) DisabledOn(value string) TableControl {
	return tc.set("disabledOn", value)
}

// Draggable 是否可以拖拽排序
func (tc TableControl) Draggable(value bool) TableControl {
	return tc.set("draggable", value)
}

// EditBtnIcon 更新按钮图标
func (tc TableControl) EditBtnIcon(value string) TableControl {
	return tc.set("editBtnIcon", value)
}

// EditBtnLabel 更新按钮名称
func (tc TableControl) EditBtnLabel(value string) TableControl {
	return tc.set("editBtnLabel", value)
}

// Editable 可否编辑
func (tc TableControl) Editable(value bool) TableControl {
	return tc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tc TableControl) EditorSetting(value string) TableControl {
	return tc.set("editorSetting", value)
}

// EnableStaticTransform 是否开启 static 状态切换
func (tc TableControl) EnableStaticTransform(value bool) TableControl {
	return tc.set("enableStaticTransform", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tc TableControl) ExtraName(value string) TableControl {
	return tc.set("extraName", value)
}

// Footable 是否开启底部展示功能，适合移动端展示
func (tc TableControl) Footable(value bool) TableControl {
	return tc.set("footable", value)
}

// FooterAddBtn 底部新增按钮配置 (底部新增按钮配置)
func (tc TableControl) FooterAddBtn(value string) TableControl {
	return tc.set("footerAddBtn", value)
}

// FooterClassName 底部外层 CSS 类名 (底部外层 CSS 类名)
func (tc TableControl) FooterClassName(value string) TableControl {
	return tc.set("footerClassName", value)
}

// FooterTitle 底部标题
func (tc TableControl) FooterTitle(value string) TableControl {
	return tc.set("footerTitle", value)
}

// FormApi 表单 API (表单 API)
func (tc TableControl) FormApi(value string) TableControl {
	return tc.set("formApi", value)
}

// HideQuickSave 隐藏快捷保存按钮
func (tc TableControl) HideQuickSave(value bool) TableControl {
	return tc.set("hideQuickSave", value)
}

// HoverAction 悬浮操作的配置
func (tc TableControl) HoverAction(value string) TableControl {
	return tc.set("hoverAction", value)
}

// HoverTip 触发时的提示文字
func (tc TableControl) HoverTip(value string) TableControl {
	return tc.set("hoverTip", value)
}

// Inline 编辑框是否行内展示，默认是 false，即展示成弹框。
func (tc TableControl) Inline(value bool) TableControl {
	return tc.set("inline", value)
}

// IsPreview 是否预览模式
func (tc TableControl) IsPreview(value bool) TableControl {
	return tc.set("isPreview", value)
}

// ItemActionItem (列表项操作项的配置)
func (tc TableControl) ItemActionItem(value string) TableControl {
	return tc.set("itemActionItem", value)
}

// ItemActions 相关操作项 (相关操作项)
func (tc TableControl) ItemActions(value string) TableControl {
	return tc.set("itemActions", value)
}

// Label 标签文字
func (tc TableControl) Label(value string) TableControl {
	return tc.set("label", value)
}

// LabelClassName 标签的类名
func (tc TableControl) LabelClassName(value string) TableControl {
	return tc.set("labelClassName", value)
}

// LabelRemark 标签文字的备注
func (tc TableControl) LabelRemark(value string) TableControl {
	return tc.set("labelRemark", value)
}

// MaxLength 最大长度
func (tc TableControl) MaxLength(value string) TableControl {
	return tc.set("maxLength", value)
}

// Merge
func (tc TableControl) Merge(value string) TableControl {
	return tc.set("merge", value)
}

// MinLength 最小长度
func (tc TableControl) MinLength(value string) TableControl {
	return tc.set("minLength", value)
}

// MultiLine 多行显示
func (tc TableControl) MultiLine(value bool) TableControl {
	return tc.set("multiLine", value)
}

// Name 配置字段名称，作为数据提交的 key
func (tc TableControl) Name(value string) TableControl {
	return tc.set("name", value)
}

// NoEmpty 忽略未填字段 (忽略未填字段)
func (tc TableControl) NoEmpty(value bool) TableControl {
	return tc.set("noEmpty", value)
}

// NoBorder 组件外层是否需要 border
func (tc TableControl) NoBorder(value bool) TableControl {
	return tc.set("noBorder", value)
}

// NoIndex 不显示序号列
func (tc TableControl) NoIndex(value bool) TableControl {
	return tc.set("noIndex", value)
}

// NoResults 没有数据时显示的内容
func (tc TableControl) NoResults(value string) TableControl {
	return tc.set("noResults", value)
}

// OnAction 触发后的回调函数
func (tc TableControl) OnAction(value string) TableControl {
	return tc.set("onAction", value)
}

// OnChange 组件值变化时的回调函数
func (tc TableControl) OnChange(value string) TableControl {
	return tc.set("onChange", value)
}

// OnConfirm 触发确认的回调
func (tc TableControl) OnConfirm(value string) TableControl {
	return tc.set("onConfirm", value)
}

// OnDelete 触发删除的回调
func (tc TableControl) OnDelete(value string) TableControl {
	return tc.set("onDelete", value)
}

// OnEvent 事件触发器配置
func (tc TableControl) OnEvent(value string) TableControl {
	return tc.set("onEvent", value)
}

// Placeholder 提示文字
func (tc TableControl) Placeholder(value string) TableControl {
	return tc.set("placeholder", value)
}

// QuickSave 快速保存时，是否弹出对话框
func (tc TableControl) QuickSave(value bool) TableControl {
	return tc.set("quickSave", value)
}

// ReadOnly 是否只读
func (tc TableControl) ReadOnly(value bool) TableControl {
	return tc.set("readOnly", value)
}

// Remarks 附加说明信息
func (tc TableControl) Remarks(value string) TableControl {
	return tc.set("remarks", value)
}

// Remote 模板的远程接口配置
func (tc TableControl) Remote(value string) TableControl {
	return tc.set("remote", value)
}

// RenderLabel 渲染标签
func (tc TableControl) RenderLabel(value string) TableControl {
	return tc.set("renderLabel", value)
}

// SaveApi 保存的 API (保存的 API)
func (tc TableControl) SaveApi(value string) TableControl {
	return tc.set("saveApi", value)
}

// SearchApi 查询的 API (查询的 API)
func (tc TableControl) SearchApi(value string) TableControl {
	return tc.set("searchApi", value)
}

// Searchable 可查询
func (tc TableControl) Searchable(value bool) TableControl {
	return tc.set("searchable", value)
}

// Size 尺寸 (可配置为：large, medium, small)
func (tc TableControl) Size(value string) TableControl {
	return tc.set("size", value)
}

// Source 远程接口配置
func (tc TableControl) Source(value string) TableControl {
	return tc.set("source", value)
}

// Static 静态数据配置 (静态数据配置)
func (tc TableControl) Static(value string) TableControl {
	return tc.set("static", value)
}

// Status 状态
func (tc TableControl) Status(value string) TableControl {
	return tc.set("status", value)
}

// SubForm 关联表单
func (tc TableControl) SubForm(value string) TableControl {
	return tc.set("subForm", value)
}

// SubmitOnChange 是否提交更改 (是否提交更改)
func (tc TableControl) SubmitOnChange(value bool) TableControl {
	return tc.set("submitOnChange", value)
}

// Tag 标签配置 (标签配置)
func (tc TableControl) Tag(value string) TableControl {
	return tc.set("tag", value)
}

// Union 合并表单 (合并表单)
func (tc TableControl) Union(value string) TableControl {
	return tc.set("union", value)
}

// Validate 是否校验
func (tc TableControl) Validate(value bool) TableControl {
	return tc.set("validate", value)
}

// Value 默认值
func (tc TableControl) Value(value string) TableControl {
	return tc.set("value", value)
}

// ValueType 值的类型
func (tc TableControl) ValueType(value string) TableControl {
	return tc.set("valueType", value)
}

// Width 控件宽度
func (tc TableControl) Width(value string) TableControl {
	return tc.set("width", value)
}
