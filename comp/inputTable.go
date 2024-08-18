package comp

// inputTable 表格控件
//
// @version 6.7.0
type inputTable schema

// InputTable 创建一个新的 TableControl 实例
func InputTable() inputTable {
	return inputTable{}.set("type", "input-table")
}

func (tc inputTable) set(key string, value interface{}) inputTable {
	tc[key] = value
	return tc
}

// AddApi 新增 API (新增 API)
func (tc inputTable) AddApi(value string) inputTable {
	return tc.set("addApi", value)
}

// AddBtnIcon 新增按钮图标
func (tc inputTable) AddBtnIcon(value string) inputTable {
	return tc.set("addBtnIcon", value)
}

// AddBtnLabel 新增按钮文字
func (tc inputTable) AddBtnLabel(value string) inputTable {
	return tc.set("addBtnLabel", value)
}

// Addable 可新增
func (tc inputTable) Addable(value bool) inputTable {
	return tc.set("addable", value)
}

// AffixFooter 是否固底
func (tc inputTable) AffixFooter(value bool) inputTable {
	return tc.set("affixFooter", value)
}

// AffixHeader 是否固定表头
func (tc inputTable) AffixHeader(value bool) inputTable {
	return tc.set("affixHeader", value)
}

// AffixRow 底部总结行
func (tc inputTable) AffixRow(value string) inputTable {
	return tc.set("affixRow", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tc inputTable) AutoFill(value string) inputTable {
	return tc.set("autoFill", value)
}

// AutoFillHeight 表格自动计算高度
func (tc inputTable) AutoFillHeight(value string) inputTable {
	return tc.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
func (tc inputTable) AutoGenerateFilter(value string) inputTable {
	return tc.set("autoGenerateFilter", value)
}

// CanAccessSuperData 是否可以访问父级数据，正常 combo 已经关联到数组成员，是不能访问父级数据的。
func (tc inputTable) CanAccessSuperData(value bool) inputTable {
	return tc.set("canAccessSuperData", value)
}

// CancelBtnIcon 取消按钮图标
func (tc inputTable) CancelBtnIcon(value string) inputTable {
	return tc.set("cancelBtnIcon", value)
}

// CancelBtnLabel 取消按钮文字
func (tc inputTable) CancelBtnLabel(value string) inputTable {
	return tc.set("cancelBtnLabel", value)
}

// ChildrenAddable 是否可以新增子项
func (tc inputTable) ChildrenAddable(value bool) inputTable {
	return tc.set("childrenAddable", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tc inputTable) ClassName(value string) inputTable {
	return tc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tc inputTable) ClearValueOnHidden(value bool) inputTable {
	return tc.set("clearValueOnHidden", value)
}

// Columns 表格的列信息
func (tc inputTable) Columns(value string) inputTable {
	return tc.set("columns", value)
}

// ColumnsTogglable 展示列显示开关，自动即：列数量大于或等于5个时自动开启
func (tc inputTable) ColumnsTogglable(value bool) inputTable {
	return tc.set("columnsTogglable", value)
}

// CombineFromIndex 合并单元格配置，配置从第几列开始合并。
func (tc inputTable) CombineFromIndex(value string) inputTable {
	return tc.set("combineFromIndex", value)
}

// CombineNum 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格。
func (tc inputTable) CombineNum(value string) inputTable {
	return tc.set("combineNum", value)
}

// ConfirmBtnIcon 确认按钮图标
func (tc inputTable) ConfirmBtnIcon(value string) inputTable {
	return tc.set("confirmBtnIcon", value)
}

// ConfirmBtnLabel 确认按钮文字
func (tc inputTable) ConfirmBtnLabel(value string) inputTable {
	return tc.set("confirmBtnLabel", value)
}

// CopyAddBtn 是否显示复制按钮
func (tc inputTable) CopyAddBtn(value bool) inputTable {
	return tc.set("copyAddBtn", value)
}

// CopyBtnIcon 复制按钮图标
func (tc inputTable) CopyBtnIcon(value string) inputTable {
	return tc.set("copyBtnIcon", value)
}

// CopyBtnLabel 复制按钮文字
func (tc inputTable) CopyBtnLabel(value string) inputTable {
	return tc.set("copyBtnLabel", value)
}

// CopyData 复制的时候用来配置复制映射的数据。默认值是 {&:$$}，相当与复制整个行数据 通常有时候需要用来标记是复制过来的，也可能需要删掉一下主键字段。
func (tc inputTable) CopyData(value string) inputTable {
	return tc.set("copyData", value)
}

// Copyable 可复制新增
func (tc inputTable) Copyable(value bool) inputTable {
	return tc.set("copyable", value)
}

// DeferApi 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。 (懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。)
func (tc inputTable) DeferApi(value string) inputTable {
	return tc.set("deferApi", value)
}

// DeleteApi 删除的 API (删除的 API)
func (tc inputTable) DeleteApi(value string) inputTable {
	return tc.set("deleteApi", value)
}

// DeleteBtnIcon 删除按钮图标
func (tc inputTable) DeleteBtnIcon(value string) inputTable {
	return tc.set("deleteBtnIcon", value)
}

// DeleteBtnLabel 删除按钮文字
func (tc inputTable) DeleteBtnLabel(value string) inputTable {
	return tc.set("deleteBtnLabel", value)
}

// DeleteConfirmText 删除确认文字
func (tc inputTable) DeleteConfirmText(value string) inputTable {
	return tc.set("deleteConfirmText", value)
}

// Desc
func (tc inputTable) Desc(value string) inputTable {
	return tc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tc inputTable) Description(value string) inputTable {
	return tc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tc inputTable) DescriptionClassName(value string) inputTable {
	return tc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tc inputTable) Disabled(value bool) inputTable {
	return tc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tc inputTable) DisabledOn(value string) inputTable {
	return tc.set("disabledOn", value)
}

// Draggable 是否可以拖拽排序
func (tc inputTable) Draggable(value bool) inputTable {
	return tc.set("draggable", value)
}

// EditBtnIcon 更新按钮图标
func (tc inputTable) EditBtnIcon(value string) inputTable {
	return tc.set("editBtnIcon", value)
}

// EditBtnLabel 更新按钮名称
func (tc inputTable) EditBtnLabel(value string) inputTable {
	return tc.set("editBtnLabel", value)
}

// Editable 可否编辑
func (tc inputTable) Editable(value bool) inputTable {
	return tc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tc inputTable) EditorSetting(value string) inputTable {
	return tc.set("editorSetting", value)
}

// EnableStaticTransform 是否开启 static 状态切换
func (tc inputTable) EnableStaticTransform(value bool) inputTable {
	return tc.set("enableStaticTransform", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tc inputTable) ExtraName(value string) inputTable {
	return tc.set("extraName", value)
}

// Footable 是否开启底部展示功能，适合移动端展示
func (tc inputTable) Footable(value bool) inputTable {
	return tc.set("footable", value)
}

// FooterAddBtn 底部新增按钮配置 (底部新增按钮配置)
func (tc inputTable) FooterAddBtn(value string) inputTable {
	return tc.set("footerAddBtn", value)
}

// FooterClassName 底部外层 CSS 类名 (底部外层 CSS 类名)
func (tc inputTable) FooterClassName(value string) inputTable {
	return tc.set("footerClassName", value)
}

// FooterTitle 底部标题
func (tc inputTable) FooterTitle(value string) inputTable {
	return tc.set("footerTitle", value)
}

// FormApi 表单 API (表单 API)
func (tc inputTable) FormApi(value string) inputTable {
	return tc.set("formApi", value)
}

// HideQuickSave 隐藏快捷保存按钮
func (tc inputTable) HideQuickSave(value bool) inputTable {
	return tc.set("hideQuickSave", value)
}

// HoverAction 悬浮操作的配置
func (tc inputTable) HoverAction(value string) inputTable {
	return tc.set("hoverAction", value)
}

// HoverTip 触发时的提示文字
func (tc inputTable) HoverTip(value string) inputTable {
	return tc.set("hoverTip", value)
}

// Inline 编辑框是否行内展示，默认是 false，即展示成弹框。
func (tc inputTable) Inline(value bool) inputTable {
	return tc.set("inline", value)
}

// IsPreview 是否预览模式
func (tc inputTable) IsPreview(value bool) inputTable {
	return tc.set("isPreview", value)
}

// ItemActionItem (列表项操作项的配置)
func (tc inputTable) ItemActionItem(value string) inputTable {
	return tc.set("itemActionItem", value)
}

// ItemActions 相关操作项 (相关操作项)
func (tc inputTable) ItemActions(value string) inputTable {
	return tc.set("itemActions", value)
}

// Label 标签文字
func (tc inputTable) Label(value string) inputTable {
	return tc.set("label", value)
}

// LabelClassName 标签的类名
func (tc inputTable) LabelClassName(value string) inputTable {
	return tc.set("labelClassName", value)
}

// LabelRemark 标签文字的备注
func (tc inputTable) LabelRemark(value string) inputTable {
	return tc.set("labelRemark", value)
}

// MaxLength 最大长度
func (tc inputTable) MaxLength(value string) inputTable {
	return tc.set("maxLength", value)
}

// Merge
func (tc inputTable) Merge(value string) inputTable {
	return tc.set("merge", value)
}

// MinLength 最小长度
func (tc inputTable) MinLength(value string) inputTable {
	return tc.set("minLength", value)
}

// MultiLine 多行显示
func (tc inputTable) MultiLine(value bool) inputTable {
	return tc.set("multiLine", value)
}

// Name 配置字段名称，作为数据提交的 key
func (tc inputTable) Name(value string) inputTable {
	return tc.set("name", value)
}

// NoEmpty 忽略未填字段 (忽略未填字段)
func (tc inputTable) NoEmpty(value bool) inputTable {
	return tc.set("noEmpty", value)
}

// NoBorder 组件外层是否需要 border
func (tc inputTable) NoBorder(value bool) inputTable {
	return tc.set("noBorder", value)
}

// NoIndex 不显示序号列
func (tc inputTable) NoIndex(value bool) inputTable {
	return tc.set("noIndex", value)
}

// NoResults 没有数据时显示的内容
func (tc inputTable) NoResults(value string) inputTable {
	return tc.set("noResults", value)
}

// OnAction 触发后的回调函数
func (tc inputTable) OnAction(value string) inputTable {
	return tc.set("onAction", value)
}

// OnChange 组件值变化时的回调函数
func (tc inputTable) OnChange(value string) inputTable {
	return tc.set("onChange", value)
}

// OnConfirm 触发确认的回调
func (tc inputTable) OnConfirm(value string) inputTable {
	return tc.set("onConfirm", value)
}

// OnDelete 触发删除的回调
func (tc inputTable) OnDelete(value string) inputTable {
	return tc.set("onDelete", value)
}

// OnEvent 事件触发器配置
func (tc inputTable) OnEvent(value string) inputTable {
	return tc.set("onEvent", value)
}

// Placeholder 提示文字
func (tc inputTable) Placeholder(value string) inputTable {
	return tc.set("placeholder", value)
}

// QuickSave 快速保存时，是否弹出对话框
func (tc inputTable) QuickSave(value bool) inputTable {
	return tc.set("quickSave", value)
}

// ReadOnly 是否只读
func (tc inputTable) ReadOnly(value bool) inputTable {
	return tc.set("readOnly", value)
}

// Remarks 附加说明信息
func (tc inputTable) Remarks(value string) inputTable {
	return tc.set("remarks", value)
}

// Remote 模板的远程接口配置
func (tc inputTable) Remote(value string) inputTable {
	return tc.set("remote", value)
}

// RenderLabel 渲染标签
func (tc inputTable) RenderLabel(value string) inputTable {
	return tc.set("renderLabel", value)
}

// SaveApi 保存的 API (保存的 API)
func (tc inputTable) SaveApi(value string) inputTable {
	return tc.set("saveApi", value)
}

// SearchApi 查询的 API (查询的 API)
func (tc inputTable) SearchApi(value string) inputTable {
	return tc.set("searchApi", value)
}

// Searchable 可查询
func (tc inputTable) Searchable(value bool) inputTable {
	return tc.set("searchable", value)
}

// Size 尺寸 (可配置为：large, medium, small)
func (tc inputTable) Size(value string) inputTable {
	return tc.set("size", value)
}

// Source 远程接口配置
func (tc inputTable) Source(value string) inputTable {
	return tc.set("source", value)
}

// Static 静态数据配置 (静态数据配置)
func (tc inputTable) Static(value string) inputTable {
	return tc.set("static", value)
}

// Status 状态
func (tc inputTable) Status(value string) inputTable {
	return tc.set("status", value)
}

// SubForm 关联表单
func (tc inputTable) SubForm(value string) inputTable {
	return tc.set("subForm", value)
}

// SubmitOnChange 是否提交更改 (是否提交更改)
func (tc inputTable) SubmitOnChange(value bool) inputTable {
	return tc.set("submitOnChange", value)
}

// Tag 标签配置 (标签配置)
func (tc inputTable) Tag(value string) inputTable {
	return tc.set("tag", value)
}

// Union 合并表单 (合并表单)
func (tc inputTable) Union(value string) inputTable {
	return tc.set("union", value)
}

// Validate 是否校验
func (tc inputTable) Validate(value bool) inputTable {
	return tc.set("validate", value)
}

// Value 默认值
func (tc inputTable) Value(value string) inputTable {
	return tc.set("value", value)
}

// ValueType 值的类型
func (tc inputTable) ValueType(value string) inputTable {
	return tc.set("valueType", value)
}

// Width 控件宽度
func (tc inputTable) Width(value string) inputTable {
	return tc.set("width", value)
}
