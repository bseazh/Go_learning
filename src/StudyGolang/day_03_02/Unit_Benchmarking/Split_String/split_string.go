package split_string

import "strings"

// Split 字符串

func Split(str string, sep string) []string {

	var res = make([]string, 0, strings.Count(str, sep)+1)

	idx := strings.Index(str, sep)
	sep_len := len(sep)
	for idx >= 0 {
		res = append(res, str[:idx])
		str = str[idx+sep_len:]
		idx = strings.Index(str, sep)
	}
	res = append(res, str)
	return res
}
