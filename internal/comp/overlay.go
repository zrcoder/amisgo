package comp

import "github.com/zrcoder/amisgo/schema"

type Overlay schema.Schema

func NewOverlay() Overlay {
	return Overlay{
		"type": "overlay",
	}
}

func (o Overlay) set(key string, value any) Overlay {
	o[key] = value
	return o
}

// Width string | number
func (o Overlay) Width(width any) Overlay {
	return o.set("width", width)
}

// Align "left" | "center" | "right" }
func (o Overlay) Align(align string) Overlay {
	return o.set("align", align)
}
