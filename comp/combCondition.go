package comp

// ComboCondition represents a composite condition rendering component
type comboCondition Schema

// ComboCondition creates a new ComboCondition instance
func ComboCondition() comboCondition {
	return make(comboCondition)
}

func (cc comboCondition) set(key string, value any) comboCondition {
	cc[key] = value
	return cc
}

// Items sets the collection of condition items
func (cc comboCondition) Items(value ...any) comboCondition {
	return cc.set("items", value)
}

// Label sets the text label for the condition component
func (cc comboCondition) Label(value string) comboCondition {
	return cc.set("label", value)
}

// Mode sets the operational mode for the condition component
func (cc comboCondition) Mode(value string) comboCondition {
	return cc.set("mode", value)
}

// Scaffold sets the structural template or framework for the condition
func (cc comboCondition) Scaffold(value string) comboCondition {
	return cc.set("scaffold", value)
}

// Test sets the condition test expression or value
func (cc comboCondition) Test(value string) comboCondition {
	return cc.set("test", value)
}
