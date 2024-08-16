package renderers

// BaseApi 基础 API 渲染器
type BaseApi struct {
	*BaseRenderer
}

// NewBaseApi 创建一个新的 BaseApi 实例
func NewBaseApi() *BaseApi {
	return &BaseApi{BaseRenderer: NewBaseRenderer()}
}

// Set 设置属性
func (b *BaseApi) set(key string, value interface{}) *BaseApi {
	b.BaseRenderer.set(key, value)
	return b
}

// AttachDataToQuery 如果 method 为 get 的接口，设置了 data 信息
func (b *BaseApi) AttachDataToQuery(value bool) *BaseApi {
	return b.set("attachDataToQuery", value)
}

// AutoRefresh 是否自动刷新
func (b *BaseApi) AutoRefresh(value bool) *BaseApi {
	return b.set("autoRefresh", value)
}

// Cache 如果设置了值，同一个接口，相同参数，指定的时间（单位：ms）内请求将直接走缓存
func (b *BaseApi) Cache(value string) *BaseApi {
	return b.set("cache", value)
}

// ConcatDataFields 是否将两次返回的数据字段做一个合并
func (b *BaseApi) ConcatDataFields(value string) *BaseApi {
	return b.set("concatDataFields", value)
}

// ConvertKeyToPath 默认数据映射中的key如果带点或带大括号，会转成对象
func (b *BaseApi) ConvertKeyToPath(value bool) *BaseApi {
	return b.set("convertKeyToPath", value)
}

// Data 用来控制携带数据
func (b *BaseApi) Data(value string) *BaseApi {
	return b.set("data", value)
}

// DataType 发送体的格式
func (b *BaseApi) DataType(value string) *BaseApi {
	return b.set("dataType", value)
}

// ForceAppendDataToQuery 强制将数据附加在 query
func (b *BaseApi) ForceAppendDataToQuery(value bool) *BaseApi {
	return b.set("forceAppendDataToQuery", value)
}

// Headers 携带 headers，用法和 data 一样
func (b *BaseApi) Headers(value string) *BaseApi {
	return b.set("headers", value)
}

// Messages 提示信息
func (b *BaseApi) Messages(value string) *BaseApi {
	return b.set("messages", value)
}

// Method API 发送类型
func (b *BaseApi) Method(value string) *BaseApi {
	return b.set("method", value)
}

// QsOptions qs 配置项
func (b *BaseApi) QsOptions(value string) *BaseApi {
	return b.set("qsOptions", value)
}

// ReplaceData 默认都是追加模式，如果想完全替换
func (b *BaseApi) ReplaceData(value bool) *BaseApi {
	return b.set("replaceData", value)
}

// ResponseData 用来做接口返回的数据映射
func (b *BaseApi) ResponseData(value string) *BaseApi {
	return b.set("responseData", value)
}

// ResponseType 如果是文件下载接口，请配置这个
func (b *BaseApi) ResponseType(value string) *BaseApi {
	return b.set("responseType", value)
}

// SendOn 设置发送条件
func (b *BaseApi) SendOn(value string) *BaseApi {
	return b.set("sendOn", value)
}

// Silent autoFill 是否显示自动填充错误提示
func (b *BaseApi) Silent(value bool) *BaseApi {
	return b.set("silent", value)
}

// TrackExpression 当开启自动刷新的时候，默认是 api 的 url 来自动跟踪变量变化的
func (b *BaseApi) TrackExpression(value string) *BaseApi {
	return b.set("trackExpression", value)
}

// Url API 发送目标地址
func (b *BaseApi) Url(value string) *BaseApi {
	return b.set("url", value)
}
