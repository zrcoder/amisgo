package comp

import "github.com/zrcoder/amisgo/model"

// Barcode represents a Barcode renderer
type Barcode model.Schema

func NewBarcode() Barcode {
	return Barcode{"type": "barcode"}
}

func (b Barcode) set(key string, value any) Barcode {
	b[key] = value
	return b
}

// ClassName sets the CSS class name for the outer container
func (b Barcode) ClassName(value string) Barcode {
	return b.set("className", value)
}
