package comp

// Page 代表 amis Page 渲染器
//
// @version 6.7.0
type Page Schema

// NewPage 创建一个新的 Page 实例
func NewPage() Page {
	return Page{}.set("type", "page")
}

// set 用于设置字段值
func (p Page) set(key string, value interface{}) Page {
	p[key] = value
	return p
}

// Aside 边栏区域
func (p Page) Aside(value string) Page {
	return p.set("aside", value)
}

// AsideClassName 边栏区 css 类名
func (p Page) AsideClassName(value string) Page {
	return p.set("asideClassName", value)
}

// AsideMaxWidth 边栏最小宽度
func (p Page) AsideMaxWidth(value string) Page {
	return p.set("asideMaxWidth", value)
}

// AsideMinWidth 边栏最小宽度
func (p Page) AsideMinWidth(value string) Page {
	return p.set("asideMinWidth", value)
}

// AsideResizor 边栏是否允许拖动
func (p Page) AsideResizor(value bool) Page {
	return p.set("asideResizor", value)
}

// AsideSticky 边栏内容是否粘住
func (p Page) AsideSticky(value bool) Page {
	return p.set("asideSticky", value)
}

// Body 内容区域
func (p Page) Body(value ...interface{}) Page {
	return p.set("body", value)
}

// BodyClassName 内容区 css 类名
func (p Page) BodyClassName(value string) Page {
	return p.set("bodyClassName", value)
}

// ClassName 配置容器 className
func (p Page) ClassName(value string) Page {
	return p.set("className", value)
}

// CSS 自定义页面级别样式表
func (p Page) CSS(value string) Page {
	return p.set("css", value)
}

// CSSVars css 变量
func (p Page) CSSVars(value string) Page {
	return p.set("cssVars", value)
}

// Data 页面级别的初始数据
func (p Page) Data(value string) Page {
	return p.set("data", value)
}

// Definitions 配置定义
func (p Page) Definitions(value string) Page {
	return p.set("definitions", value)
}

// Disabled 是否禁用
func (p Page) Disabled(value bool) Page {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p Page) DisabledOn(value string) Page {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p Page) EditorSetting(value string) Page {
	return p.set("editorSetting", value)
}

// HeaderClassName 配置 header 容器 className
func (p Page) HeaderClassName(value string) Page {
	return p.set("headerClassName", value)
}

// Hidden 是否隐藏
func (p Page) Hidden(value bool) Page {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p Page) HiddenOn(value string) Page {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p Page) ID(value string) Page {
	return p.set("id", value)
}

// InitApi 页面初始化 API
func (p Page) InitApi(value string) Page {
	return p.set("initApi", value)
}

// InitFetch 是否默认就拉取
func (p Page) InitFetch(value bool) Page {
	return p.set("initFetch", value)
}

// InitFetchOn 是否默认就拉取表达式
func (p Page) InitFetchOn(value string) Page {
	return p.set("initFetchOn", value)
}

// Interval 配置轮询间隔
func (p Page) Interval(value string) Page {
	return p.set("interval", value)
}

// LoadingConfig 加载配置
func (p Page) LoadingConfig(value string) Page {
	return p.set("loadingConfig", value)
}

// Messages 消息文案配置
func (p Page) Messages(value string) Page {
	return p.set("messages", value)
}

// MobileCSS 移动端下的样式表
func (p Page) MobileCSS(value string) Page {
	return p.set("mobileCSS", value)
}

// Name 组件名字
func (p Page) Name(value string) Page {
	return p.set("name", value)
}

// OnEvent 事件动作配置
func (p Page) OnEvent(value string) Page {
	return p.set("onEvent", value)
}

// PullRefresh 下拉刷新配置
func (p Page) PullRefresh(value string) Page {
	return p.set("pullRefresh", value)
}

// Regions 默认不设置自动感觉内容来决定要不要展示这些区域
func (p Page) Regions(value string) Page {
	return p.set("regions", value)
}

// Remark 页面描述
func (p Page) Remark(value string) Page {
	return p.set("remark", value)
}

// ShowErrorMsg 是否显示错误信息
func (p Page) ShowErrorMsg(value bool) Page {
	return p.set("showErrorMsg", value)
}

// SilentPolling 是否要静默加载
func (p Page) SilentPolling(value bool) Page {
	return p.set("silentPolling", value)
}

// Static 是否静态展示
func (p Page) Static(value bool) Page {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p Page) StaticClassName(value string) Page {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p Page) StaticInputClassName(value string) Page {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p Page) StaticLabelClassName(value string) Page {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p Page) StaticOn(value string) Page {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p Page) StaticPlaceholder(value string) Page {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p Page) StaticSchema(value string) Page {
	return p.set("staticSchema", value)
}

// StopAutoRefreshWhen 配置停止轮询的条件
func (p Page) StopAutoRefreshWhen(value string) Page {
	return p.set("stopAutoRefreshWhen", value)
}

// Style 自定义样式
func (p Page) Style(value string) Page {
	return p.set("style", value)
}

// SubTitle 页面副标题
func (p Page) SubTitle(value string) Page {
	return p.set("subTitle", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p Page) TestIdBuilder(value string) Page {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p Page) Testid(value string) Page {
	return p.set("testid", value)
}

// Title 页面标题
func (p Page) Title(value string) Page {
	return p.set("title", value)
}

// Toolbar 页面顶部区域
func (p Page) Toolbar(value string) Page {
	return p.set("toolbar", value)
}

// ToolbarClassName 配置 toolbar 容器 className
func (p Page) ToolbarClassName(value string) Page {
	return p.set("toolbarClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p Page) UseMobileUI(value bool) Page {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p Page) Visible(value bool) Page {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p Page) VisibleOn(value string) Page {
	return p.set("visibleOn", value)
}
