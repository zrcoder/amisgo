package comp

// qrCodeImageSettings

type qrCodeImageSettings Schema

func QRCodeImageSettings() qrCodeImageSettings {
	return qrCodeImageSettings{}
}

func (q qrCodeImageSettings) set(key string, value any) qrCodeImageSettings {
	q[key] = value
	return q
}

// Excavate sets the "excavate" property of qrCodeImageSettings and returns the updated settings.
func (q qrCodeImageSettings) Excavate(value bool) qrCodeImageSettings {
	return q.set("excavate", value)
}

// Height
func (q qrCodeImageSettings) Height(value string) qrCodeImageSettings {
	return q.set("height", value)
}

// Src
func (q qrCodeImageSettings) Src(value string) qrCodeImageSettings {
	return q.set("src", value)
}

// Width
func (q qrCodeImageSettings) Width(value string) qrCodeImageSettings {
	return q.set("width", value)
}

// X
func (q qrCodeImageSettings) X(value string) qrCodeImageSettings {
	return q.set("x", value)
}

// Y
func (q qrCodeImageSettings) Y(value string) qrCodeImageSettings {
	return q.set("y", value)
}
