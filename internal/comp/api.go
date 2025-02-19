package comp

import "github.com/zrcoder/amisgo/schema"

type Api schema.Schema

func NewApi() Api {
	return make(Api)
}

func (b Api) Set(key string, value any) Api {
	b[key] = value
	return b
}

// AttachDataToQuery sets data in the query if the method is GET
func (b Api) AttachDataToQuery(value bool) Api {
	return b.Set("attachDataToQuery", value)
}

// AutoRefresh enables or disables auto-refresh
func (b Api) AutoRefresh(value bool) Api {
	return b.Set("autoRefresh", value)
}

// Cache sets the cache duration in milliseconds
func (b Api) Cache(value int) Api {
	return b.Set("cache", value)
}

// ConcatDataFields merges data fields from multiple responses
func (b Api) ConcatDataFields(value string) Api {
	return b.Set("concatDataFields", value)
}

// ConvertKeyToPath converts keys with dots or braces to objects
func (b Api) ConvertKeyToPath(value bool) Api {
	return b.Set("convertKeyToPath", value)
}

// Data sets the data to be sent
func (b Api) Data(value any) Api {
	return b.Set("data", value)
}

// DataType sets the format of the request body
func (b Api) DataType(value string) Api {
	return b.Set("dataType", value)
}

// ForceAppendDataToQuery forces data to be appended to the query
func (b Api) ForceAppendDataToQuery(value bool) Api {
	return b.Set("forceAppendDataToQuery", value)
}

// Headers sets the request headers
func (b Api) Headers(value any) Api {
	return b.Set("headers", value)
}

// Messages sets the message to be displayed
func (b Api) Messages(value any) Api {
	return b.Set("messages", value)
}

// Method sets the HTTP method
func (b Api) Method(value string) Api {
	return b.Set("method", value)
}

// QsOptions sets the query string options
func (b Api) QsOptions(value any) Api {
	return b.Set("qsOptions", value)
}

// ReplaceData replaces the existing data
func (b Api) ReplaceData(value bool) Api {
	return b.Set("replaceData", value)
}

// ResponseData maps the response data
func (b Api) ResponseData(value any) Api {
	return b.Set("responseData", value)
}

// ResponseType sets the response type, e.g., for file downloads
func (b Api) ResponseType(value string) Api {
	return b.Set("responseType", value)
}

// SendOn sets the condition for sending the request
func (b Api) SendOn(value string) Api {
	return b.Set("sendOn", value)
}

// Silent enables or disables silent mode for auto-fill errors
func (b Api) Silent(value bool) Api {
	return b.Set("silent", value)
}

// TrackExpression sets the expression to track for auto-refresh
func (b Api) TrackExpression(value string) Api {
	return b.Set("trackExpression", value)
}

// Url sets the target URL for the API request
func (b Api) Url(value string) Api {
	return b.Set("url", value)
}

// RequestAdaptor customizes the request before sending
func (b Api) RequestAdaptor(value string) Api {
	return b.Set("requestAdaptor", value)
}

// Adaptor processes the response to match the expected format
func (b Api) Adaptor(value string) Api {
	return b.Set("adaptor", value)
}
