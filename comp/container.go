package comp

// container 代表一个容器渲染器，支持多种配置项
type container schema

// Container 创建一个新的 Container 实例，并设置默认的 type
func Container() container {
	return make(container).set("type", "container")
}

func (c container) set(key string, value interface{}) container {
	c[key] = value
	return c
}

// Body 设置内容
func (c container) Body(value ...interface{}) container {
	return c.set("body", value)
}

// BodyClassName 设置 body 类名
func (c container) BodyClassName(value string) container {
	return c.set("bodyClassName", value)
}

// ClassName 设置容器 css 类名
func (c container) ClassName(value string) container {
	return c.set("className", value)
}

// Disabled 设置是否禁用
func (c container) Disabled(value bool) container {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c container) DisabledOn(value string) container {
	return c.set("disabledOn", value)
}

// Draggable 设置是否开启容器拖拽
func (c container) Draggable(value bool) container {
	return c.set("draggable", value)
}

// DraggableConfig 设置是否开启容器拖拽配置
func (c container) DraggableConfig(value string) container {
	return c.set("draggableConfig", value)
}

// EditorSetting 设置编辑器配置
func (c container) EditorSetting(value string) container {
	return c.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (c container) Hidden(value bool) container {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c container) HiddenOn(value string) container {
	return c.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (c container) ID(value string) container {
	return c.set("id", value)
}

// OnEvent 设置事件动作配置
func (c container) OnEvent(value string) container {
	return c.set("onEvent", value)
}

// Static 设置是否静态展示
func (c container) Static(value bool) container {
	return c.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (c container) StaticClassName(value string) container {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (c container) StaticInputClassName(value string) container {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (c container) StaticLabelClassName(value string) container {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (c container) StaticOn(value string) container {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c container) StaticPlaceholder(value string) container {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (c container) StaticSchema(value string) container {
	return c.set("staticSchema", value)
}

// Style 设置自定义样式
func (c container) Style(value string) container {
	return c.set("style", value)
}

// TestIdBuilder 设置测试 id 构建器
func (c container) TestIdBuilder(value string) container {
	return c.set("testIdBuilder", value)
}

// Testid 设置测试 id
func (c container) Testid(value string) container {
	return c.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c container) UseMobileUI(value bool) container {
	return c.set("useMobileUI", value)
}

// Visible 设置是否显示
func (c container) Visible(value bool) container {
	return c.set("visible", value)
}

// VisibleOn 设置是否显示表达式
func (c container) VisibleOn(value string) container {
	return c.set("visibleOn", value)
}

// WrapperBody 设置是否需要对 body 加一层 div 包裹
func (c container) WrapperBody(value bool) container {
	return c.set("wrapperBody", value)
}

// WrapperComponent 设置使用的标签
func (c container) WrapperComponent(value string) container {
	return c.set("wrapperComponent", value)
}
