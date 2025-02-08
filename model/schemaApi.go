package model

// SchemaApi represents a schema API configuration.
type SchemaApi Schema

func NewSchemaApi() SchemaApi {
	return SchemaApi{}
}

func (s SchemaApi) set(key string, value any) SchemaApi {
	s[key] = value
	return s
}

// AttachDataToQuery sets whether to attach data to the query.
func (s SchemaApi) AttachDataToQuery(value bool) SchemaApi {
	return s.set("attachDataToQuery", value)
}

// AutoRefresh sets whether to auto refresh.
func (s SchemaApi) AutoRefresh(value bool) SchemaApi {
	return s.set("autoRefresh", value)
}

// Cache sets the cache duration.
func (s SchemaApi) Cache(value string) SchemaApi {
	return s.set("cache", value)
}

// ConvertKeyToPath sets whether to convert key to path.
func (s SchemaApi) ConvertKeyToPath(value bool) SchemaApi {
	return s.set("convertKeyToPath", value)
}

// Data sets the data configuration.
func (s SchemaApi) Data(value any) SchemaApi {
	return s.set("data", value)
}

// DataType sets the data type.
func (s SchemaApi) DataType(value string) SchemaApi {
	return s.set("dataType", value)
}

// DownloadFileName sets the download file name.
func (s SchemaApi) DownloadFileName(value string) SchemaApi {
	return s.set("downloadFileName", value)
}

// ForceAppendDataToQuery sets whether to force append data to query.
func (s SchemaApi) ForceAppendDataToQuery(value bool) SchemaApi {
	return s.set("forceAppendDataToQuery", value)
}

// Headers sets the headers configuration.
func (s SchemaApi) Headers(value string) SchemaApi {
	return s.set("headers", value)
}

// Method sets the HTTP method.
func (s SchemaApi) Method(value string) SchemaApi {
	return s.set("method", value)
}

// QsOptions sets the query string options.
func (s SchemaApi) QsOptions(value ...any) SchemaApi {
	return s.set("qsOptions", value)
}

// ReplaceData sets whether to replace data.
func (s SchemaApi) ReplaceData(value bool) SchemaApi {
	return s.set("replaceData", value)
}

// ResponseData sets the response data configuration.
func (s SchemaApi) ResponseData(value string) SchemaApi {
	return s.set("responseData", value)
}

// ResponseType sets the response type.
func (s SchemaApi) ResponseType(value string) SchemaApi {
	return s.set("responseType", value)
}

// SendOn sets the send condition.
func (s SchemaApi) SendOn(value string) SchemaApi {
	return s.set("sendOn", value)
}

// Silent sets whether to silence errors.
func (s SchemaApi) Silent(value bool) SchemaApi {
	return s.set("silent", value)
}

// TrackExpression sets the track expression.
func (s SchemaApi) TrackExpression(value string) SchemaApi {
	return s.set("trackExpression", value)
}

// Url sets the URL.
func (s SchemaApi) Url(value string) SchemaApi {
	return s.set("url", value)
}
