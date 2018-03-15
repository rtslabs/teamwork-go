package util

import "regexp"

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