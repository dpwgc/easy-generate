package generate

import (
	"fmt"
)

type MySQLGenerator struct {
	config Config
}

func NewMySQLGenerator(config Config) Generator {
	return &MySQLGenerator{
		config: config,
	}
}

const (
	MySQLCreateStart  = "CREATE TABLE `%s` ("
	MySQLCreateEnd    = ") COMMENT='%s';\n"
	MySQLField        = "  `%s` %s NOT NULL DEFAULT %s COMMENT '%s',\n"
	MySQLCanNullField = "  `%s` %s COMMENT '%s',\n"
)

func (g *MySQLGenerator) Build() string {
	res := ""
	res = res + fmt.Sprintln(fmt.Sprintf(MySQLCreateStart, g.config.Table))
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
		case "text":
			res = res + fmt.Sprintf(MySQLCanNullField, v.Name, "text", v.Desc)
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
