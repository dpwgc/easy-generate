package generate

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type JavaPOGenerator struct {
	config Config
}

func NewJavaPOGenerator(text string) Generator {
	var config Config
	err := yaml.Unmarshal([]byte(text), &config)
	if err != nil {
		panic(err)
	}
	return &JavaPOGenerator{
		config: config,
	}
}

const (
	JavaPOClassStart = `
@Data
@TableName(value = "%s")
public class %sPO {
`
	JavaPOClassEnd = `}`
	JavaPOField    = `    /**
     * %s
     */
    private %s %s;

`
)

func (g *JavaPOGenerator) Build() string {
	res := ""
	res = res + fmt.Sprintln(fmt.Sprintf(JavaPOClassStart, g.config.Table, camelString(g.config.Table)))
	for _, v := range g.config.Fields {
		switch v.Type {
		case "int":
			res = res + fmt.Sprintf(JavaPOField, v.Desc, "Integer", firstToLower(camelString(v.Name)))
			break
		case "long":
			res = res + fmt.Sprintf(JavaPOField, v.Desc, "Long", firstToLower(camelString(v.Name)))
			break
		case "float":
			res = res + fmt.Sprintf(JavaPOField, v.Desc, "Double", firstToLower(camelString(v.Name)))
			break
		case "string":
			res = res + fmt.Sprintf(JavaPOField, v.Desc, "String", firstToLower(camelString(v.Name)))
			break
		case "time":
			res = res + fmt.Sprintf(JavaPOField, v.Desc, "LocalDateTime", firstToLower(camelString(v.Name)))
			break
		case "date":
			res = res + fmt.Sprintf(JavaPOField, v.Desc, "LocalDate", firstToLower(camelString(v.Name)))
			break
		}
	}
	res = res + fmt.Sprintln(JavaPOClassEnd)
	return res
}
