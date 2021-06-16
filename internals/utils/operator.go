package utils

func OperatorQuery(rowVal string, val string, conditional string) bool {
	switch conditional {
	case "=":
		return rowVal == val
	case ">":
		return rowVal > val
	case "<":
		return rowVal < val
	case ">=":
		return rowVal >= val
	case "<=":
		return rowVal <= val
	case "!=":
		return rowVal != val
	}
	return false
}
