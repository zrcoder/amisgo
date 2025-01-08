package comp

import "github.com/zrcoder/amisgo/model"

// qrCode

type qrCode model.Schema

// QRCode creates a new QRCode instance
func QRCode() qrCode {
	return qrCode{}.
		set("type", "qrcode")
}

func (q qrCode) set(key string, value any) qrCode {
	q[key] = value
	return q
}

// BackgroundColor sets the background color
func (q qrCode) BackgroundColor(value string) qrCode {
	return q.set("backgroundColor", value)
}

// ClassName sets the container CSS class name
func (q qrCode) ClassName(value string) qrCode {
	return q.set("className", value)
}

// CodeSize sets the QR code size, default is 128
func (q qrCode) CodeSize(value int) qrCode {
	return q.set("codeSize", value)
}

// Disabled sets whether the QR code is disabled
func (q qrCode) Disabled(value bool) qrCode {
	return q.set("disabled", value)
}

// DisabledOn sets the expression to disable the QR code
func (q qrCode) DisabledOn(value string) qrCode {
	return q.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (q qrCode) EditorSetting(value string) qrCode {
	return q.set("editorSetting", value)
}

// ForegroundColor sets the foreground color
func (q qrCode) ForegroundColor(value string) qrCode {
	return q.set("foregroundColor", value)
}

// Hidden sets whether the QR code is hidden
func (q qrCode) Hidden(value bool) qrCode {
	return q.set("hidden", value)
}

// HiddenOn sets the expression to hide the QR code
func (q qrCode) HiddenOn(value string) qrCode {
	return q.set("hiddenOn", value)
}

// ID sets the unique component ID
func (q qrCode) ID(value string) qrCode {
	return q.set("id", value)
}

// ImageSettings sets the image configuration
func (q qrCode) ImageSettings(value string) qrCode {
	return q.set("imageSettings", value)
}

// Level sets the QR code complexity level ('L', 'M', 'Q', 'H')
func (q qrCode) Level(value string) qrCode {
	return q.set("level", value)
}

// Mode sets the rendering mode
func (q qrCode) Mode(value string) qrCode {
	return q.set("mode", value)
}

// Name sets the associated field name
func (q qrCode) Name(value string) qrCode {
	return q.set("name", value)
}

// Value sets the raw data, the content obtained after scanning the QR code
func (q qrCode) Value(value string) qrCode {
	return q.set("value", value)
}

// OnEvent sets the event action configuration
func (q qrCode) OnEvent(value any) qrCode {
	return q.set("onEvent", value)
}

// Placeholder sets the placeholder
func (q qrCode) Placeholder(value string) qrCode {
	return q.set("placeholder", value)
}

// QrcodeClassName sets the QR code CSS class name
func (q qrCode) QrcodeClassName(value string) qrCode {
	return q.set("qrcodeClassName", value)
}

// Static sets whether the QR code is statically displayed
func (q qrCode) Static(value bool) qrCode {
	return q.set("static", value)
}

// StaticClassName sets the static display form item class name
func (q qrCode) StaticClassName(value string) qrCode {
	return q.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (q qrCode) StaticInputClassName(value string) qrCode {
	return q.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (q qrCode) StaticLabelClassName(value string) qrCode {
	return q.set("staticLabelClassName", value)
}

// StaticOn sets the expression to statically display the QR code
func (q qrCode) StaticOn(value string) qrCode {
	return q.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values
func (q qrCode) StaticPlaceholder(value string) qrCode {
	return q.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (q qrCode) StaticSchema(value string) qrCode {
	return q.set("staticSchema", value)
}

// Style sets the component style
func (q qrCode) Style(value any) qrCode {
	return q.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (q qrCode) TestIdBuilder(value string) qrCode {
	return q.set("testIdBuilder", value)
}

// TestID sets the test ID
func (q qrCode) TestID(value string) qrCode {
	return q.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (q qrCode) UseMobileUI(value bool) qrCode {
	return q.set("useMobileUI", value)
}

// Visible sets whether the QR code is visible
func (q qrCode) Visible(value bool) qrCode {
	return q.set("visible", value)
}

// VisibleOn sets the expression to display the QR code
func (q qrCode) VisibleOn(value string) qrCode {
	return q.set("visibleOn", value)
}
