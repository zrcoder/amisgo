package comp

import "github.com/zrcoder/amisgo/model"

// schemaApi represents a schema API configuration.
type schemaApi model.Schema

// model.SchemaApi creates a new schemaApi instance.
func SchemaApi() schemaApi {
	return schemaApi{}
}

func (s schemaApi) set(key string, value any) schemaApi {
	s[key] = value
	return s
}

// AttachDataToQuery sets whether to attach data to the query.
func (s schemaApi) AttachDataToQuery(value bool) schemaApi {
	return s.set("attachDataToQuery", value)
}

// AutoRefresh sets whether to auto refresh.
func (s schemaApi) AutoRefresh(value bool) schemaApi {
	return s.set("autoRefresh", value)
}

// Cache sets the cache duration.
func (s schemaApi) Cache(value string) schemaApi {
	return s.set("cache", value)
}

// ConvertKeyToPath sets whether to convert key to path.
func (s schemaApi) ConvertKeyToPath(value bool) schemaApi {
	return s.set("convertKeyToPath", value)
}

// Data sets the data configuration.
func (s schemaApi) Data(value any) schemaApi {
	return s.set("data", value)
}

// DataType sets the data type.
func (s schemaApi) DataType(value string) schemaApi {
	return s.set("dataType", value)
}

// DownloadFileName sets the download file name.
func (s schemaApi) DownloadFileName(value string) schemaApi {
	return s.set("downloadFileName", value)
}

// ForceAppendDataToQuery sets whether to force append data to query.
func (s schemaApi) ForceAppendDataToQuery(value bool) schemaApi {
	return s.set("forceAppendDataToQuery", value)
}

// Headers sets the headers configuration.
func (s schemaApi) Headers(value string) schemaApi {
	return s.set("headers", value)
}

// Method sets the HTTP method.
func (s schemaApi) Method(value string) schemaApi {
	return s.set("method", value)
}

// QsOptions sets the query string options.
func (s schemaApi) QsOptions(value ...any) schemaApi {
	return s.set("qsOptions", value)
}

// ReplaceData sets whether to replace data.
func (s schemaApi) ReplaceData(value bool) schemaApi {
	return s.set("replaceData", value)
}

// ResponseData sets the response data configuration.
func (s schemaApi) ResponseData(value string) schemaApi {
	return s.set("responseData", value)
}

// ResponseType sets the response type.
func (s schemaApi) ResponseType(value string) schemaApi {
	return s.set("responseType", value)
}

// SendOn sets the send condition.
func (s schemaApi) SendOn(value string) schemaApi {
	return s.set("sendOn", value)
}

// Silent sets whether to silence errors.
func (s schemaApi) Silent(value bool) schemaApi {
	return s.set("silent", value)
}

// TrackExpression sets the track expression.
func (s schemaApi) TrackExpression(value string) schemaApi {
	return s.set("trackExpression", value)
}

// Url sets the URL.
func (s schemaApi) Url(value string) schemaApi {
	return s.set("url", value)
}
