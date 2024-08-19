package comp

// comboCondition 表示组合条件渲染器
type comboCondition schema

// ComboCondition 创建一个新的 ComboCondition 实例
func ComboCondition() comboCondition {
	return make(comboCondition)
}

func (cc comboCondition) set(key string, value any) comboCondition {
	cc[key] = value
	return cc
}

// Items 设置项
func (cc comboCondition) Items(value ...any) comboCondition {
	return cc.set("items", value)
}

// Label 设置标签
func (cc comboCondition) Label(value string) comboCondition {
	return cc.set("label", value)
}

// Mode 设置模式
func (cc comboCondition) Mode(value string) comboCondition {
	return cc.set("mode", value)
}

// Scaffold 设置脚手架
func (cc comboCondition) Scaffold(value string) comboCondition {
	return cc.set("scaffold", value)
}

// Test 设置测试值
func (cc comboCondition) Test(value string) comboCondition {
	return cc.set("test", value)
}
