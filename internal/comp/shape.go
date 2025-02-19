package comp

import "github.com/zrcoder/amisgo/schema"

type Shape schema.Schema

func NewShape() Shape {
	return Shape{"type": "shape"}
}

func (s Shape) set(key string, value any) Shape {
	s[key] = value
	return s
}

// ShapeType sets the shape type:
// square, triangle, right-triangle, rectangle, convex-arc-rectangle, concave-arc-rectangle,
// double-convex-arc-rectangle, double-concave-arc-rectangle, barrel-rectangle, rhombus,
// parallelogram, rectangle-1, rectangle-2, rectangle-3, heart, pentagon, hexagon, octagon,
// star, hexagon-star, circle, arrow
// custom
func (s Shape) ShapeType(value string) Shape {
	return s.set("shapeType", value)
}

// ClassName sets the CSS class name for the code component
func (s Shape) ClassName(value string) Shape {
	return s.set("className", value)
}

// Color sets the color of the fill color
func (s Shape) Color(value string) Shape {
	return s.set("color", value)
}

// Width sets the shape's width
func (s Shape) Width(value float64) Shape {
	return s.set("width", value)
}

// Height sets the shape's height
func (s Shape) Height(value float64) Shape {
	return s.set("height", value)
}

// Radius sets the corner radius, negative value means inner arc
func (s Shape) Radius(value float64) Shape {
	return s.set("radius", value)
}

// Stroke sets the stroke color
func (s Shape) Stroke(value string) Shape {
	return s.set("stroke", value)
}

// StrokeWidth sets the stroke width
func (s Shape) StrokeWidth(value float64) Shape {
	return s.set("strokeWidth", value)
}

// StrokeType sets the stroke type:
// line, dash, dot
func (s Shape) StrokeType(value string) Shape {
	return s.set("strokeType", value)
}

// Paths sets the paths when shapge type is custom
func (s Shape) Paths(value ...string) Shape {
	return s.set("paths", value)
}

// OnEvent sets the event action configuration
func (s Shape) OnEvent(value any) Shape {
	return s.set("onEvent", value)
}
