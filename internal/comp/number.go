package comp

import "github.com/zrcoder/amisgo/model"

type Number model.Schema

func NewNumber() Number {
	return Number{"type": "number"}
}

func (n Number) set(key string, value any) Number {
	n[key] = value
	return n
}

// ClassName sets the style classname
func (n Number) ClassName(value string) Number {
	return n.set("className", value)
}

// Value sets the number's value
func (n Number) Value(value string) Number {
	return n.set("value", value)
}

// Name sets the component's name
func (n Number) Name(value string) Number {
	return n.set("name", value)
}

// Placeholder sets the placeholder
func (n Number) Placeholder(value string) Number {
	return n.set("placeholder", value)
}

// KilobitSeparator sets whether to display the number with thousands separator
func (n Number) KilobitSeparator(value bool) Number {
	return n.set("kilobitSeparator", value)
}

// Precision sets the number of decimal places
func (n Number) Precision(value int) Number {
	return n.set("precision", value)
}

// Percent sets whether to display as percentage.
// If value is a number, it controls the decimal places of the percentage
func (n Number) Percent(value any) Number {
	return n.set("percent", value)
}

// Prefix sets the prefix
func (n Number) Prefix(value string) Number {
	return n.set("prefix", value)
}

// Affix sets the suffix
func (n Number) Affix(value string) Number {
	return n.set("affix", value)
}
