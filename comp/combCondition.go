package comp

// ComboCondition 表示组合条件渲染器
type ComboCondition Schema

// NewComboCondition 创建一个新的 ComboCondition 实例
func NewComboCondition() ComboCondition {
	return make(ComboCondition)
}

func (cc ComboCondition) set(key string, value interface{}) ComboCondition {
	cc[key] = value
	return cc
}

// Items 设置项
func (cc ComboCondition) Items(value string) ComboCondition {
	return cc.set("items", value)
}

// Label 设置标签
func (cc ComboCondition) Label(value string) ComboCondition {
	return cc.set("label", value)
}

// Mode 设置模式
func (cc ComboCondition) Mode(value string) ComboCondition {
	return cc.set("mode", value)
}

// Scaffold 设置脚手架
func (cc ComboCondition) Scaffold(value string) ComboCondition {
	return cc.set("scaffold", value)
}

// Test 设置测试值
func (cc ComboCondition) Test(value string) ComboCondition {
	return cc.set("test", value)
}
