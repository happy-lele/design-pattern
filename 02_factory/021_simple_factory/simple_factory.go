package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// NewRuleConfigParser 由于 Go 本身是没有构造函数的
// 一般而言我们采用 NewName 的方式创建对象/接口
// 当它返回的是接口的时候，其实就是简单工厂模式
func NewRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}