package model

// ComboCondition represents a composite condition rendering component
type ComboCondition Schema

func NewComboCondition() ComboCondition {
	return make(ComboCondition)
}

func (cc ComboCondition) set(key string, value any) ComboCondition {
	cc[key] = value
	return cc
}

// Items sets the collection of condition items
func (cc ComboCondition) Items(value ...any) ComboCondition {
	return cc.set("items", value)
}

// Label sets the text label for the condition component
func (cc ComboCondition) Label(value string) ComboCondition {
	return cc.set("label", value)
}

// Mode sets the operational mode for the condition component
func (cc ComboCondition) Mode(value string) ComboCondition {
	return cc.set("mode", value)
}

// Scaffold sets the structural template or framework for the condition
func (cc ComboCondition) Scaffold(value string) ComboCondition {
	return cc.set("scaffold", value)
}

// Test sets the condition test expression or value
func (cc ComboCondition) Test(value string) ComboCondition {
	return cc.set("test", value)
}
