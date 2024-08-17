package comp

// Breadcrumb 面包屑
type Breadcrumb Schema

// NewBreadcrumb 创建一个新的 Breadcrumb 实例
func NewBreadcrumb() Breadcrumb {
	return make(Breadcrumb)
}

// Set 设置属性
func (b Breadcrumb) set(key string, value interface{}) Breadcrumb {
	b[key] = value
	return b
}

// ClassName 外层类名
func (b Breadcrumb) ClassName(value string) Breadcrumb {
	return b.set("className", value)
}

// DropdownClassName 下拉菜单类名
func (b Breadcrumb) DropdownClassName(value string) Breadcrumb {
	return b.set("dropdownClassName", value)
}

// DropdownItemClassName 下拉菜单项类名
func (b Breadcrumb) DropdownItemClassName(value string) Breadcrumb {
	return b.set("dropdownItemClassName", value)
}

// ItemClassName 导航项类名
func (b Breadcrumb) ItemClassName(value string) Breadcrumb {
	return b.set("itemClassName", value)
}

// Items 文本
func (b Breadcrumb) Items(value string) Breadcrumb {
	return b.set("items", value)
}

// LabelMaxLength 最大展示长度
func (b Breadcrumb) LabelMaxLength(value string) Breadcrumb {
	return b.set("labelMaxLength", value)
}

// Separator 分隔符
func (b Breadcrumb) Separator(value string) Breadcrumb {
	return b.set("separator", value)
}

// SeparatorClassName 分割符类名
func (b Breadcrumb) SeparatorClassName(value string) Breadcrumb {
	return b.set("separatorClassName", value)
}

// Source 动态数据
func (b Breadcrumb) Source(value string) Breadcrumb {
	return b.set("source", value)
}

// TooltipPosition 浮窗提示位置
func (b Breadcrumb) TooltipPosition(value string) Breadcrumb {
	return b.set("tooltipPosition", value)
}
