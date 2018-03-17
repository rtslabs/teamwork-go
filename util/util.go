package util

import (
	"regexp"
	"reflect"
	"log"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Reverse(list []string) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

var blankRgx = regexp.MustCompile("^\\s*$")

func NotBlank(str string) (b bool) {
	return !blankRgx.MatchString(str)
}

func Overwrite(in interface{}, out interface{}) {

	inPtr := reflect.ValueOf(in).Elem()
	outPtr := reflect.ValueOf(out).Elem()

	for i := 0; i < inPtr.NumField(); i++ {

		inField := inPtr.Field(i)
		outField := outPtr.Field(i)

		switch inField.Type().Kind() {
		case reflect.String:
			str := inField.Interface().(string)
			if NotBlank(str) {
				outField.SetString(str)
			}
		case reflect.Struct:
			Overwrite(inField.Addr().Interface(), outField.Addr().Interface())
		default:
			log.Println("here", in, out, inField.Type().Kind(), inField)

		}
		if inField.Type() == reflect.TypeOf("") {
		}
	}
}
