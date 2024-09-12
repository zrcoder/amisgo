package comp

// service 服务类控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/service

type service schema

// Service 创建一个新的 Service 实例
func Service() service {
	return service{}.set("type", "service")
}

func (s service) set(key string, value any) service {
	s[key] = value
	return s
}

// Api 页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。 (页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。)
func (s service) Api(value string) service {
	return s.set("api", value)
}

// FetchData 通过内置 api 获取数据
func (s service) FetchData(getter func() any) service {
	return s.Api(serveData(getter))
}

// Body 内容区域 (内容区域)
func (s service) Body(value ...any) service {
	return s.set("body", value)
}

// Data 数据
func (s service) Data(value any) service {
	return s.set("data", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s service) ClassName(value string) service {
	return s.set("className", value)
}

// DataProvider 通过调用外部函数来获取数据 (通过调用外部函数来获取数据)
func (s service) DataProvider(value string) service {
	return s.set("dataProvider", value)
}

// Disabled 是否禁用
func (s service) Disabled(value bool) service {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s service) DisabledOn(value string) service {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s service) EditorSetting(value string) service {
	return s.set("editorSetting", value)
}

// FetchOn 表达式，语法 `data.xxx > 5`。
func (s service) FetchOn(value string) service {
	return s.set("fetchOn", value)
}

// Hidden 是否隐藏
func (s service) Hidden(value bool) service {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s service) HiddenOn(value string) service {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s service) Id(value string) service {
	return s.set("id", value)
}

// InitFetch 是否默认就拉取？
func (s service) InitFetch(value bool) service {
	return s.set("initFetch", value)
}

// InitFetchOn 是否默认就拉取？通过表达式来决定. (表达式，语法 `data.xxx > 5`。)
func (s service) InitFetchOn(value string) service {
	return s.set("initFetchOn", value)
}

// InitFetchSchema 是否默认加载 schemaApi
func (s service) InitFetchSchema(value bool) service {
	return s.set("initFetchSchema", value)
}

// InitFetchSchemaOn 用表达式来配置。 (表达式，语法 `data.xxx > 5`。)
func (s service) InitFetchSchemaOn(value string) service {
	return s.set("initFetchSchemaOn", value)
}

// Interval 是否轮询拉取
func (s service) Interval(value string) service {
	return s.set("interval", value)
}

// LoadingConfig
func (s service) LoadingConfig(value string) service {
	return s.set("loadingConfig", value)
}

// Messages 消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。
func (s service) Messages(value string) service {
	return s.set("messages", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (s service) Name(value string) service {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s service) OnEvent(value any) service {
	return s.set("onEvent", value)
}

// SchemaApi 用来获取远程 Schema 的 api (用来获取远程 Schema 的 api)
func (s service) SchemaApi(value string) service {
	return s.set("schemaApi", value)
}

// ShowErrorMsg 是否以Alert的形式显示api接口响应的错误信息，默认展示
func (s service) ShowErrorMsg(value bool) service {
	return s.set("showErrorMsg", value)
}

// SilentPolling 是否静默拉取
func (s service) SilentPolling(value bool) service {
	return s.set("silentPolling", value)
}

// Static 是否静态展示
func (s service) Static(value bool) service {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s service) StaticClassName(value string) service {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s service) StaticInputClassName(value string) service {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s service) StaticLabelClassName(value string) service {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s service) StaticOn(value string) service {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s service) StaticPlaceholder(value string) service {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s service) StaticSchema(value string) service {
	return s.set("staticSchema", value)
}

// StopAutoRefreshWhen 关闭轮询的条件。 (表达式，语法 `data.xxx > 5`。)
func (s service) StopAutoRefreshWhen(value string) service {
	return s.set("stopAutoRefreshWhen", value)
}

// Style 组件样式
func (s service) Style(value any) service {
	return s.set("style", value)
}

// TestIdBuilder
func (s service) TestIdBuilder(value string) service {
	return s.set("testIdBuilder", value)
}

// Testid
func (s service) Testid(value string) service {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s service) UseMobileUI(value bool) service {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s service) Visible(value bool) service {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s service) VisibleOn(value string) service {
	return s.set("visibleOn", value)
}

// Ws WebScocket 地址，用于实时获取数据
func (s service) Ws(value string) service {
	return s.set("ws", value)
}
