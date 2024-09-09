package comp

// tasks 渲染器，格式说明 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/tasks

type tasks schema

// Tasks 创建一个新的 Tasks 实例
func Tasks() tasks {
	return tasks{}.set("type", "tasks")
}

// set 用于设置属性值，返回更新后的 Tasks 实例
func (t tasks) set(key string, value any) tasks {
	t[key] = value
	return t
}

// BtnClassName css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     }
func (t tasks) BtnClassName(value string) tasks {
	return t.set("btnClassName", value)
}

// BtnText 操作按钮文字
func (t tasks) BtnText(value string) tasks {
	return t.set("btnText", value)
}

// CanRetryStatusCode
func (t tasks) CanRetryStatusCode(value string) tasks {
	return t.set("canRetryStatusCode", value)
}

// CheckApi 用来获取任务状态的 API，当没有进行时任务时不会发送。 (用来获取任务状态的 API，当没有进行时任务时不会发送。)
func (t tasks) CheckApi(value string) tasks {
	return t.set("checkApi", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t tasks) ClassName(value string) tasks {
	return t.set("className", value)
}

// Disabled 是否禁用
func (t tasks) Disabled(value bool) tasks {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t tasks) DisabledOn(value string) tasks {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t tasks) EditorSetting(value string) tasks {
	return t.set("editorSetting", value)
}

// ErrorStatusCode
func (t tasks) ErrorStatusCode(value string) tasks {
	return t.set("errorStatusCode", value)
}

// FinishStatusCode
func (t tasks) FinishStatusCode(value string) tasks {
	return t.set("finishStatusCode", value)
}

// Hidden 是否隐藏
func (t tasks) Hidden(value bool) tasks {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t tasks) HiddenOn(value string) tasks {
	return t.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t tasks) Id(value string) tasks {
	return t.set("id", value)
}

// InitialStatusCode
func (t tasks) InitialStatusCode(value string) tasks {
	return t.set("initialStatusCode", value)
}

// Interval 当有任务进行中，会每隔一段时间再次检测，而时间间隔就是通过此项配置，默认 3s。
func (t tasks) Interval(value string) tasks {
	return t.set("interval", value)
}

// Items
func (t tasks) Items(value ...any) tasks {
	return t.set("items", value)
}

// LoadingConfig
func (t tasks) LoadingConfig(value string) tasks {
	return t.set("loadingConfig", value)
}

// LoadingStatusCode
func (t tasks) LoadingStatusCode(value string) tasks {
	return t.set("loadingStatusCode", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (t tasks) Name(value string) tasks {
	return t.set("name", value)
}

// OnEvent 事件动作配置
func (t tasks) OnEvent(value any) tasks {
	return t.set("onEvent", value)
}

// OperationLabel 操作列说明
func (t tasks) OperationLabel(value string) tasks {
	return t.set("operationLabel", value)
}

// ReSubmitApi 如果任务失败，且可以重试，提交的时候会使用此 API (如果任务失败，且可以重试，提交的时候会使用此 API)
func (t tasks) ReSubmitApi(value string) tasks {
	return t.set("reSubmitApi", value)
}

// ReadyStatusCode
func (t tasks) ReadyStatusCode(value string) tasks {
	return t.set("readyStatusCode", value)
}

// RemarkLabel 备注列说明
func (t tasks) RemarkLabel(value string) tasks {
	return t.set("remarkLabel", value)
}

// RetryBtnClassName 配置容器重试按钮 className (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t tasks) RetryBtnClassName(value string) tasks {
	return t.set("retryBtnClassName", value)
}

// RetryBtnText 重试操作按钮文字
func (t tasks) RetryBtnText(value string) tasks {
	return t.set("retryBtnText", value)
}

// Static 是否静态展示
func (t tasks) Static(value bool) tasks {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t tasks) StaticClassName(value string) tasks {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t tasks) StaticInputClassName(value string) tasks {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t tasks) StaticLabelClassName(value string) tasks {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t tasks) StaticOn(value string) tasks {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t tasks) StaticPlaceholder(value string) tasks {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t tasks) StaticSchema(value string) tasks {
	return t.set("staticSchema", value)
}

// StatusLabel 状态列说明
func (t tasks) StatusLabel(value string) tasks {
	return t.set("statusLabel", value)
}

// StatusLabelMap 状态显示对应的类名配置。
func (t tasks) StatusLabelMap(value string) tasks {
	return t.set("statusLabelMap", value)
}

// StatusTextMap 状态显示对应的文字显示配置。
func (t tasks) StatusTextMap(value string) tasks {
	return t.set("statusTextMap", value)
}

// Style 组件样式
func (t tasks) Style(value any) tasks {
	return t.set("style", value)
}

// SubmitApi 提交任务使用的 API (提交任务使用的 API)
func (t tasks) SubmitApi(value string) tasks {
	return t.set("submitApi", value)
}

// TableClassName 配置 table className (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t tasks) TableClassName(value string) tasks {
	return t.set("tableClassName", value)
}

// TaskNameLabel 任务名称列说明
func (t tasks) TaskNameLabel(value string) tasks {
	return t.set("taskNameLabel", value)
}

// TestIdBuilder
func (t tasks) TestIdBuilder(value string) tasks {
	return t.set("testIdBuilder", value)
}

// Testid
func (t tasks) Testid(value string) tasks {
	return t.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t tasks) UseMobileUI(value bool) tasks {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t tasks) Visible(value bool) tasks {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t tasks) VisibleOn(value string) tasks {
	return t.set("visibleOn", value)
}
