package generate

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type JavaDTOGenerator struct {
	config Config
}

func NewJavaDTOGenerator(text string) Generator {
	var config Config
	err := yaml.Unmarshal([]byte(text), &config)
	if err != nil {
		panic(err)
	}
	return &JavaDTOGenerator{
		config: config,
	}
}

const (
	JavaDTOClassStart = `
@Data
@ApiModel(value = "%s")
public class %sDTO {
`
	JavaDTOClassEnd = `}`
	JavaDTOField    = `    /**
     * %s
     */
    @ApiModelProperty(value = "%s")
    private %s %s;

`
)

func (g *JavaDTOGenerator) Build() string {
	res := ""
	res = res + fmt.Sprintln(fmt.Sprintf(JavaDTOClassStart, g.config.Desc, camelString(g.config.Table)))
	for _, v := range g.config.Fields {
		switch v.Type {
		case "int":
			res = res + fmt.Sprintf(JavaDTOField, v.Desc, v.Desc, "Integer", firstToLower(camelString(v.Name)))
			break
		case "long":
			res = res + fmt.Sprintf(JavaDTOField, v.Desc, v.Desc, "Long", firstToLower(camelString(v.Name)))
			break
		case "float":
			res = res + fmt.Sprintf(JavaDTOField, v.Desc, v.Desc, "Double", firstToLower(camelString(v.Name)))
			break
		case "string":
			res = res + fmt.Sprintf(JavaDTOField, v.Desc, v.Desc, "String", firstToLower(camelString(v.Name)))
			break
		case "time":
			res = res + fmt.Sprintf(JavaDTOField, v.Desc, v.Desc, "LocalDateTime", firstToLower(camelString(v.Name)))
			break
		case "date":
			res = res + fmt.Sprintf(JavaDTOField, v.Desc, v.Desc, "LocalDate", firstToLower(camelString(v.Name)))
			break
		}
	}
	res = res + fmt.Sprintln(JavaDTOClassEnd)
	return res
}
