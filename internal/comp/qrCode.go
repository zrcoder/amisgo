package comp

import "github.com/zrcoder/amisgo/schema"

// QRCode
type QRCode schema.Schema

func NewQRCode() QRCode {
	return QRCode{"type": "qrcode"}
}

func (q QRCode) set(key string, value any) QRCode {
	q[key] = value
	return q
}

// BackgroundColor sets the background color
func (q QRCode) BackgroundColor(value string) QRCode {
	return q.set("backgroundColor", value)
}

// ClassName sets the container CSS class name
func (q QRCode) ClassName(value string) QRCode {
	return q.set("className", value)
}

// CodeSize sets the QR code size, default is 128
func (q QRCode) CodeSize(value int) QRCode {
	return q.set("codeSize", value)
}

// Disabled sets whether the QR code is disabled
func (q QRCode) Disabled(value bool) QRCode {
	return q.set("disabled", value)
}

// DisabledOn sets the expression to disable the QR code
func (q QRCode) DisabledOn(value string) QRCode {
	return q.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (q QRCode) EditorSetting(value string) QRCode {
	return q.set("editorSetting", value)
}

// ForegroundColor sets the foreground color
func (q QRCode) ForegroundColor(value string) QRCode {
	return q.set("foregroundColor", value)
}

// Hidden sets whether the QR code is hidden
func (q QRCode) Hidden(value bool) QRCode {
	return q.set("hidden", value)
}

// HiddenOn sets the expression to hide the QR code
func (q QRCode) HiddenOn(value string) QRCode {
	return q.set("hiddenOn", value)
}

// ID sets the unique component ID
func (q QRCode) ID(value string) QRCode {
	return q.set("id", value)
}

// ImageSettings sets the image configuration
func (q QRCode) ImageSettings(value string) QRCode {
	return q.set("imageSettings", value)
}

// Level sets the QR code complexity level ('L', 'M', 'Q', 'H')
func (q QRCode) Level(value string) QRCode {
	return q.set("level", value)
}

// Mode sets the rendering mode
func (q QRCode) Mode(value string) QRCode {
	return q.set("mode", value)
}

// Name sets the associated field name
func (q QRCode) Name(value string) QRCode {
	return q.set("name", value)
}

// Value sets the raw data, the content obtained after scanning the QR code
func (q QRCode) Value(value string) QRCode {
	return q.set("value", value)
}

// OnEvent sets the event action configuration
func (q QRCode) OnEvent(value Event) QRCode {
	return q.set("onEvent", value)
}

// Placeholder sets the placeholder
func (q QRCode) Placeholder(value string) QRCode {
	return q.set("placeholder", value)
}

// QrcodeClassName sets the QR code CSS class name
func (q QRCode) QrcodeClassName(value string) QRCode {
	return q.set("qrcodeClassName", value)
}

// Static sets whether the QR code is statically displayed
func (q QRCode) Static(value bool) QRCode {
	return q.set("static", value)
}

// StaticClassName sets the static display form item class name
func (q QRCode) StaticClassName(value string) QRCode {
	return q.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (q QRCode) StaticInputClassName(value string) QRCode {
	return q.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (q QRCode) StaticLabelClassName(value string) QRCode {
	return q.set("staticLabelClassName", value)
}

// StaticOn sets the expression to statically display the QR code
func (q QRCode) StaticOn(value string) QRCode {
	return q.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values
func (q QRCode) StaticPlaceholder(value string) QRCode {
	return q.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema
func (q QRCode) StaticSchema(value string) QRCode {
	return q.set("staticSchema", value)
}

// Style sets the component style
func (q QRCode) Style(value any) QRCode {
	return q.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (q QRCode) TestIdBuilder(value string) QRCode {
	return q.set("testIdBuilder", value)
}

// TestID sets the test ID
func (q QRCode) TestID(value string) QRCode {
	return q.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (q QRCode) UseMobileUI(value bool) QRCode {
	return q.set("useMobileUI", value)
}

// Visible sets whether the QR code is visible
func (q QRCode) Visible(value bool) QRCode {
	return q.set("visible", value)
}

// VisibleOn sets the expression to display the QR code
func (q QRCode) VisibleOn(value string) QRCode {
	return q.set("visibleOn", value)
}
