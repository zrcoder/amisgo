package comp

// combo 组合输入框类型
type combo schema

// Combo 创建一个新的 Combo 实例
func Combo() combo {
	return make(combo).set("type", "combo")
}

func (cc combo) set(key string, value any) combo {
	cc[key] = value
	return cc
}

// AddButtonClassName 设置新增按钮CSS类名
func (cc combo) AddButtonClassName(value string) combo {
	return cc.set("addButtonClassName", value)
}

// AddButtonText 设置新增按钮文字
func (cc combo) AddButtonText(value string) combo {
	return cc.set("addButtonText", value)
}

// Addable 设置是否可新增
func (cc combo) Addable(value bool) combo {
	return cc.set("addable", value)
}

// Addattop 设置 Add at top
func (cc combo) Addattop(value bool) combo {
	return cc.set("addattop", value)
}

// AutoFill 设置自动填充
func (cc combo) AutoFill(value string) combo {
	return cc.set("autoFill", value)
}

// CanAccessSuperData 设置是否可以访问父级数据
func (cc combo) CanAccessSuperData(value bool) combo {
	return cc.set("canAccessSuperData", value)
}

// ClassName 设置容器 css 类名
func (cc combo) ClassName(value string) combo {
	return cc.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否在当前 Form 中删除掉该表单项值
func (cc combo) ClearValueOnHidden(value bool) combo {
	return cc.set("clearValueOnHidden", value)
}

// Conditions 设置符合某类条件后才渲染的schema
func (cc combo) Conditions(value string) combo {
	return cc.set("conditions", value)
}

// DeleteApi 设置删除时调用的api
func (cc combo) DeleteApi(value string) combo {
	return cc.set("deleteApi", value)
}

// DeleteConfirmText 设置确认删除时的提示
func (cc combo) DeleteConfirmText(value string) combo {
	return cc.set("deleteConfirmText", value)
}

// Delimiter 设置分隔符
func (cc combo) Delimiter(value string) combo {
	return cc.set("delimiter", value)
}

// Desc 设置描述
func (cc combo) Desc(value string) combo {
	return cc.set("desc", value)
}

// Description 设置描述内容
func (cc combo) Description(value string) combo {
	return cc.set("description", value)
}

// DescriptionClassName 设置描述上的 className
func (cc combo) DescriptionClassName(value string) combo {
	return cc.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (cc combo) Disabled(value bool) combo {
	return cc.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (cc combo) DisabledOn(value string) combo {
	return cc.set("disabledOn", value)
}

// Draggable 设置是否可拖拽排序
func (cc combo) Draggable(value bool) combo {
	return cc.set("draggable", value)
}

// DraggableTip 设置可拖拽排序的提示信息
func (cc combo) DraggableTip(value string) combo {
	return cc.set("draggableTip", value)
}

// EditorSetting 设置编辑器配置
func (cc combo) EditorSetting(value string) combo {
	return cc.set("editorSetting", value)
}

// ExtraName 设置额外的字段名
func (cc combo) ExtraName(value string) combo {
	return cc.set("extraName", value)
}

// Flat 设置是否将结果扁平化
func (cc combo) Flat(value bool) combo {
	return cc.set("flat", value)
}

// FormClassName 设置内部单组表单项的类名
func (cc combo) FormClassName(value string) combo {
	return cc.set("formClassName", value)
}

// Hidden 设置是否隐藏
func (cc combo) Hidden(value bool) combo {
	return cc.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (cc combo) HiddenOn(value string) combo {
	return cc.set("hiddenOn", value)
}

// Hint 设置输入提示
func (cc combo) Hint(value string) combo {
	return cc.set("hint", value)
}

// Horizontal 设置水平布局具体左右分配
func (cc combo) Horizontal(value string) combo {
	return cc.set("horizontal", value)
}

// Id 设置组件唯一 id
func (cc combo) Id(value string) combo {
	return cc.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (cc combo) InitAutoFill(value string) combo {
	return cc.set("initAutoFill", value)
}

// Inline 设置表单 control 是否为 inline 模式
func (cc combo) Inline(value bool) combo {
	return cc.set("inline", value)
}

// InputClassName 设置 input className
func (cc combo) InputClassName(value string) combo {
	return cc.set("inputClassName", value)
}

// Items 设置数组输入框的子项
func (cc combo) Items(value ...any) combo {
	return cc.set("items", value)
}

// JoinValues 设置当扁平化开启时是否用分隔符的形式发送给后端
func (cc combo) JoinValues(value bool) combo {
	return cc.set("joinValues", value)
}

// Label 设置描述标题
func (cc combo) Label(value string) combo {
	return cc.set("label", value)
}

// LabelAlign 设置描述标题对齐方式
func (cc combo) LabelAlign(value string) combo {
	return cc.set("labelAlign", value)
}

// LabelClassName 设置 label className
func (cc combo) LabelClassName(value string) combo {
	return cc.set("labelClassName", value)
}

// LabelRemark 设置 label 旁的小图标提示
func (cc combo) LabelRemark(value string) combo {
	return cc.set("labelRemark", value)
}

// LabelWidth 设置 label 自定义宽度
func (cc combo) LabelWidth(value string) combo {
	return cc.set("labelWidth", value)
}

// LazyLoad 设置是否开启懒加载
func (cc combo) LazyLoad(value bool) combo {
	return cc.set("lazyLoad", value)
}

// MaxLength 设置限制最大个数
func (cc combo) MaxLength(value string) combo {
	return cc.set("maxLength", value)
}

// Messages 设置提示信息
func (cc combo) Messages(value string) combo {
	return cc.set("messages", value)
}

// MinLength 设置限制最小个数
func (cc combo) MinLength(value string) combo {
	return cc.set("minLength", value)
}

// Mode 设置当前表单项展示模式
func (cc combo) Mode(value string) combo {
	return cc.set("mode", value)
}

// MultiLine 设置是否多行模式
func (cc combo) MultiLine(value bool) combo {
	return cc.set("multiLine", value)
}

// Multiple 设置是否可多选
func (cc combo) Multiple(value bool) combo {
	return cc.set("multiple", value)
}

// Name 设置字段名
func (cc combo) Name(value string) combo {
	return cc.set("name", value)
}

// NoBorder 设置是否含有边框
func (cc combo) NoBorder(value bool) combo {
	return cc.set("noBorder", value)
}

// Nullable 设置是否允许为空
func (cc combo) Nullable(value bool) combo {
	return cc.set("nullable", value)
}

// OnEvent 设置事件动作配置
func (cc combo) OnEvent(value any) combo {
	return cc.set("onEvent", value)
}

// Placeholder 设置没有成员时显示的内容
func (cc combo) Placeholder(value string) combo {
	return cc.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (cc combo) ReadOnly(value bool) combo {
	return cc.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (cc combo) ReadOnlyOn(value string) combo {
	return cc.set("readOnlyOn", value)
}

// Remark 设置显示的小图标提示
func (cc combo) Remark(value string) combo {
	return cc.set("remark", value)
}

// Removable 设置是否可删除
func (cc combo) Removable(value bool) combo {
	return cc.set("removable", value)
}

// Required 设置是否为必填
func (cc combo) Required(value bool) combo {
	return cc.set("required", value)
}

// Row 设置单组表单项的行
func (cc combo) Row(value string) combo {
	return cc.set("row", value)
}

// SaveAs 设置保存时的格式
func (cc combo) SaveAs(value string) combo {
	return cc.set("saveAs", value)
}

// Size 设置组件大小
func (cc combo) Size(value string) combo {
	return cc.set("size", value)
}

// Source 设置数据源接口地址
func (cc combo) Source(value string) combo {
	return cc.set("source", value)
}

// StaticClassName 设置静态组件的类名
func (cc combo) StaticClassName(value string) combo {
	return cc.set("staticClassName", value)
}

// Striped 设置是否有条纹
func (cc combo) Striped(value bool) combo {
	return cc.set("striped", value)
}

// SubmitOnChange 设置是否在修改时触发提交
func (cc combo) SubmitOnChange(value bool) combo {
	return cc.set("submitOnChange", value)
}

// Tip 设置提示信息
func (cc combo) Tip(value string) combo {
	return cc.set("tip", value)
}

// Value 设置值
func (cc combo) Value(value string) combo {
	return cc.set("value", value)
}

// Vertical 设置垂直布局具体上下分配
func (cc combo) Vertical(value string) combo {
	return cc.set("vertical", value)
}

// Visible 设置是否可见
func (cc combo) Visible(value bool) combo {
	return cc.set("visible", value)
}

// VisibleOn 设置是否可见条件
func (cc combo) VisibleOn(value string) combo {
	return cc.set("visibleOn", value)
}

// Width 设置组件宽度
func (cc combo) Width(value string) combo {
	return cc.set("width", value)
}

// Wrap 设置外层 div 的 css 类名
func (cc combo) Wrap(value string) combo {
	return cc.set("wrap", value)
}
