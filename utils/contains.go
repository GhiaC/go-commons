package utils

func Contains(arrString []string, value string) bool {
	set := make(map[string]bool)
	for _, v := range arrString {
		set[v] = true
	}
	return set[value]
}
