package util

import (
	"regexp"
	"reflect"
	"log"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"errors"
	"strings"
	"time"
	"fmt"
	"math"
)

func Contains(s []string, e string) bool {
	return IndexOf(s, e) >= 0
}

func IndexOf(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func IndexMatching(s []interface{}, f func(interface{}) bool) int {
	for i, a := range s {
		if f(a) {
			return i
		}
	}
	return -1
}

func AnyMatch(s []interface{}, f func(interface{}) bool) bool {
	return IndexMatching(s, f) >= 0
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

func ToString(value interface{}, format string) (str string, err error) {

	var data []byte
	switch strings.ToLower(format) {
	case "json":
		data, err = json.MarshalIndent(value, "", "  ")
	case "minified":
		data, err = json.Marshal(value)
	case "yml", "yaml":
		data, err = yaml.Marshal(value)
	default:
		err = errors.New("unrecognized format type: " + format)
	}

	return string(data[:]), err
}

func DurationToString(duration time.Duration) string {

	hours := math.Floor(duration.Hours())
	minutes := math.Floor(duration.Minutes())

	if hours > 0 {
		return fmt.Sprintf("%dm", minutes)
	} else {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
}

func StringToDuration(str string) (time.Duration, error) {
	return time.ParseDuration(strings.Replace(str, " ", "", -1))
}

func PrintLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
