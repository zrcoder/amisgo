package comp

// QRCodeImageSettings
//
// @version 6.7.0
type QRCodeImageSettings Schema

// NewQRCodeImageSettings 创建一个新的 QRCodeImageSettings 实例
func NewQRCodeImageSettings() QRCodeImageSettings {
	return QRCodeImageSettings{}
}

func (q QRCodeImageSettings) set(key string, value interface{}) QRCodeImageSettings {
	q[key] = value
	return q
}

// Excavate
func (q QRCodeImageSettings) Excavate(value bool) QRCodeImageSettings {
	return q.set("excavate", value)
}

// Height
func (q QRCodeImageSettings) Height(value string) QRCodeImageSettings {
	return q.set("height", value)
}

// Src
func (q QRCodeImageSettings) Src(value string) QRCodeImageSettings {
	return q.set("src", value)
}

// Width
func (q QRCodeImageSettings) Width(value string) QRCodeImageSettings {
	return q.set("width", value)
}

// X
func (q QRCodeImageSettings) X(value string) QRCodeImageSettings {
	return q.set("x", value)
}

// Y
func (q QRCodeImageSettings) Y(value string) QRCodeImageSettings {
	return q.set("y", value)
}
