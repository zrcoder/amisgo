package comp

import "github.com/zrcoder/amisgo/model"

// Barcode represents a barcode renderer
type barcode model.Schema

// Barcode creates a new Barcode instance
func Barcode() barcode {
	return barcode{"type": "barcode"}
}

func (b barcode) set(key string, value any) barcode {
	b[key] = value
	return b
}

// ClassName sets the CSS class name for the outer container
func (b barcode) ClassName(value string) barcode {
	return b.set("className", value)
}
