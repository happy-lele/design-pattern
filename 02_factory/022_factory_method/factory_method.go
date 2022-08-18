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

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// IRuleConfigParseFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct {
}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type jsonRuleConfigParserFactory struct {
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}









