package comp

// avatar 头像渲染器
type avatar schema

// Avatar 创建一个新的 Avatar 实例
func Avatar() avatar {
	return make(avatar).set("type", "avatar").set("crossOrigin", "anonymous")
}

func (a avatar) set(key string, value interface{}) avatar {
	a[key] = value
	return a
}

// Alt 设置图片无法显示时的替换文字地址
func (a avatar) Alt(value string) avatar {
	return a.set("alt", value)
}

// Badge 设置角标
func (a avatar) Badge(value string) avatar {
	return a.set("badge", value)
}

// ClassName 设置类名
func (a avatar) ClassName(value string) avatar {
	return a.set("className", value)
}

// CrossOrigin 设置图片CORS属性
func (a avatar) CrossOrigin(value string) avatar {
	return a.set("crossOrigin", value)
}

// DefaultAvatar 设置默认头像
func (a avatar) DefaultAvatar(value string) avatar {
	return a.set("defaultAvatar", value)
}

// Disabled 设置是否禁用
func (a avatar) Disabled(value bool) avatar {
	return a.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (a avatar) DisabledOn(value string) avatar {
	return a.set("disabledOn", value)
}

// Draggable 设置图片是否允许拖动
func (a avatar) Draggable(value bool) avatar {
	return a.set("draggable", value)
}

// EditorSetting 设置编辑器配置
func (a avatar) EditorSetting(value string) avatar {
	return a.set("editorSetting", value)
}

// Fit 设置图片相对于容器的缩放方式
func (a avatar) Fit(value string) avatar {
	return a.set("fit", value)
}

// Gap 设置字符类型距离左右两侧边界单位像素
func (a avatar) Gap(value string) avatar {
	return a.set("gap", value)
}

// Hidden 设置是否隐藏
func (a avatar) Hidden(value bool) avatar {
	return a.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (a avatar) HiddenOn(value string) avatar {
	return a.set("hiddenOn", value)
}

// Icon 设置图标
func (a avatar) Icon(value string) avatar {
	return a.set("icon", value)
}

// ID 设置组件唯一 id
func (a avatar) ID(value string) avatar {
	return a.set("id", value)
}

// OnError 设置图片加载失败的处理
func (a avatar) OnError(value string) avatar {
	return a.set("onError", value)
}

// OnEvent 设置事件动作配置
func (a avatar) OnEvent(value string) avatar {
	return a.set("onEvent", value)
}

// Shape 设置形状
func (a avatar) Shape(value string) avatar {
	return a.set("shape", value)
}

// Size 设置大小
func (a avatar) Size(value string) avatar {
	return a.set("size", value)
}

// Src 设置图片地址
func (a avatar) Src(value string) avatar {
	return a.set("src", value)
}

// Static 设置是否静态展示
func (a avatar) Static(value bool) avatar {
	return a.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (a avatar) StaticClassName(value string) avatar {
	return a.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (a avatar) StaticInputClassName(value string) avatar {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (a avatar) StaticLabelClassName(value string) avatar {
	return a.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (a avatar) StaticOn(value string) avatar {
	return a.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (a avatar) StaticPlaceholder(value string) avatar {
	return a.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (a avatar) StaticSchema(value string) avatar {
	return a.set("staticSchema", value)
}

// Style 设置自定义样式
func (a avatar) Style(value string) avatar {
	return a.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (a avatar) TestIdBuilder(value string) avatar {
	return a.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (a avatar) Testid(value string) avatar {
	return a.set("testid", value)
}

// Text 设置文本
func (a avatar) Text(value string) avatar {
	return a.set("text", value)
}

// UseMobileUI 设置是否使用移动端样式
func (a avatar) UseMobileUI(value bool) avatar {
	return a.set("useMobileUI", value)
}

// Visible 设置是否显示
func (a avatar) Visible(value bool) avatar {
	return a.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (a avatar) VisibleOn(value string) avatar {
	return a.set("visibleOn", value)
}
