package comp

// baseApi 基础 API 渲染器
type baseApi schema

// BaseApi 创建一个新的 BaseApi 实例
func BaseApi() baseApi {
	return make(baseApi)
}

func (b baseApi) set(key string, value interface{}) baseApi {
	b[key] = value
	return b
}

// AttachDataToQuery 如果 method 为 get 的接口，设置了 data 信息
func (b baseApi) AttachDataToQuery(value bool) baseApi {
	return b.set("attachDataToQuery", value)
}

// AutoRefresh 是否自动刷新
func (b baseApi) AutoRefresh(value bool) baseApi {
	return b.set("autoRefresh", value)
}

// Cache 如果设置了值，同一个接口，相同参数，指定的时间（单位：ms）内请求将直接走缓存
func (b baseApi) Cache(value string) baseApi {
	return b.set("cache", value)
}

// ConcatDataFields 是否将两次返回的数据字段做一个合并
func (b baseApi) ConcatDataFields(value string) baseApi {
	return b.set("concatDataFields", value)
}

// ConvertKeyToPath 默认数据映射中的key如果带点或带大括号，会转成对象
func (b baseApi) ConvertKeyToPath(value bool) baseApi {
	return b.set("convertKeyToPath", value)
}

// Data 用来控制携带数据
func (b baseApi) Data(value string) baseApi {
	return b.set("data", value)
}

// DataType 发送体的格式
func (b baseApi) DataType(value string) baseApi {
	return b.set("dataType", value)
}

// ForceAppendDataToQuery 强制将数据附加在 query
func (b baseApi) ForceAppendDataToQuery(value bool) baseApi {
	return b.set("forceAppendDataToQuery", value)
}

// Headers 携带 headers，用法和 data 一样
func (b baseApi) Headers(value string) baseApi {
	return b.set("headers", value)
}

// Messages 提示信息
func (b baseApi) Messages(value string) baseApi {
	return b.set("messages", value)
}

// Method API 发送类型
func (b baseApi) Method(value string) baseApi {
	return b.set("method", value)
}

// QsOptions qs 配置项
func (b baseApi) QsOptions(value string) baseApi {
	return b.set("qsOptions", value)
}

// ReplaceData 默认都是追加模式，如果想完全替换
func (b baseApi) ReplaceData(value bool) baseApi {
	return b.set("replaceData", value)
}

// ResponseData 用来做接口返回的数据映射
func (b baseApi) ResponseData(value string) baseApi {
	return b.set("responseData", value)
}

// ResponseType 如果是文件下载接口，请配置这个
func (b baseApi) ResponseType(value string) baseApi {
	return b.set("responseType", value)
}

// SendOn 设置发送条件
func (b baseApi) SendOn(value string) baseApi {
	return b.set("sendOn", value)
}

// Silent autoFill 是否显示自动填充错误提示
func (b baseApi) Silent(value bool) baseApi {
	return b.set("silent", value)
}

// TrackExpression 当开启自动刷新的时候，默认是 api 的 url 来自动跟踪变量变化的
func (b baseApi) TrackExpression(value string) baseApi {
	return b.set("trackExpression", value)
}

// Url API 发送目标地址
func (b baseApi) Url(value string) baseApi {
	return b.set("url", value)
}
