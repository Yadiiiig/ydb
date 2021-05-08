package utils

func RowExists(name string, l map[string][]string) bool {
	for _, v := range l {
		for _, i := range v {
			if i == name {
				return true
			}
		}
	}
	return false
}
