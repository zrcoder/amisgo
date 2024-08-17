package comp

// ComboControl 组合输入框类型
type ComboControl BaseRenderer

// NewComboControl 创建一个新的 ComboControl 实例
func NewComboControl() ComboControl {
	cc := ComboControl(make(BaseRenderer))
	return cc.set("type", "combo")
}

func (cc ComboControl) set(key string, value interface{}) ComboControl {
	cc[key] = value
	return cc
}

// AddButtonClassName 设置新增按钮CSS类名
func (cc ComboControl) AddButtonClassName(value string) ComboControl {
	return cc.set("addButtonClassName", value)
}

// AddButtonText 设置新增按钮文字
func (cc ComboControl) AddButtonText(value string) ComboControl {
	return cc.set("addButtonText", value)
}

// Addable 设置是否可新增
func (cc ComboControl) Addable(value bool) ComboControl {
	return cc.set("addable", value)
}

// Addattop 设置 Add at top
func (cc ComboControl) Addattop(value bool) ComboControl {
	return cc.set("addattop", value)
}

// AutoFill 设置自动填充
func (cc ComboControl) AutoFill(value string) ComboControl {
	return cc.set("autoFill", value)
}

// CanAccessSuperData 设置是否可以访问父级数据
func (cc ComboControl) CanAccessSuperData(value bool) ComboControl {
	return cc.set("canAccessSuperData", value)
}

