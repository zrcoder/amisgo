package comp

// Avatar 头像渲染器
type Avatar Schema

// NewAvatar 创建一个新的 Avatar 实例
func NewAvatar() Avatar {
	return make(Avatar).set("type", "avatar").set("crossOrigin", "anonymous")
}

func (a Avatar) set(key string, value interface{}) Avatar {
	a[key] = value
	return a
}

// Alt 设置图片无法显示时的替换文字地址
func (a Avatar) Alt(value string) Avatar {
	return a.set("alt", value)
}

// Badge 设置角标
func (a Avatar) Badge(value string) Avatar {
	return a.set("badge", value)
}

// ClassName 设置类名
func (a Avatar) ClassName(value string) Avatar {
	return a.set("className", value)
}

// CrossOrigin 设置图片CORS属性
func (a Avatar) CrossOrigin(value string) Avatar {
	return a.set("crossOrigin", value)
}

// DefaultAvatar 设置默认头像
func (a Avatar) DefaultAvatar(value string) Avatar {
	return a.set("defaultAvatar", value)
}

// Disabled 设置是否禁用
func (a Avatar) Disabled(value bool) Avatar {
	return a.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (a Avatar) DisabledOn(value string) Avatar {
	return a.set("disabledOn", value)
}

// Draggable 设置图片是否允许拖动
func (a Avatar) Draggable(value bool) Avatar {
	return a.set("draggable", value)
}

// EditorSetting 设置编辑器配置
func (a Avatar) EditorSetting(value string) Avatar {
	return a.set("editorSetting", value)
}

// Fit 设置图片相对于容器的缩放方式
func (a Avatar) Fit(value string) Avatar {
	return a.set("fit", value)
}

// Gap 设置字符类型距离左右两侧边界单位像素
func (a Avatar) Gap(value string) Avatar {
	return a.set("gap", value)
}

// Hidden 设置是否隐藏
func (a Avatar) Hidden(value bool) Avatar {
	return a.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (a Avatar) HiddenOn(value string) Avatar {
	return a.set("hiddenOn", value)
}

// Icon 设置图标
func (a Avatar) Icon(value string) Avatar {
	return a.set("icon", value)
}

// ID 设置组件唯一 id
func (a Avatar) ID(value string) Avatar {
	return a.set("id", value)
}

// OnError 设置图片加载失败的处理
func (a Avatar) OnError(value string) Avatar {
	return a.set("onError", value)
}

// OnEvent 设置事件动作配置
func (a Avatar) OnEvent(value string) Avatar {
	return a.set("onEvent", value)
}

// Shape 设置形状
func (a Avatar) Shape(value string) Avatar {
	return a.set("shape", value)
}

// Size 设置大小
func (a Avatar) Size(value string) Avatar {
	return a.set("size", value)
}

// Src 设置图片地址
func (a Avatar) Src(value string) Avatar {
	return a.set("src", value)
}

// Static 设置是否静态展示
func (a Avatar) Static(value bool) Avatar {
	return a.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (a Avatar) StaticClassName(value string) Avatar {
	return a.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (a Avatar) StaticInputClassName(value string) Avatar {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (a Avatar) StaticLabelClassName(value string) Avatar {
	return a.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (a Avatar) StaticOn(value string) Avatar {
	return a.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (a Avatar) StaticPlaceholder(value string) Avatar {
	return a.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (a Avatar) StaticSchema(value string) Avatar {
	return a.set("staticSchema", value)
}

// Style 设置自定义样式
func (a Avatar) Style(value string) Avatar {
	return a.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (a Avatar) TestIdBuilder(value string) Avatar {
	return a.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (a Avatar) Testid(value string) Avatar {
	return a.set("testid", value)
}

// Text 设置文本
func (a Avatar) Text(value string) Avatar {
	return a.set("text", value)
}

// UseMobileUI 设置是否使用移动端样式
func (a Avatar) UseMobileUI(value bool) Avatar {
	return a.set("useMobileUI", value)
}

// Visible 设置是否显示
func (a Avatar) Visible(value bool) Avatar {
	return a.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (a Avatar) VisibleOn(value string) Avatar {
	return a.set("visibleOn", value)
}
