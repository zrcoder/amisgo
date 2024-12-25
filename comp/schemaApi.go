package comp

// schemaApi

type schemaApi schema

// SchemaApi 创建一个新的 SchemaApi 实例
func SchemaApi() schemaApi {
	return schemaApi{}
}

func (s schemaApi) set(key string, value any) schemaApi {
	s[key] = value
	return s
}

// AttachDataToQuery 如果 method 为 get 的接口，设置了 data 信息。 默认 data 会自动附带在 query 里面发送给后端。如果想通过 body 发送给后端，那么请把这个配置成 false。但是，浏览器还不支持啊，设置了只是摆设。除非服务端支持 method-override
func (s schemaApi) AttachDataToQuery(value bool) schemaApi {
	return s.set("attachDataToQuery", value)
}

// AutoRefresh 是否自动刷新，当 url 中的取值结果变化时，自动刷新数据。
func (s schemaApi) AutoRefresh(value bool) schemaApi {
	return s.set("autoRefresh", value)
}

// Cache 如果设置了值，同一个接口，相同参数，指定的时间（单位：ms）内请求将直接走缓存。
func (s schemaApi) Cache(value string) schemaApi {
	return s.set("cache", value)
}

// ConvertKeyToPath 默认数据映射中的key如果带点，或者带大括号，会转成对象比如：{   'a.b': '123' }经过数据映射后变成 {  a: {   b: '123  } }如果想要关闭此功能，请设置 convertKeyToPath 为 false
func (s schemaApi) ConvertKeyToPath(value bool) schemaApi {
	return s.set("convertKeyToPath", value)
}

// Data 用来控制携带数据. 当key 为 `&` 值为 `$$` 时, 将所有原始数据打平设置到 data 中. 当值为 $$ 将所有原始数据赋值到对应的 key 中. 当值为 $ 打头时, 将变量值设置到 key 中.
func (s schemaApi) Data(value Data) schemaApi {
	return s.set("data", value)
}

// DataType 发送体的格式 可选值: json | form-data | form
func (s schemaApi) DataType(value string) schemaApi {
	return s.set("dataType", value)
}

// DownloadFileName 文件下载时，指定文件名
func (s schemaApi) DownloadFileName(value string) schemaApi {
	return s.set("downloadFileName", value)
}

// ForceAppendDataToQuery 强制将数据附加在 query，默认只有 api 地址中没有用变量的时候 crud 查询接口才会 自动附加数据到 query 部分，如果想强制附加请设置这个属性。 对于那种临时加了个变量但是又不想全部参数写一遍的时候配置很有用。
func (s schemaApi) ForceAppendDataToQuery(value bool) schemaApi {
	return s.set("forceAppendDataToQuery", value)
}

// Headers 携带 headers，用法和 data 一样，可以用变量。
func (s schemaApi) Headers(value string) schemaApi {
	return s.set("headers", value)
}

// Method API 发送类型 可选值: get | post | put | delete | patch | jsonp | js
func (s schemaApi) Method(value string) schemaApi {
	return s.set("method", value)
}

// QsOptions qs 配置项
func (s schemaApi) QsOptions(value ...MOption) schemaApi {
	return s.set("qsOptions", value)
}

// ReplaceData 默认都是追加模式，如果想完全替换把这个配置成 true
func (s schemaApi) ReplaceData(value bool) schemaApi {
	return s.set("replaceData", value)
}

// ResponseData 用来做接口返回的数据映射。
func (s schemaApi) ResponseData(value string) schemaApi {
	return s.set("responseData", value)
}

// ResponseType 如果是文件下载接口，请配置这个。
func (s schemaApi) ResponseType(value string) schemaApi {
	return s.set("responseType", value)
}

// SendOn 设置发送条件 (表达式，语法 `data.xxx > 5`。)
func (s schemaApi) SendOn(value string) schemaApi {
	return s.set("sendOn", value)
}

// Silent autoFill 是否显示自动填充错误提示
func (s schemaApi) Silent(value bool) schemaApi {
	return s.set("silent", value)
}

// TrackExpression 当开启自动刷新的时候，默认是 api 的 url 来自动跟踪变量变化的。 如果你希望监控 url 外的变量，请配置 trackExpression。
func (s schemaApi) TrackExpression(value string) schemaApi {
	return s.set("trackExpression", value)
}

// Url API 发送目标地址 (API 发送目标地址)
func (s schemaApi) Url(value string) schemaApi {
	return s.set("url", value)
}
