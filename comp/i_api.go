package comp

type MApi schema

func Api() MApi {
	return make(MApi)
}

func (b MApi) set(key string, value any) MApi {
	b[key] = value
	return b
}

// AttachDataToQuery 如果 method 为 get 的接口，设置了 data 信息
func (b MApi) AttachDataToQuery(value bool) MApi {
	return b.set("attachDataToQuery", value)
}

// AutoRefresh 是否自动刷新
func (b MApi) AutoRefresh(value bool) MApi {
	return b.set("autoRefresh", value)
}

// Cache 如果设置了值，同一个接口，相同参数，指定的时间（单位：ms）内请求将直接走缓存
func (b MApi) Cache(value string) MApi {
	return b.set("cache", value)
}

// ConcatDataFields 是否将两次返回的数据字段做一个合并
func (b MApi) ConcatDataFields(value string) MApi {
	return b.set("concatDataFields", value)
}

// ConvertKeyToPath 默认数据映射中的key如果带点或带大括号，会转成对象
func (b MApi) ConvertKeyToPath(value bool) MApi {
	return b.set("convertKeyToPath", value)
}

// Data 用来控制携带数据
func (b MApi) Data(value Data) MApi {
	return b.set("data", value)
}

// DataType 发送体的格式
func (b MApi) DataType(value string) MApi {
	return b.set("dataType", value)
}

// ForceAppendDataToQuery 强制将数据附加在 query
func (b MApi) ForceAppendDataToQuery(value bool) MApi {
	return b.set("forceAppendDataToQuery", value)
}

// Headers 携带 headers，用法和 data 一样
func (b MApi) Headers(value any) MApi {
	return b.set("headers", value)
}

// Messages 提示信息
func (b MApi) Messages(value string) MApi {
	return b.set("messages", value)
}

// Method API 发送类型
func (b MApi) Method(value string) MApi {
	return b.set("method", value)
}

// QsOptions qs 配置项
func (b MApi) QsOptions(value ...MOption) MApi {
	return b.set("qsOptions", value)
}

// ReplaceData 默认都是追加模式，如果想完全替换
func (b MApi) ReplaceData(value bool) MApi {
	return b.set("replaceData", value)
}

// ResponseData 用来做接口返回的数据映射
func (b MApi) ResponseData(value string) MApi {
	return b.set("responseData", value)
}

// ResponseType 如果是文件下载接口，请配置这个
func (b MApi) ResponseType(value string) MApi {
	return b.set("responseType", value)
}

// SendOn 设置发送条件
func (b MApi) SendOn(value string) MApi {
	return b.set("sendOn", value)
}

// Silent autoFill 是否显示自动填充错误提示
func (b MApi) Silent(value bool) MApi {
	return b.set("silent", value)
}

// TrackExpression 当开启自动刷新的时候，默认是 api 的 url 来自动跟踪变量变化的
func (b MApi) TrackExpression(value string) MApi {
	return b.set("trackExpression", value)
}

// Url API 发送目标地址
func (b MApi) Url(value string) MApi {
	return b.set("url", value)
}

// RequestAdaptor 发送适配器，在接口请求前，对请求进行一些自定义处理，例如修改发送数据体、添加请求头、等等，基本用法是，获取暴露的api参数，并且对该参数进行一些修改
func (b MApi) RequestAdaptor(value string) MApi {
	return b.set("requestAdaptor", value)
}

// Adaptor 接收适配器，如果接口返回不符合要求，可以通过配置一个适配器来处理成 amis 需要的数据
// 如： "console.log(context); // 打印上下文数据 \nreturn {\n    ...payload,\n    status: payload.code === 200 ? 0 : payload.code\n}"
func (b MApi) Adaptor(value string) MApi {
	return b.set("adaptor", value)
}
