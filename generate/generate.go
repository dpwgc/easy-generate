package generate

import "unicode"

type Generator interface {
	Build() string
}

type Config struct {
	Table  string  `yaml:"table"`
	Desc   string  `yaml:"desc"`
	Fields []Field `yaml:"fields"`
}

type Field struct {
	Name string `yaml:"name"`
	Desc string `yaml:"desc"`
	Type string `yaml:"type"`
}

// 下划线转驼峰
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 首字母小写
func firstToLower(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