// ClassName 设置容器 css 类名
func (cc ComboControl) ClassName(value string) ComboControl {
	return cc.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否在当前 Form 中删除掉该表单项值
func (cc ComboControl) ClearValueOnHidden(value bool) ComboControl {
	return cc.set("clearValueOnHidden", value)
}

// Conditions 设置符合某类条件后才渲染的schema
func (cc ComboControl) Conditions(value string) ComboControl {
	return cc.set("conditions", value)
}

// DeleteApi 设置删除时调用的api
func (cc ComboControl) DeleteApi(value string) ComboControl {
	return cc.set("deleteApi", value)
}

// DeleteConfirmText 设置确认删除时的提示
func (cc ComboControl) DeleteConfirmText(value string) ComboControl {
	return cc.set("deleteConfirmText", value)
}

// Delimiter 设置分隔符
func (cc ComboControl) Delimiter(value string) ComboControl {
	return cc.set("delimiter", value)
}

// Desc 设置描述
func (cc ComboControl) Desc(value string) ComboControl {
	return cc.set("desc", value)
}

// Description 设置描述内容
func (cc ComboControl) Description(value string) ComboControl {
	return cc.set("description", value)
}

// DescriptionClassName 设置描述上的 className
func (cc ComboControl) DescriptionClassName(value string) ComboControl {
	return cc.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (cc ComboControl) Disabled(value bool) ComboControl {
	return cc.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (cc ComboControl) DisabledOn(value string) ComboControl {
	return cc.set("disabledOn", value)
}

// Draggable 设置是否可拖拽排序
func (cc ComboControl) Draggable(value bool) ComboControl {
	return cc.set("draggable", value)
}

// DraggableTip 设置可拖拽排序的提示信息
func (cc ComboControl) DraggableTip(value string) ComboControl {
	return cc.set("draggableTip", value)
}

// EditorSetting 设置编辑器配置
func (cc ComboControl) EditorSetting(value string) ComboControl {
	return cc.set("editorSetting", value)
}

// ExtraName 设置额外的字段名
func (cc ComboControl) ExtraName(value string) ComboControl {
	return cc.set("extraName", value)
}

// Flat 设置是否将结果扁平化
func (cc ComboControl) Flat(value bool) ComboControl {
	return cc.set("flat", value)
}

// FormClassName 设置内部单组表单项的类名
func (cc ComboControl) FormClassName(value string) ComboControl {
	return cc.set("formClassName", value)
}

// Hidden 设置是否隐藏
func (cc ComboControl) Hidden(value bool) ComboControl {
	return cc.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (cc ComboControl) HiddenOn(value string) ComboControl {
	return cc.set("hiddenOn", value)
}

// Hint 设置输入提示
func (cc ComboControl) Hint(value string) ComboControl {
	return cc.set("hint", value)
}

// Horizontal 设置水平布局具体左右分配
func (cc ComboControl) Horizontal(value string) ComboControl {
	return cc.set("horizontal", value)
}

// Id 设置组件唯一 id
func (cc ComboControl) Id(value string) ComboControl {
	return cc.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (cc ComboControl) InitAutoFill(value string) ComboControl {
	return cc.set("initAutoFill", value)
}

// Inline 设置表单 control 是否为 inline 模式
func (cc ComboControl) Inline(value bool) ComboControl {
	return cc.set("inline", value)
}

// InputClassName 设置 input className
func (cc ComboControl) InputClassName(value string) ComboControl {
	return cc.set("inputClassName", value)
}

// Items 设置数组输入框的子项
func (cc ComboControl) Items(value string) ComboControl {
	return cc.set("items", value)
}

// JoinValues 设置当扁平化开启时是否用分隔符的形式发送给后端
func (cc ComboControl) JoinValues(value bool) ComboControl {
	return cc.set("joinValues", value)
}

// Label 设置描述标题
func (cc ComboControl) Label(value string) ComboControl {
	return cc.set("label", value)
}

// LabelAlign 设置描述标题对齐方式
func (cc ComboControl) LabelAlign(value string) ComboControl {
	return cc.set("labelAlign", value)
}

// LabelClassName 设置 label className
func (cc ComboControl) LabelClassName(value string) ComboControl {
	return cc.set("labelClassName", value)
}

// LabelRemark 设置 label 旁的小图标提示
func (cc ComboControl) LabelRemark(value string) ComboControl {
	return cc.set("labelRemark", value)
}

// LabelWidth 设置 label 自定义宽度
func (cc ComboControl) LabelWidth(value string) ComboControl {
	return cc.set("labelWidth", value)
}

// LazyLoad 设置是否开启懒加载
func (cc ComboControl) LazyLoad(value bool) ComboControl {
	return cc.set("lazyLoad", value)
}

// MaxLength 设置限制最大个数
func (cc ComboControl) MaxLength(value string) ComboControl {
	return cc.set("maxLength", value)
}

// Messages 设置提示信息
func (cc ComboControl) Messages(value string) ComboControl {
	return cc.set("messages", value)
}

// MinLength 设置限制最小个数
func (cc ComboControl) MinLength(value string) ComboControl {
	return cc.set("minLength", value)
}

// Mode 设置当前表单项展示模式
func (cc ComboControl) Mode(value string) ComboControl {
	return cc.set("mode", value)
}

// MultiLine 设置是否多行模式
func (cc ComboControl) MultiLine(value bool) ComboControl {
	return cc.set("multiLine", value)
}

// Multiple 设置是否可多选
func (cc ComboControl) Multiple(value bool) ComboControl {
	return cc.set("multiple", value)
}

// Name 设置字段名
func (cc ComboControl) Name(value string) ComboControl {
	return cc.set("name", value)
}

// NoBorder 设置是否含有边框
func (cc ComboControl) NoBorder(value bool) ComboControl {
	return cc.set("noBorder", value)
}

// Nullable 设置是否允许为空
func (cc ComboControl) Nullable(value bool) ComboControl {
	return cc.set("nullable", value)
}

// OnEvent 设置事件动作配置
func (cc ComboControl) OnEvent(value string) ComboControl {
	return cc.set("onEvent", value)
}

// Placeholder 设置没有成员时显示的内容
func (cc ComboControl) Placeholder(value string) ComboControl {
	return cc.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (cc ComboControl) ReadOnly(value bool) ComboControl {
	return cc.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (cc ComboControl) ReadOnlyOn(value string) ComboControl {
	return cc.set("readOnlyOn", value)
}

// Remark 设置显示的小图标提示
func (cc ComboControl) Remark(value string) ComboControl {
	return cc.set("remark", value)
}

// Removable 设置是否可删除
func (cc ComboControl) Removable(value bool) ComboControl {
	return cc.set("removable", value)
}

// Required 设置是否为必填
func (cc ComboControl) Required(value bool) ComboControl {
	return cc.set("required", value)
}

// Row 设置单组表单项的行
func (cc ComboControl) Row(value string) ComboControl {
	return cc.set("row", value)
}

// SaveAs 设置保存时的格式
func (cc ComboControl) SaveAs(value string) ComboControl {
	return cc.set("saveAs", value)
}

// Size 设置组件大小
func (cc ComboControl) Size(value string) ComboControl {
	return cc.set("size", value)
}

// Source 设置数据源接口地址
func (cc ComboControl) Source(value string) ComboControl {
	return cc.set("source", value)
}

// StaticClassName 设置静态组件的类名
func (cc ComboControl) StaticClassName(value string) ComboControl {
	return cc.set("staticClassName", value)
}

// Striped 设置是否有条纹
func (cc ComboControl) Striped(value bool) ComboControl {
	return cc.set("striped", value)
}

// SubmitOnChange 设置是否在修改时触发提交
func (cc ComboControl) SubmitOnChange(value bool) ComboControl {
	return cc.set("submitOnChange", value)
}

// Tip 设置提示信息
func (cc ComboControl) Tip(value string) ComboControl {
	return cc.set("tip", value)
}

// Value 设置值
func (cc ComboControl) Value(value string) ComboControl {
	return cc.set("value", value)
}

// Vertical 设置垂直布局具体上下分配
func (cc ComboControl) Vertical(value string) ComboControl {
	return cc.set("vertical", value)
}

// Visible 设置是否可见
func (cc ComboControl) Visible(value bool) ComboControl {
	return cc.set("visible", value)
}

// VisibleOn 设置是否可见条件
func (cc ComboControl) VisibleOn(value string) ComboControl {
	return cc.set("visibleOn", value)
}

// Width 设置组件宽度
func (cc ComboControl) Width(value string) ComboControl {
	return cc.set("width", value)
}

// Wrap 设置外层 div 的 css 类名
func (cc ComboControl) Wrap(value string) ComboControl {
	return cc.set("wrap", value)
}
