package tool

func Argument(s map[string]string) (string, bool) {
	for k, v := range s {
		if v == "" {
			return k, false
		}
	}
	return "", true
}
