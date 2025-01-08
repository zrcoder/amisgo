package comp

import "github.com/zrcoder/amisgo/model"

type MApi model.Schema

func Api() MApi {
	return make(MApi)
}

func (b MApi) set(key string, value any) MApi {
	b[key] = value
	return b
}

// AttachDataToQuery sets data in the query if the method is GET
func (b MApi) AttachDataToQuery(value bool) MApi {
	return b.set("attachDataToQuery", value)
}

// AutoRefresh enables or disables auto-refresh
func (b MApi) AutoRefresh(value bool) MApi {
	return b.set("autoRefresh", value)
}

// Cache sets the cache duration in milliseconds
func (b MApi) Cache(value string) MApi {
	return b.set("cache", value)
}

// ConcatDataFields merges data fields from multiple responses
func (b MApi) ConcatDataFields(value string) MApi {
	return b.set("concatDataFields", value)
}

// ConvertKeyToPath converts keys with dots or braces to objects
func (b MApi) ConvertKeyToPath(value bool) MApi {
	return b.set("convertKeyToPath", value)
}

// Data sets the data to be sent
func (b MApi) Data(value model.Data) MApi {
	return b.set("data", value)
}

// DataType sets the format of the request body
func (b MApi) DataType(value string) MApi {
	return b.set("dataType", value)
}

// ForceAppendDataToQuery forces data to be appended to the query
func (b MApi) ForceAppendDataToQuery(value bool) MApi {
	return b.set("forceAppendDataToQuery", value)
}

// Headers sets the request headers
func (b MApi) Headers(value any) MApi {
	return b.set("headers", value)
}

// Messages sets the message to be displayed
func (b MApi) Messages(value string) MApi {
	return b.set("messages", value)
}

// Method sets the HTTP method
func (b MApi) Method(value string) MApi {
	return b.set("method", value)
}

// QsOptions sets the query string options
func (b MApi) QsOptions(value ...any) MApi {
	return b.set("qsOptions", value)
}

// ReplaceData replaces the existing data
func (b MApi) ReplaceData(value bool) MApi {
	return b.set("replaceData", value)
}

// ResponseData maps the response data
func (b MApi) ResponseData(value string) MApi {
	return b.set("responseData", value)
}

// ResponseType sets the response type, e.g., for file downloads
func (b MApi) ResponseType(value string) MApi {
	return b.set("responseType", value)
}

// SendOn sets the condition for sending the request
func (b MApi) SendOn(value string) MApi {
	return b.set("sendOn", value)
}

// Silent enables or disables silent mode for auto-fill errors
func (b MApi) Silent(value bool) MApi {
	return b.set("silent", value)
}

// TrackExpression sets the expression to track for auto-refresh
func (b MApi) TrackExpression(value string) MApi {
	return b.set("trackExpression", value)
}

// Url sets the target URL for the API request
func (b MApi) Url(value string) MApi {
	return b.set("url", value)
}

// RequestAdaptor customizes the request before sending
func (b MApi) RequestAdaptor(value string) MApi {
	return b.set("requestAdaptor", value)
}

// Adaptor processes the response to match the expected format
func (b MApi) Adaptor(value string) MApi {
	return b.set("adaptor", value)
}
