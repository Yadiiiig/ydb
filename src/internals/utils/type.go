package utils

import "reflect"

func GetType(ty string) reflect.Type {
	switch ty {
	case "string":
		return reflect.TypeOf("")

	case "int":
		return reflect.TypeOf(int(0))

	default:
		return nil
	}
}
