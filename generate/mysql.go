package generate

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type MySQLGenerator struct {
	config Config
}

func NewMySQLGenerator(text string) Generator {
	var config Config
	err := yaml.Unmarshal([]byte(text), &config)
	if err != nil {
		panic(err)
	}
	return &MySQLGenerator{
		config: config,
	}
}

const (
	MySQLCreateStart = "\nCREATE TABLE `%s` ("
	MySQLCreateEnd   = ") COMMENT='%s';\n"
	MySQLField       = "  `%s` %s NOT NULL DEFAULT %s COMMENT '%s',\n"
)

func (g *MySQLGenerator) Build() string {
	res := ""
	fmt.Println(fmt.Sprintf(MySQLCreateStart, g.config.Table))
	for _, v := range g.config.Fields {
		switch v.Type {
		case "int":
			res = res + fmt.Sprintf(MySQLField, v.Name, "int(11)", "0", v.Desc)
			break
		case "long":
			res = res + fmt.Sprintf(MySQLField, v.Name, "bigint(20)", "0", v.Desc)
			break
		case "float":
			res = res + fmt.Sprintf(MySQLField, v.Name, "decimal(10,2)", "0.0", v.Desc)
			break
		case "string":
			res = res + fmt.Sprintf(MySQLField, v.Name, "varchar(200)", "''", v.Desc)
			break
		case "time":
			res = res + fmt.Sprintf(MySQLField, v.Name, "datetime", "'1900-01-01 00:00:00'", v.Desc)
			break
		case "date":
			res = res + fmt.Sprintf(MySQLField, v.Name, "date", "'1900-01-01'", v.Desc)
			break
		}
	}
	res = res + fmt.Sprintf(MySQLCreateEnd, g.config.Desc)
	return res
}
