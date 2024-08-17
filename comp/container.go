package comp

// Container 代表一个容器渲染器，支持多种配置项
type Container BaseRenderer

// NewContainer 创建一个新的 Container 实例，并设置默认的 type
func NewContainer() Container {
	c := Container(make(BaseRenderer))
	c.set("type", "container")
	return c
}

func (c Container) set(key string, value interface{}) Container {
	c[key] = value
	return c
}

// Body 设置内容
func (c Container) Body(value ...interface{}) Container {
	return c.set("body", value)
}

// BodyClassName 设置 body 类名
func (c Container) BodyClassName(value string) Container {
	return c.set("bodyClassName", value)
}

// ClassName 设置容器 css 类名
func (c Container) ClassName(value string) Container {
	return c.set("className", value)
}

// Disabled 设置是否禁用
func (c Container) Disabled(value bool) Container {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c Container) DisabledOn(value string) Container {
	return c.set("disabledOn", value)
}

// Draggable 设置是否开启容器拖拽
func (c Container) Draggable(value bool) Container {
	return c.set("draggable", value)
}

// DraggableConfig 设置是否开启容器拖拽配置
func (c Container) DraggableConfig(value string) Container {
	return c.set("draggableConfig", value)
}

// EditorSetting 设置编辑器配置
func (c Container) EditorSetting(value string) Container {
	return c.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (c Container) Hidden(value bool) Container {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c Container) HiddenOn(value string) Container {
	return c.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (c Container) ID(value string) Container {
	return c.set("id", value)
}

// OnEvent 设置事件动作配置
func (c Container) OnEvent(value string) Container {
	return c.set("onEvent", value)
}

// Static 设置是否静态展示
func (c Container) Static(value bool) Container {
	return c.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (c Container) StaticClassName(value string) Container {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (c Container) StaticInputClassName(value string) Container {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (c Container) StaticLabelClassName(value string) Container {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (c Container) StaticOn(value string) Container {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c Container) StaticPlaceholder(value string) Container {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (c Container) StaticSchema(value string) Container {
	return c.set("staticSchema", value)
}

// Style 设置自定义样式
func (c Container) Style(value string) Container {
	return c.set("style", value)
}

// TestIdBuilder 设置测试 id 构建器
func (c Container) TestIdBuilder(value string) Container {
	return c.set("testIdBuilder", value)
}

// Testid 设置测试 id
func (c Container) Testid(value string) Container {
	return c.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c Container) UseMobileUI(value bool) Container {
	return c.set("useMobileUI", value)
}

// Visible 设置是否显示
func (c Container) Visible(value bool) Container {
	return c.set("visible", value)
}

// VisibleOn 设置是否显示表达式
func (c Container) VisibleOn(value string) Container {
	return c.set("visibleOn", value)
}

// WrapperBody 设置是否需要对 body 加一层 div 包裹
func (c Container) WrapperBody(value bool) Container {
	return c.set("wrapperBody", value)
}

// WrapperComponent 设置使用的标签
func (c Container) WrapperComponent(value string) Container {
	return c.set("wrapperComponent", value)
}
