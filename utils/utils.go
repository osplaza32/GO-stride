package utils
import (
	"reflect"
)

func TypeOf(s interface{}) string {
	return reflect.TypeOf(s).String()
}