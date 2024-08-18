package comp

// breadcrumb 面包屑
type breadcrumb schema

// Breadcrumb 创建一个新的 Breadcrumb 实例
func Breadcrumb() breadcrumb {
	return make(breadcrumb)
}

func (b breadcrumb) set(key string, value interface{}) breadcrumb {
	b[key] = value
	return b
}

// ClassName 外层类名
func (b breadcrumb) ClassName(value string) breadcrumb {
	return b.set("className", value)
}

// DropdownClassName 下拉菜单类名
func (b breadcrumb) DropdownClassName(value string) breadcrumb {
	return b.set("dropdownClassName", value)
}

// DropdownItemClassName 下拉菜单项类名
func (b breadcrumb) DropdownItemClassName(value string) breadcrumb {
	return b.set("dropdownItemClassName", value)
}

// ItemClassName 导航项类名
func (b breadcrumb) ItemClassName(value string) breadcrumb {
	return b.set("itemClassName", value)
}

// Items 文本
func (b breadcrumb) Items(value string) breadcrumb {
	return b.set("items", value)
}

// LabelMaxLength 最大展示长度
func (b breadcrumb) LabelMaxLength(value string) breadcrumb {
	return b.set("labelMaxLength", value)
}

// Separator 分隔符
func (b breadcrumb) Separator(value string) breadcrumb {
	return b.set("separator", value)
}

// SeparatorClassName 分割符类名
func (b breadcrumb) SeparatorClassName(value string) breadcrumb {
	return b.set("separatorClassName", value)
}

// Source 动态数据
func (b breadcrumb) Source(value string) breadcrumb {
	return b.set("source", value)
}

// TooltipPosition 浮窗提示位置
func (b breadcrumb) TooltipPosition(value string) breadcrumb {
	return b.set("tooltipPosition", value)
}
